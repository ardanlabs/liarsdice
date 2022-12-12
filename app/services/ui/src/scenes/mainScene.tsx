import React from 'react'
import Phaser from 'phaser'
import { DEFAULT_HEIGHT, DEFAULT_WIDTH } from '../utils/config'
import { apiUrl, axiosConfig } from '../utils/axiosConfig'
import { defaultApiError } from '../types/responses.d'
import axios, { AxiosError, AxiosResponse } from 'axios'
import { bet, dice, die, game, user } from '../types/index.d'
import assureGameType from '../utils/assureGameType'
import getActivePlayersLength from '../utils/getActivePlayers'
import { shortenIfAddress } from '../utils/address'
// import PlayerBalance from '../components/playerBalance'
// var cursors: Phaser.Types.Input.Keyboard.CursorKeys
// var player: Phaser.Types.Physics.Arcade.SpriteWithDynamicBody
var pointer: Phaser.GameObjects.Image
var localGame: game

var account: string | null = window.sessionStorage.getItem('account')

var ENV = 'DEV'

// Create an axios instance to keep the token updated
const axiosInstance = axios.create({
  headers: {
    authorization: window.sessionStorage.getItem('token') as string,
  },
})

// Variables
var playerDice = window.localStorage.getItem('playerDice')

var player: user

// Details bar
var statusText: Phaser.GameObjects.Text,
  roundText: Phaser.GameObjects.Text,
  lastbetText: Phaser.GameObjects.Text,
  lastwinText: Phaser.GameObjects.Text,
  lastlooserText: Phaser.GameObjects.Text,
  accountText: Phaser.GameObjects.Text,
  playerDiceText: Phaser.GameObjects.Text,
  playerOutsText: Phaser.GameObjects.Text

export default class MainScene extends Phaser.Scene {
  ws: WebSocket
  constructor() {
    super({ key: 'MainScene' })
    this.ws = new WebSocket(`ws://${apiUrl}/events`)
    localGame = {
      status: 'nogame',
      lastOut: '-',
      lastWin: '-',
      currentPlayer: '-',
      currentCup: 0,
      round: 1,
      cups: [],
      balances: [],
      playerOrder: [],
      bets: [] as bet[],
      currentID: '-',
      anteUSD: 0,
    }
  }

  preload() {
    this.load.path = 'assets/'
    this.load.image('background', 'images/background.png')
    this.load.image('table', 'images/table.png')
    this.load.image('pointer', 'images/pointer.png')
  }

  create() {
    // We set the background
    this.add.image(DEFAULT_WIDTH / 2, DEFAULT_HEIGHT / 2, 'background')

    // Details bar
    if (ENV === 'DEV') {
      const textSpacing = 20
      statusText = this.add.text(
        textSpacing,
        textSpacing,
        `Status: ${localGame.status}`,
      )
      roundText = this.add.text(
        textSpacing,
        textSpacing * 2,
        `Round: ${localGame.round}`,
      )
      lastbetText = this.add.text(
        textSpacing,
        textSpacing * 3,
        `Last Bet: ${localGame.bets[0]}`,
      )
      lastwinText = this.add.text(
        textSpacing,
        textSpacing * 4,
        `Last Winner: ${localGame.lastWin}`,
      )
      lastlooserText = this.add.text(
        textSpacing,
        textSpacing * 5,
        `Last Looser: ${localGame.lastOut}`,
      )
      accountText = this.add.text(
        textSpacing,
        textSpacing * 6,
        `Account: ${account}`,
      )

      playerDiceText = this.add.text(
        textSpacing,
        textSpacing * 7,
        `Dices: [0,0,0,0,0]`,
      )

      playerOutsText = this.add.text(textSpacing, textSpacing * 8, `Outs: 0`)
    }

    // =========================================================================

    const table = this.physics.add.staticGroup()
    table
      .create(DEFAULT_WIDTH / 2, DEFAULT_HEIGHT / 2, 'table')
      .setScale(0.5)
      .refreshBody()
    pointer = this.add
      .image(DEFAULT_WIDTH / 2, DEFAULT_HEIGHT / 2, 'pointer')
      .setOrigin(0.5, 0.4)
      .setScale(0.2)

    // ws.onopen binds an event listener that triggers with the "open" event.
    this.ws.onopen = (event: any) => {
      console.log(event)
    }

    // ws.onmessage binds an event listener that triggers with "message" event.
    this.ws.onmessage = (evt: MessageEvent) => {
      this.updateStatus()
      if (evt.data) {
        let message = JSON.parse(evt.data)
        const messageAccount = shortenIfAddress(message.address)
        console.log(message)

        // We force a switch in order to check for every type of message.
        switch (message.type) {
          // Message received when the game starts.
          case 'start':
            this.rolldice()
            break
          // // Message received when bet is maded.
          // case 'bet':
          //   this.restart()
          //   break
          // // Message received when next turn is started.
          // case 'nextturn':
          //   this.restart()
          //   break
          // // Message received when a player gets called a liar.
          // case 'callliar':
          //   this.restart()
          //   break
        }
      }
      return
    }

    this.initGame()
  }

  update() {
    pointer.rotation += 0.011

    player = localGame.cups.filter((player: user) => {
      return player.account === localGame.currentID
    })[0]

    if (ENV === 'DEV') {
      statusText.setText(`Status: ${localGame.status}`)
      roundText.setText(`Round: ${localGame.round}`)
      lastbetText.setText(`Last Bet: ${localGame.bets[0]}`)
      lastwinText.setText(`Last Win: ${localGame.lastWin}`)
      lastlooserText.setText(`Last Looser: ${localGame.lastOut}`)
      accountText.setText(`Last Account:  ${account}`)
      playerOutsText.setText(`Outs:  ${player?.outs}`)
      playerDiceText.setText(`Dice:  ${playerDice}`)
    }
  }

  // game functions

  initGame() {
    const initGameAxiosFn = (response: AxiosResponse) => {
      this.setNewGame(response.data)
      if (
        response.data &&
        (localGame.status === 'nogame' || localGame.status === 'reconciled')
      ) {
        this.createNewGame()
        return
      }
      this.joinGame()
    }

    const initGameAxiosErrorFn = (error: AxiosError) => {
      this.createNewGame()
      console.error((error as any).response.data.error)
    }

    axios
      .get(`http://${apiUrl}/status`, axiosConfig)
      .then(initGameAxiosFn)
      .catch(initGameAxiosErrorFn)
  }

  joinGame() {
    // toast.info('Joining game...')

    // catchFn catches the error
    const catchFn = (error: defaultApiError) => {
      const errorMessage = error.response.data.error.replace(/\[[^\]]+\]/gm, '')

      console.log(errorMessage.replace(/\[[^\]]+\]/gm, ''))

      // toast(capitalize(errorMessage))
      console.group()
      console.error('Error:', error.response.data.error)
      console.groupEnd()
    }

    axios
      .get(`http://${apiUrl}/join`, {
        headers: {
          authorization: window.sessionStorage.getItem('token') as string,
        },
      })
      .then(() => {
        console.log('welcome to the game')
        // toast.info('Welcome to the game')
      })
      .catch(catchFn)
  }

  createNewGame() {
    // Sets a new game in the gameContext.
    const createGameFn = (response: AxiosResponse) => {
      if (response.data) {
        const newGame = assureGameType(response.data)
        this.setGame(newGame)
      }
    }

    // Catches the error from the axios call.
    const createGameCatchFn = (error: defaultApiError) => {
      // Figure out regex
      console.log(error)

      // let errorMessage = error.response.data.error.replace(/\[[^\]]+\]/gm, '')
      // toast(capitalize(errorMessage))
      console.group()
      console.error('Error:', error.response.data.error)
      console.groupEnd()
    }

    axiosInstance
      .get(`http://${apiUrl}/new`)
      .then(createGameFn)
      .catch(createGameCatchFn)
  }

  // SetNewGame updates the instance of the game in the local state.
  // Also sets the player dice.
  setNewGame(data: game) {
    const newGame = assureGameType(data)
    if (newGame.cups.length) {
      // We filter the connected player
      const player = newGame.cups.filter((cup) => {
        return cup.account === account
      })
      if (player.length) {
        this.setPlayerDice(player[0].dice)
      }
    }
    this.setGame(newGame)
    return newGame
  }

  setPlayerDice(dice: dice) {
    dice = dice
  }

  setGame(game: game) {
    localGame = game
  }

  // updateStatus calls to the status endpoint and updates the local state.
  updateStatus() {
    // updatesStatusAxiosFn handles the backend answer.
    const updateStatusAxiosFn = (response: AxiosResponse) => {
      console.log(response)

      if (response.data) {
        const parsedGame = this.setNewGame(response.data)
        switch (parsedGame.status) {
          case 'newgame':
            playerDice = '[0,0,0,0,0]'
            if (getActivePlayersLength(parsedGame.cups) >= 2) {
              this.startGame()
            }
            break
          case 'gameover':
            if (
              getActivePlayersLength(parsedGame.cups) >= 1 &&
              parsedGame.lastWin === account
            ) {
              axiosInstance
                .get(`http://${apiUrl}/reconcile`)
                .then(() => {
                  this.updateStatus()
                })
                .catch((error: AxiosError) => {
                  console.error(error)
                })
            }
            break
        }
      }
    }

    axiosInstance
      .get(`http://${apiUrl}/status`)
      .then(updateStatusAxiosFn)
      .catch(function (error: AxiosError) {
        console.error(error as any)
      })
  }

  // startGame starts the game
  startGame() {
    axiosInstance
      .get(`http://${apiUrl}/start`)
      .then(function () {})
      .catch(function (error: AxiosError) {
        console.error(error)
      })
  }

  // nextTurn calls to nextTurn and then updates the status.
  nextTurn() {
    const nextTurnAxiosFn = () => {
      this.updateStatus()
    }

    axiosInstance
      .get(`http://${apiUrl}/next`)
      .then(nextTurnAxiosFn)
      .catch(function (error: AxiosError) {
        console.error(error)
      })
  }

  // Takes an account address and adds an out to that account
  addOut(accountToOut = localGame.currentID) {
    const player = localGame.cups.filter((player: user) => {
      return player.account === accountToOut
    })

    const addOutAxiosFn = (response: AxiosResponse) => {
      this.setNewGame(response.data)
      // If the game didn't stop we call next-turn
      if (response.data.status === 'playing') {
        this.nextTurn()
      }
    }

    axiosInstance
      .get(`http://${apiUrl}/out/${player[0].outs + 1}`)
      .then(addOutAxiosFn)
      .catch(function (error: AxiosError) {
        console.group('Something went wrong, try again.')
        console.error(error)
        console.groupEnd()
      })
  }

  // sendBet sends the player bet to the backend.
  sendBet(number: number, suite: die) {
    axiosInstance
      .get(`http://${apiUrl}/bet/${number}/${suite}`)
      .then()
      .catch(function (error: AxiosError) {
        console.error(error)
      })
  }

  // callLiar triggers the callLiar endpoint
  callLiar() {
    axiosInstance
      .get(`http://${apiUrl}/liar`)
      .catch(function (error: AxiosError) {
        console.error(error)
      })
  }

  // rolldice rolls the player dice.
  rolldice(): void {
    axiosInstance
      .get(`http://${apiUrl}/rolldice`)
      .catch(function (error: AxiosError) {
        console.error(error)
      })
  }
}
