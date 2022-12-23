import React from 'react'
import Phaser from 'phaser'
import { DEFAULT_HEIGHT, DEFAULT_WIDTH, DIE_PER_PLAYER } from '../utils/config'
import { apiUrl, axiosConfig } from '../utils/axiosConfig'
import { defaultApiError } from '../types/responses.d'
import axios, { AxiosError, AxiosResponse } from 'axios'
import { bet, dice, die, game, user } from '../types/index.d'
import assureGameType from '../utils/assureGameType'
import getActivePlayersLength from '../utils/getActivePlayers'
import { shortenIfAddress } from '../utils/address'
import getRotationOnCircle from '../utils/getRotationOnCircle'

// Configs
var ENV = 'PROD'
// Create an axios instance to keep the token updated
const axiosInstance = axios.create({
  headers: {
    authorization: window.sessionStorage.getItem('token') as string,
  },
})

// BackendGame Variables
var playerDice = window.localStorage.getItem('playerDice')
var localGame: game
var account: string | null = window.localStorage.getItem('account')
var player: user
var currentBet: { number: number; suite: die } = { number: 1, suite: 1 }

// UI Variables
var pointer: Phaser.GameObjects.Image
var table: Phaser.GameObjects.Image
var currentDiceAmountText: Phaser.GameObjects.Text
var diceBetButtonsGroup: Phaser.GameObjects.Group
var playersGroup: Phaser.GameObjects.Group
var diceBetButtons: Phaser.GameObjects.Sprite[] = []
var showBetButtons: boolean
var firstPlayerPosition: number
const textSpacing = 20

// Details bar
var statusText: Phaser.GameObjects.Text,
  roundText: Phaser.GameObjects.Text,
  lastbetText: Phaser.GameObjects.Text,
  lastwinText: Phaser.GameObjects.Text,
  lastlooserText: Phaser.GameObjects.Text,
  accountText: Phaser.GameObjects.Text,
  playerDiceText: Phaser.GameObjects.Text,
  playerOutsText: Phaser.GameObjects.Text,
  resetGameButton: Phaser.GameObjects.Sprite,
  phaserDice: {
    [key: number]: {
      dieNumber: number
      die: Phaser.GameObjects.Sprite
    }[]
  },
  playersOuts: {
    [key: number]: {
      star: Phaser.GameObjects.Sprite
    }[]
  }

export default class MainScene extends Phaser.Scene {
  ws: WebSocket
  dieConfig
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

    this.dieConfig = {
      key: 'dice',
      scale: 0.8,
    }
  }

  preload() {
    this.load.path = 'assets/'
    this.load.image('background', 'images/background.png')
    this.load.image('table', 'images/table.png')
    this.load.image('pointer', 'images/pointer.png')
    this.load.atlas('dice', 'animations/dice.png', 'animations/dice.json')
    this.load.image('die_0', 'images/die_0.png')
    this.load.image('resetGame', 'images/resetGame.png')
    this.load.image('placeBet', 'images/placeBet.png')
    this.load.image('callLiar', 'images/callLiar.png')
    this.load.image('player0', 'images/player0.png')
    this.load.image('player1', 'images/player1.png')
    this.load.image('player2', 'images/player2.png')
    this.load.image('player3', 'images/player3.png')
    this.load.image('player4', 'images/player4.png')
    this.load.image('filledStar', 'images/filledStar.png')
    this.load.image('emptyStar', 'images/emptyStar.png')
  }

  create() {
    // We set the background
    this.add.image(this.center.x, this.center.y, 'background')
    // Set the table
    table = this.add.image(this.center.x, this.center.y, 'table').setScale(0.4)

    // pointer = this.add
    //   .image(this.center.x, this.center.y, 'pointer')
    //   .setOrigin(0.5, 0.4)
    //   .setScale(0.2)

    this.anims.create({
      key: 'dieAnimation',
      frames: this.anims.generateFrameNames('dice', {
        prefix: 'die_',
        frames: [7, 8, 9, 10, 11],
      }),
      frameRate: 8,
      repeat: -1,
    })

    lastbetText = this.add
      .text(
        this.center.x,
        this.center.y,
        `${localGame.bets[localGame.bets.length - 1]?.number || '-'} X ${
          localGame.bets[localGame.bets.length - 1]?.suite || '-'
        }`,
        { fontSize: '50px' },
      )
      .setOrigin(0.5, 0.5)

    roundText = this.add
      .text(
        this.center.x,
        this.center.y + textSpacing * 2,
        `Round: ${localGame.round}`,
      )
      .setOrigin(0.5, 0.5)

    statusText = this.add
      .text(
        this.center.x,
        this.center.y + textSpacing * 3,
        `Status: ${localGame.status}`,
      )
      .setOrigin(0.5, 0.5)

    firstPlayerPosition = this.center.x - 250

    // Details bar
    if (ENV === 'DEV') {
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

      resetGameButton = this.add
        .sprite(textSpacing, textSpacing * 9, 'resetGame')
        .setOrigin(0, 0)
        .setInteractive()

      const resetGameFn = () => {
        this.createNewGame()
      }

      resetGameButton.on('pointerdown', resetGameFn)
    }

    // =========================================================================
    diceBetButtonsGroup = this.add.group()
    const firstDiceBetPosition = this.center.x - 234.5
    const diceBetPosition = DEFAULT_HEIGHT - 80

    currentDiceAmountText = this.add.text(
      firstDiceBetPosition - 100,
      diceBetPosition - 25,
      `${currentBet.number} X`,
      {
        fontSize: '50px',
      },
    )

    diceBetButtonsGroup.add(currentDiceAmountText)

    for (let i: die = 1; i < 7; i++) {
      const diceBetButton = this.add
        .sprite(
          firstDiceBetPosition + 67 * i,
          diceBetPosition,
          'dice',
          `die_${i}`,
        )
        .setInteractive()

      diceBetButtonsGroup.add(diceBetButton)
      diceBetButtons[i] = diceBetButton

      const setCurrentBet = () => {
        diceBetButtons.forEach((button) => {
          button.clearTint()
        })
        diceBetButton.setTint(0xffff)
        if (currentBet.suite === i) {
          currentBet.number++
          return
        }
        currentBet.suite = i
        currentBet.number = localGame.bets[localGame.bets.length - 1]
          ? localGame.bets[localGame.bets.length - 1].number
          : 1
      }

      diceBetButton.on('pointerdown', setCurrentBet)
    }

    const placeBetButton = this.add
      .sprite(this.center.x + 100, DEFAULT_HEIGHT - 25, 'placeBet')
      .setInteractive()

    const placeBet = () => {
      this.sendBet(currentBet.number, currentBet.suite)
      diceBetButtons.forEach((button) => {
        button.clearTint()
      })
      currentBet.number = localGame.bets[localGame.bets.length - 1]
        ? localGame.bets[localGame.bets.length - 1].number
        : 1
      currentBet.suite = 1
    }

    placeBetButton.on('pointerdown', placeBet)

    diceBetButtonsGroup.add(placeBetButton)

    const callLiarButton = this.add
      .sprite(this.center.x - 100, DEFAULT_HEIGHT - 25, 'callLiar')
      .setInteractive()

    const callLiarFn = () => {
      this.callLiar()
    }

    callLiarButton.on('pointerdown', callLiarFn)

    diceBetButtonsGroup.add(callLiarButton)

    // =========================================================================
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

        // We force a switch in order to check for every type of message.
        switch (message.type) {
          // Message received when the game starts.
          case 'start':
            this.rolldice()
            break
          case 'rolldice':
            this.renderDice(localGame)
            break
          case 'newgame':
            this.joinGame()
            break
          // Message received when bet is maded.
          case 'bet':
            currentBet.number =
              localGame.bets[localGame.bets.length - 1]?.number || 1
            this.update()
            break
          // // Message received when next turn is started.
          // case 'nextturn':
          //   this.restart()
          //   break
          // Message received when a player gets called a liar.
          case 'reconcile':
            this.deleteDice()
            showBetButtons = false
            break
        }
      }
      return
    }

    this.initGame()
  }

  update() {
    player = localGame.cups.filter((player: user) => {
      return player.account === localGame.currentID
    })[0]

    lastbetText.setText(
      `${localGame.bets[localGame.bets.length - 1]?.number || '-'} X ${
        localGame.bets[localGame.bets.length - 1]?.suite || '-'
      }`,
    )

    if ('children' in diceBetButtonsGroup && diceBetButtonsGroup.getLength()) {
      currentDiceAmountText.setText(`${currentBet.number} X`)
      diceBetButtonsGroup.setVisible(showBetButtons)
    }

    statusText.setText(`Status: ${localGame.status}`)
    roundText.setText(`Round: ${localGame.round}`)
    if (ENV === 'DEV') {
      lastwinText.setText(`Last Win: ${localGame.lastWin}`)
      lastlooserText.setText(`Last Looser: ${localGame.lastOut}`)
      accountText.setText(`Account:  ${account}`)
      playerOutsText.setText(`Outs:  ${player?.outs}`)
      playerDiceText.setText(`Dice:  ${playerDice}`)
    }
  }

  renderPlayers() {
    playersGroup = this.add.group()

    playersOuts = {}

    localGame.cups?.forEach((player, i: number) => {
      playersOuts[i] = []
      const playerSprite = this.add.sprite(
        firstPlayerPosition + 150 * i,
        60,
        `player${i}`,
      )

      const playerSpriteText = this.add
        .text(
          firstPlayerPosition + 150 * i,
          60,
          `${shortenIfAddress(player.account)}`,
          {
            fontSize: '14px',
          },
        )
        .setOrigin(0.5, 0.5)

      playersGroup.addMultiple([playerSprite, playerSpriteText])
    })
  }

  renderDice(game: game) {
    this.deleteDice()

    phaserDice = {}
    const cups = game.cups.sort((a: user, b: user) => {
      switch (true) {
        case a.account === b.account:
          return 0
        case a.account === account?.toLowerCase():
          return -1
        default:
          return 1
      }
    })

    // Position dices and multiple them by amount of players.
    cups.forEach((user: user, p: number) => {
      phaserDice[p] = []
      const userDice: dice = user.dice || [0, 0, 0, 0, 0]

      // Angle in radians
      let angle = 1.97
      // Distance between dice
      let angleStep = -1 / DIE_PER_PLAYER

      // Circle to serve as container for the dice
      const diceCircle = new Phaser.Geom.Circle(
        this.center.x,
        this.center.y,
        table.displayWidth / 2 - 70,
      )

      const angleOffset = 0.74

      const angleEqualDistribution = 1.256638

      angle = angleOffset + angleEqualDistribution * (p + 1)

      userDice.forEach((dieNumber: die) => {
        const position = getRotationOnCircle(diceCircle, angle, angleStep)

        const { x, y, rotation } = position.position

        if (dieNumber !== 0) {
          const die = this.add.sprite(x, y, 'dice', `die_${dieNumber}`)
          die.setAngle(rotation)
          die.setScale(0.6)
          phaserDice[p].push({ dieNumber, die })
        }

        if (dieNumber === 0) {
          const die = this.add.sprite(x, y, 'die_0')
          die.setAngle(rotation)
          die.setScale(0.7)
          phaserDice[p].push({ dieNumber, die })
        }
        angle = position.angle
      })

      console.log(playersOuts)
    })
  }

  renderOuts(game: game) {
    this.deleteStars()
    const cups = game.cups.sort((a: user, b: user) => {
      switch (true) {
        case a.account === b.account:
          return 0
        case a.account === account?.toLowerCase():
          return -1
        default:
          return 1
      }
    })
    // Position dices and multiple them by amount of players.
    cups.forEach((user: user, p: number) => {
      playersOuts[p] = []
      // Angle in radians
      let starAngle = 2.1
      // Distance between stars (3 outs)
      let starAngleStep = -1 / 3
      const starAngleOffset = 0.68

      const starAngleEqualDistribution = 1.256638

      starAngle = starAngleOffset + starAngleEqualDistribution * (p + 1)

      const starCircle = new Phaser.Geom.Circle(
        this.center.x,
        this.center.y,
        table.displayWidth / 2 - 20,
      )

      for (let i = 1; i <= 3; i++) {
        const position = getRotationOnCircle(
          starCircle,
          starAngle,
          starAngleStep,
        )

        const { x, y, rotation } = position.position
        const star = this.add.sprite(
          x,
          y,
          i > user.outs ? 'emptyStar' : 'filledStar',
        )
        star.setAngle(rotation)
        star.setScale(0.6)
        playersOuts[p].push({ star })
        playersGroup.add(star)
        starAngle = position.angle
      }
    })
  }

  stopDiceAnimation() {
    for (let i = 0; i < Object.keys(phaserDice).length; i++) {
      const element = phaserDice[i]
      element.forEach((die) => {
        if (die.die.anims.isPlaying) {
          die.die.stop()
          die.die.setFrame(`die_${die.dieNumber}`)
          return
        }
      })
    }
  }
  startDiceAnimation(playerPosition: number) {
    phaserDice[playerPosition].forEach((die) => {
      die.die.play('dieAnimation')
    })
  }
  async deleteDice() {
    if (phaserDice) {
      for (let i = 0; i < Object.keys(phaserDice).length; i++) {
        const element = phaserDice[i]
        element.forEach((die) => {
          die.die.destroy()
        })
      }
    }
  }

  async deleteStars() {
    if (playersOuts) {
      for (let i = 0; i < Object.keys(playersOuts).length; i++) {
        const element = playersOuts[i]
        element.forEach((star) => {
          star.star.destroy()
        })
      }
    }
  }

  // game functions

  initGame() {
    const initGameAxiosFn = (response: AxiosResponse) => {
      const parsedGame = this.setNewGame(response.data)
      this.renderPlayers()
      this.renderOuts(response.data)

      if (
        parsedGame &&
        (parsedGame.status === 'nogame' || parsedGame.status === 'reconciled')
      ) {
        this.createNewGame()
        return
      }
      if (parsedGame.status === 'newgame') this.joinGame()
      if (parsedGame.status === 'playing') {
        currentBet = {
          number: parsedGame.bets[localGame.bets.length - 1]?.number || 1,
          suite: parsedGame.bets[localGame.bets.length - 1]?.suite || 1,
        }
        showBetButtons = parsedGame.currentID === account?.toLowerCase()
        this.renderDice(response.data)
      }
    }

    const initGameAxiosErrorFn = (error: AxiosError) => {
      console.log(error)

      // console.error((error as any).response.data.error)
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
      // console.group()
      // console.error('Error:', error.response.data.error)
      // console.groupEnd()
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
    this.setGame(newGame)
    this.update()
    if (newGame.cups.length && newGame.status === 'playing') {
      // We filter the connected player
      const player = newGame.cups.filter((cup) => {
        return cup.account === account?.toLowerCase()
      })
      if (player.length) {
        this.setPlayerDice(player[0].dice)
      }
    }
    return newGame
  }

  setPlayerDice(dice: dice) {
    const parsedDice = JSON.stringify(dice)
    window.localStorage.setItem('playerDice', parsedDice)
    playerDice = parsedDice
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
        this.renderPlayers()
        this.renderOuts(response.data)
        switch (parsedGame.status) {
          case 'newgame':
            window.localStorage.removeItem('playerDice')
            this.deleteDice()
            if (getActivePlayersLength(parsedGame.cups) >= 2) {
              this.startGame()
            }
            break
          case 'gameover':
            if (
              getActivePlayersLength(parsedGame.cups) >= 1 &&
              parsedGame.lastWin === account?.toLowerCase()
            ) {
              axiosInstance
                .get(`http://${apiUrl}/reconcile`)
                .then(() => {
                  this.deleteDice()
                  showBetButtons = false
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
          case 'playing':
            currentBet = {
              number: parsedGame.bets[localGame.bets.length - 1]?.number || 1,
              suite: parsedGame.bets[localGame.bets.length - 1]?.suite || 1,
            }
            // If it's player turn we show the betting section
            showBetButtons = parsedGame.currentID === account?.toLowerCase()
            this.renderDice(parsedGame)
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
    this.renderDice(localGame)
    axiosInstance
      .get(`http://${apiUrl}/rolldice`)
      .catch(function (error: AxiosError) {
        console.error(error)
      })
  }
}
