import React, { useEffect, useContext, useState, useRef, useMemo } from 'react'
import SideBar from './sidebar'
import GameTable from './gameTable'
import axios, { AxiosError, AxiosResponse } from 'axios'
import { useEthers } from '@usedapp/core'
import { GameContext } from '../gameContext'
import { dice, game, user } from '../types/index.d'

interface MainRoomProps {}
const MainRoom = (props: MainRoomProps) => {
  const { account } = useEthers()
  const gameAnte = 10
  let { game, setGame } = useContext(GameContext)
  const gamePot = useMemo(
    () => gameAnte * game.cups.length,
    [game.cups.length, gameAnte],
  )
  let [playerDice, setPlayerDice] = useState([] as dice)

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
    timeoutTime = 2,
    // Get the timer that's set inside the sessionStorage
    sessionTimer = window.sessionStorage.getItem('round_timer')
      ? parseInt(window.sessionStorage.getItem('round_timer') as string) - 1
      : timeoutTime

  let wsStatus = useRef('closed'),
    roundInterval: NodeJS.Timer,
    [timer, setTimer] = useState(sessionTimer)

  const updateStatus = () => {
    axios
      .get(`http://${process.env.REACT_APP_GO_HOST}/status`)
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
              } else if (getActivePlayersLength(response.data) === 1) {
                axios
                  .get(`http://${process.env.REACT_APP_GO_HOST}/reconcile`)
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
        .get(`http://${process.env.REACT_APP_GO_HOST}/start`)
        .then(function () {
          console.info('Game has started!')
        })
        .catch(function (error: AxiosError) {
          console.error(error)
        })
    }
  }

  const rolldice = () => {
    axios
      .get(`http://${process.env.REACT_APP_GO_HOST}/rolldice/${account}`)
      .then(function (response: AxiosResponse) {
        console.info('Rolling the dice')
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
      .get(`http://${process.env.REACT_APP_GO_HOST}/new/${ante}`)
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
      .get(`http://${process.env.REACT_APP_GO_HOST}/newround`)
      .then(function (response: AxiosResponse) {
        console.info('New round!')
        if (response.data) {
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
    const ws = new WebSocket(`ws://${process.env.REACT_APP_GO_HOST}/events`)
    if (
      wsStatus.current !== 'open' &&
      wsStatus.current !== 'attemptingConnection'
    ) {
      ws.onopen = () => {
        updateStatus()
        wsStatus.current = 'open'
        console.info('Socket connected')
      }
      ws.onmessage = (evt: MessageEvent) => {
        updateStatus()
        if (evt.data) {
          let message = evt.data
          switch (message) {
            case 'start':
              rolldice()
              break
            case 'claim':
              window.sessionStorage.removeItem('round_timer')
              setTimer(timeoutTime)
              break
            case 'callliar':
              if (getActivePlayersLength() === 1) {
                console.info('Game finished! Winner is ' + game.cups[0])
                break
              }
              newRound()
              break
          }
        }
        return
      }
      ws.onclose = (evt: CloseEvent) => {
        console.info(
          'Socket is closed. Reconnect will be attempted in 1 second.',
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
    console.info('Joining game...')
    axios
      .get(`http://${process.env.REACT_APP_GO_HOST}/join/${account}`)
      .then(function (response: AxiosResponse) {
        console.info('Welcome to the game!!')
        updateStatus()
      })
      .catch(function (error: AxiosError) {
        console.group('Something went wrong, try again.')
        console.error(error.message)
        console.error('Error:', (error as any).response.data.error)

        console.groupEnd()
      })
  }

  const nextTurn = () => {
    console.info('Next turn')
    axios
      .get(`http://${process.env.REACT_APP_GO_HOST}/next`)
      .then(function (response: AxiosResponse) {
        window.sessionStorage.removeItem('round_timer')
        setTimer(timeoutTime)
        updateStatus()
      })
      .catch(function (error: AxiosError) {
        console.group('Something went wrong, try again.')
        console.error(error.message)
        console.error('Error:', (error as any).response.data.error)

        console.groupEnd()
      })
  }

  const addOut = (
    accountToOut = (game.player_order as string[])[game.current_cup],
  ) => {
    const player = game.cups.filter((player: user) => {
      return player.account === accountToOut
    })
    axios
      .get(
        `http://${process.env.REACT_APP_GO_HOST}/out/${accountToOut}/${
          player[0].outs + 1
        }`,
      )
      .then(function (response: AxiosResponse) {
        console.info('Player timed out and got striked')
        console.log(response.data)
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
      <GameTable playerDice={playerDice} timer={timer} />
    </div>
  )
}

export default MainRoom
