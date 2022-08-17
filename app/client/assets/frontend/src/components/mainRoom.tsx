import React, { useEffect, useContext, useState, useRef } from 'react'
import SideBar from './sidebar'
import GameTable from './gameTable'
import axios, { AxiosError, AxiosResponse } from 'axios'
import { useEthers } from '@usedapp/core'
import { GameContext } from '../gameContext'
import { game, user } from '../types/index.d'

interface MainRoomProps {}
const MainRoom = (props: MainRoomProps) => {
  const { account } = useEthers()
  let { game, setGame } = useContext(GameContext)
  const setNewGame = (data: game) => {
      let newGame = data
      newGame = newGame.cups ? newGame : { ...newGame, cups: [] }
      newGame = newGame.player_order
        ? newGame
        : { ...newGame, player_order: [] }
      setGame(newGame)
      if (newGame.cups.length >= 2 && newGame.status === 'open') {
        startGame()
      }
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
      .get(`http://${process.env.REACT_APP_GO_HOST}/status`)
      .then(function (response: AxiosResponse) {
        if (response.data) {
          setNewGame(response.data)
        }
      })
      .catch(function (error: AxiosError) {
        console.log(error)
      })
  }

  const startGame = () => {
    if (game.status === 'gameover') {
      axios
        .get(`http://${process.env.REACT_APP_GO_HOST}/start`)
        .then(function () {
          console.log('Game has started!')
        })
        .catch(function (error: AxiosError) {
          console.log(error)
        })
    }
  }

  const rolldice = () => {
    axios
      .get(`http://${process.env.REACT_APP_GO_HOST}/rolldice/${account}`)
      .then(function (response: AxiosResponse) {
        console.log('Rolling the dice')
        if (response.data) {
          setNewGame(response.data)
        }
      })
      .catch(function (error: AxiosError) {
        console.log(error)
      })
  }

  const createNewGame = (ante: number = 10000) => {
    axios
      .get(`http://${process.env.REACT_APP_GO_HOST}/new/${ante}`)
      .then(function (response: AxiosResponse) {
        if (response.data) {
          setNewGame(response.data)
        }
      })
      .catch(function (error: AxiosError) {
        console.log(error)
      })
  }

  const newRound = () => {
    axios
      .get(`http://${process.env.REACT_APP_GO_HOST}/newround`)
      .then(function (response: AxiosResponse) {
        console.log('New round!')
        if (response.data) {
          setNewGame(response.data)
        }
      })
      .catch(function (error: AxiosError) {
        console.log(error)
      })
  }

  const connect = () => {
    // console.log(wsStatus)
    const ws = new WebSocket(`ws://${process.env.REACT_APP_GO_HOST}/events`)
    if (
      wsStatus.current !== 'open' &&
      wsStatus.current !== 'attemptingConnection'
    ) {
      ws.onopen = () => {
        wsStatus.current = 'open'
        createNewGame()
        updateStatus()
        console.log('Socket connected')
      }
      ws.onmessage = (evt: MessageEvent) => {
        updateStatus()
        if (evt.data) {
          let message = evt.data
          switch (message) {
            case 'start':
              rolldice()
              break
            case 'callliar':
              if (game.cups.length !== 1) {
                console.log('Game finished! Winner is ' + game.cups[0])
                break
              }
              newRound()
              break
          }
        }
        return
      }
      ws.onclose = (evt: CloseEvent) => {
        console.log(
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
    console.log('Joining game...')
    axios
      .get(`http://${process.env.REACT_APP_GO_HOST}/join/${account}`)
      .then(function (response: AxiosResponse) {
        console.log('Welcome to the game!!')
        updateStatus()
      })
      .catch(function (error: AxiosError) {
        console.group('Something went wrong, try again.')
        console.log(error)
        console.groupEnd()
      })
  }

  const addOut = (accountToOut = game.current_player) => {
    console.log('Player timed out and got striked')
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
        console.log('Welcome to the game!!')
        updateStatus()
      })
      .catch(function (error: AxiosError) {
        console.group('Something went wrong, try again.')
        console.log(error)
        console.groupEnd()
      })
  }

  const decreaseTimer = () => {
    setTimer((prevState) => {
      return prevState - 1
    })
  }

  useEffect(() => {
    if (game.current_player === account) {
      // eslint-disable-next-line react-hooks/exhaustive-deps
      roundInterval = setInterval(() => {
        if (timer > 0) {
          decreaseTimer()
        } else {
          addOut()
          window.sessionStorage.removeItem('round_timer')
          clearInterval(roundInterval)
        }
      }, 1000)
    }
    return () => clearInterval(roundInterval)
  }, [timer, account, game.current_player])

  useEffect(() => {
    connect()
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
      <SideBar joinGame={joinGame} />
      <GameTable timer={timer} />
    </div>
  )
}

export default MainRoom
