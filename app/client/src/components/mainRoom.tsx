import React, { useEffect, useContext, useState, useRef, useMemo } from 'react'
import SideBar from './sidebar'
import GameTable from './gameTable'
import axios, { AxiosError, AxiosResponse } from 'axios'
import { shortenAddress, useEthers } from '@usedapp/core'
import { GameContext } from '../gameContext'
import { dice, game, user } from '../types/index.d'
import { toast } from 'react-toastify'
import { capitalize } from '../utils/capitalize'
import Test from './test'

interface MainRoomProps {}
const MainRoom = (props: MainRoomProps) => {
  const { account } = useEthers()
  const gameAnte = 10
  const [timeoutsCount, setTimeoutsCount] = useState(0)
  let { game, setGame } = useContext(GameContext)
  const gamePot = useMemo(
    () => gameAnte * game.cups.length,
    [game.cups.length, gameAnte],
  )
  let [playerDice, setPlayerDice] = useState([] as dice)
  const apiUrl = process.env.REACT_APP_GO_HOST
    ? process.env.REACT_APP_GO_HOST
    : 'localhost:3000/v1/game'

  useEffect(() => {
    setPlayerDice(
      JSON.parse(window.localStorage.getItem('playerDice') as string),
    )
  }, [])

  useEffect(() => {
    window.localStorage.setItem('playerDice', JSON.stringify(playerDice))
  }, [playerDice])

  const setNewGame = (data: game) => {
      let newGame = data
      newGame = newGame.claims ? newGame : { ...newGame, claims: [] }
      newGame = newGame.cups ? newGame : { ...newGame, cups: [] }
      newGame = newGame.player_order
        ? newGame
        : { ...newGame, player_order: [] }
      setGame(newGame)
    },
    // Timer time in seconds
    timeoutTime = 30,
    // Get the timer that's set inside the sessionStorage
    sessionTimer = window.sessionStorage.getItem('round_timer')
      ? parseInt(window.sessionStorage.getItem('round_timer') as string) - 1
      : timeoutTime

  let wsStatus = useRef('closed'),
    roundInterval: NodeJS.Timer,
    [timer, setTimer] = useState(sessionTimer)

  const updateStatus = () => {
    axios
      .get(`http://${apiUrl}/status`)
      .then(function (response: AxiosResponse) {
        if (response.data) {
          setNewGame(response.data)
          switch (response.data.status) {
            case 'roundover':
              newRound()
              break
            case 'gameover':
              if (getActivePlayersLength(response.data) >= 2) {
                startGame()
              } else if (
                getActivePlayersLength(response.data) === 1 &&
                game.last_win !== ''
              ) {
                axios
                  .get(`http://${apiUrl}/reconcile`)
                  .then((response: AxiosResponse) => {
                    console.info(response)
                  })
                  .catch((error: AxiosError) => {
                    console.error(error)
                  })
              }
              break
          }
        }
      })
      .catch(function (error: AxiosError) {
        if (
          error.response &&
          (error as any).response.data.error === 'no game exists'
        ) {
          createNewGame()
        }
        console.error(error)
      })
  }

  const startGame = () => {
    if (game.status === 'gameover') {
      axios
        .get(`http://${apiUrl}/start`)
        .then(function () {})
        .catch(function (error: AxiosError) {
          console.error(error)
        })
    }
  }

  const rolldice = () => {
    axios
      .get(`http://${apiUrl}/rolldice/${account}`)
      .then(function (response: AxiosResponse) {
        if (response.data) {
          setPlayerDice(response.data.dice)
        }
      })
      .catch(function (error: AxiosError) {
        console.error(error)
      })
  }

  const createNewGame = (ante: number = gameAnte) => {
    axios
      .get(`http://${apiUrl}/new/${ante}`)
      .then(function (response: AxiosResponse) {
        if (response.data) {
          setNewGame(response.data)
        }
      })
      .catch(function (error: AxiosError) {
        console.error(error)
      })
  }

  const newRound = () => {
    axios
      .get(`http://${apiUrl}/newround`)
      .then(function (response: AxiosResponse) {
        if (response.data) {
          window.sessionStorage.removeItem('round_timer')
          setTimer(timeoutTime)
          setTimeoutsCount(0)
          updateStatus()
          rolldice()
        }
      })
      .catch(function (error: AxiosError) {
        console.error(error)
      })
  }

  const getActivePlayersLength = (gameToFilter = game) => {
    return gameToFilter.player_order.filter((player: string) => player.length)
      .length
  }

  const connect = () => {
    const ws = new WebSocket(`ws://${apiUrl}/events`)
    if (
      wsStatus.current !== 'open' &&
      wsStatus.current !== 'attemptingConnection'
    ) {
      ws.onopen = () => {
        updateStatus()
        wsStatus.current = 'open'
        toast.info('Connection established')
      }
      ws.onmessage = (evt: MessageEvent) => {
        updateStatus()
        if (evt.data) {
          let message = evt.data
          switch (true) {
            case message === 'start':
              toast.info('Game has started!')
              rolldice()
              break
            case message === 'rolldice':
              toast.info(`Rolling dice's`)
              break
            case message.startsWith('join:'):
              const joinStart = 'join:'
              const joinAccount = shortenAddress(
                message.substring(joinStart.length),
              )
              toast.info(`Account ${joinAccount} just joined`)
              break
            case message === 'claim':
              window.sessionStorage.removeItem('round_timer')
              setTimer(timeoutTime)
              break
            case message === 'newround':
              toast.info('New round!')
              break
            case message === 'nextturn':
              toast.info('Next turn!')
              break
            case message.startsWith('outs:'):
              const outsStart = 'join:'
              const strikedAccount = shortenAddress(
                message.substring(outsStart.length),
              )
              toast.info(`Player ${strikedAccount} timed out and got striked`)
              break
            case message === 'callliar':
              if (getActivePlayersLength() === 1) {
                toast.info('Game finished! Winner is ' + game.cups[0])
                break
              }
              toast.info('A player was called a liar and loose!')
              newRound()
              break
          }
        }
        return
      }
      ws.onclose = (evt: CloseEvent) => {
        toast.error(
          'Connection is closed. Reconnect will be attempted in 1 second. ' +
            evt.reason,
        )
        wsStatus.current = 'closed'
        setTimeout(function () {
          setGame({
            status: 'gameover',
            last_out: '',
            last_win: '',
            current_player: '',
            current_cup: 0,
            round: 1,
            cups: [],
            player_order: [],
            claims: [],
          })
          connect()
        }, 1000)
      }

      ws.onerror = function (err) {
        console.error('Socket encountered error: ', err, 'Closing socket')
        ws.close()
        wsStatus.current = 'close'
      }
    }
  }

  const joinGame = () => {
    toast.info('Joining game...')
    axios
      .get(`http://${apiUrl}/join/${account}`)
      .then(function (response: AxiosResponse) {
        toast.info('Welcome to the game')
        updateStatus()
      })
      .catch(function (error: AxiosError) {
        let errorMessage = (error as any).response.data.error.replace(
          / \[.+\]/gm,
          '',
        )
        toast.error(
          <div style={{ textAlign: 'start' }}>{capitalize(errorMessage)}</div>,
        )
        console.group()
        console.error('Error:', (error as any).response.data.error)
        console.groupEnd()
      })
  }

  const nextTurn = () => {
    axios
      .get(`http://${apiUrl}/next`)
      .then(function (response: AxiosResponse) {
        window.sessionStorage.removeItem('round_timer')
        setTimer(timeoutTime)
        if (getActivePlayersLength(response.data) - 1 === timeoutsCount) {
          newRound()
        } else {
          updateStatus()
        }
      })
      .catch(function (error: AxiosError) {
        // To be discused
      })
  }

  const addOut = (
    accountToOut = (game.player_order as string[])[game.current_cup],
  ) => {
    const player = game.cups.filter((player: user) => {
      return player.account === accountToOut
    })
    axios
      .get(`http://${apiUrl}/out/${accountToOut}/${player[0].outs + 1}`)
      .then(function (response: AxiosResponse) {
        setTimeoutsCount((prev) => prev + 1)
        setNewGame(response.data)
        if (response.data.status === 'playing') {
          nextTurn()
        }
      })
      .catch(function (error: AxiosError) {
        console.group('Something went wrong, try again.')
        console.error(error)
        console.groupEnd()
      })
  }

  const decreaseTimer = () => {
    setTimer((prevState) => {
      return prevState - 1
    })
  }
  useEffect(() => {
    window.sessionStorage.setItem('round_timer', `${timer}`)
  }, [timer])

  useEffect(() => {
    if (
      (game.player_order as string[])[game.current_cup] === account &&
      game.status === 'playing'
    ) {
      clearInterval(roundInterval)
      // eslint-disable-next-line react-hooks/exhaustive-deps
      roundInterval = setInterval(() => {
        if (timer > 0 && game.status === 'playing') {
          decreaseTimer()
        } else {
          addOut()
          window.sessionStorage.removeItem('round_timer')
          clearInterval(roundInterval)
        }
      }, 1000)
    } else {
      clearInterval(roundInterval)
    }
    return () => clearInterval(roundInterval)
  }, [timer, account, game.player_order, game.current_cup, game.status])

  useEffect(() => {
    connect()
    window.sessionStorage.removeItem('round_timer')
    wsStatus.current = 'attemptingConnection'
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  return (
    <div
      style={{
        width: '100%',
        display: 'flex',
        justifyContent: 'start',
        alignItems: 'center',
        maxWidth: '100vw',
        marginTop: '15px',
      }}
      id="mainRoom"
    >
      <SideBar ante={gameAnte} gamePot={gamePot} joinGame={joinGame} />
      <Test />
      <GameTable playerDice={playerDice} timer={timer} />
    </div>
  )
}

export default MainRoom
