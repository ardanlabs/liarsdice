import React from 'react'
import Phaser from 'phaser'
import { DEFAULT_HEIGHT, DEFAULT_WIDTH, DICE_SPACING } from '../utils/config'
import { apiUrl, axiosConfig } from '../utils/axiosConfig'
import { defaultApiError } from '../types/responses.d'
import axios, { AxiosError, AxiosResponse } from 'axios'
import { bet, dice, DiceConfigs, die, game, user } from '../types/index.d'
import assureGameType from '../utils/assureGameType'
import getActivePlayersLength from '../utils/getActivePlayers'
import { shortenIfAddress } from '../utils/address'
import getDicePosition from '../utils/diceRotations'

// Configs
var ENV = 'DEV'
// Create an axios instance to keep the token updated
const axiosInstance = axios.create({
  headers: {
    authorization: window.sessionStorage.getItem('token') as string,
  },
})

// BackendGame Variables
var playerDice = window.localStorage.getItem('playerDice')
var localGame: game
var account: string | null = window.sessionStorage.getItem('account')
var player: user

// UI Variables
var pointer: Phaser.GameObjects.Image
var table: Phaser.GameObjects.Image
// var diceContainer2: Phaser.GameObjects.Container
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
  dieConfig: Phaser.Types.GameObjects.Sprite.SpriteConfig
  x: number
  y: number
  rotation: number
  initialYPos: number
  center = { x: DEFAULT_WIDTH / 2, y: DEFAULT_HEIGHT / 2 }
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

    this.initialYPos = DEFAULT_HEIGHT / 2 + 310
    this.x = this.center.x - DICE_SPACING * 2.5
    this.y = this.initialYPos
    this.rotation = 0.3

    this.dieConfig = {
      key: 'dice',
      scale: 0.8,
      // anims: {
      //   key: 'die',
      //   repeat: -1,
      //   repeatDelay: { randInt: [1000, 4000] },
      //   delayedPlay: function () {
      //     return Math.random() * 6000
      //   },
      // },
    }
  }

  preload() {
    this.load.path = 'assets/'
    this.load.image('background', 'images/background.png')
    this.load.image('table', 'images/table.png')
    this.load.image('pointer', 'images/pointer.png')
    this.load.atlas('dice', 'animations/dice.png', 'animations/dice.json')
    this.load.image('die_0', 'images/die_0.png')
  }

  create() {
    // We set the background
    this.add.image(this.center.x, this.center.y, 'background')
    table = this.add.image(this.center.x, this.center.y, 'table').setScale(0.65)
    // table = this.physics.add.staticGroup()
    // table
    //   .create(this.center.x, this.center.y, 'table')
    //   .setScale(0.65)
    //   .refreshBody()
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

    this.anims.create({
      key: 'die',
      frames: this.anims.generateFrameNames('dice', {
        prefix: 'die_',
        end: 6,
        zeroPad: 4,
      }),
      repeat: -1,
    })

    // const die = this.physics.add.sprite(
    //   this.center.x,
    //   200,
    //   'dice',
    //   'die_6',
    // )
    // die.setCollideWorldBounds(true)
    // this.physics.add.collider(die, table)

    // =========================================================================
    pointer = this.add
      .image(this.center.x, this.center.y, 'pointer')
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

  renderDice(game: game) {
    // Position dices and multiple them by amount of players.
    // Figure out an algorithm to calculate the position of the players
    // game.cups.forEach((user: user) => {
    const userDice: dice = [
      1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5,
    ]
    //

    userDice.forEach((dieNumber: die, i: number) => {
      const position = getDicePosition(
        this.initialYPos,
        this.x,
        this.y,
        this.rotation,
        i,
      )

      this.initialYPos = position.initialYPos
      this.x = position.position.x
      this.y = position.position.y
      this.rotation = position.position.rotation

      if (dieNumber !== 0) {
        const die = this.make.sprite({
          ...this.dieConfig,
          frame: `die_${dieNumber}`,
          ...position.position,
        })
      }
      // if (dieNumber === 0) {
      //   const die = this.make
      //     .sprite({ key: 'die_0', x: this.x, y: this.y })
      //     .setRotation(this.rotation)
      //   diceContainer.add(die)
      // }
    })
    // })
  }

  // game functions

  initGame() {
    this.renderDice(localGame)
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
      if (response.data) {
        const parsedGame = this.setNewGame(response.data)
        this.renderDice(parsedGame)
        switch (parsedGame.status) {
          case 'newgame':
            window.localStorage.removeItem('playerDice')
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
                  window.localStorage.removeItem('playerDice')
                  this.updateStatus()
                })
                .catch((error: AxiosError) => {
                  console.error(error)
                })
            }
            break
          case 'nogame':
            window.localStorage.removeItem('playerDice')
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
