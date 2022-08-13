import React, { useEffect, useContext, useState } from 'react'
import SideBar from './sidebar'
import GameTable from './gameTable'
import axios, { AxiosError, AxiosResponse } from 'axios'
import { useEthers } from '@usedapp/core'
import { GameContext } from '../gameContext'
import { game } from '../types/index.d'

interface MainRoomProps {}
const MainRoom = (props: MainRoomProps) => {
  const [wsStatus, setWsStatus] = useState('closed')
  const { account } = useEthers()
  const { game, setGame } = useContext(GameContext)

  const setNewGame = (data: game) => {
    let newGame = data
    newGame = newGame.players ? newGame : { ...newGame, players: [] }
    newGame = newGame.player_order
      ? newGame
      : { ...newGame, player_order: [] }
    setGame(newGame)
    if (newGame.players.length >= 2 && newGame.status === 'open') {
      startGame()
    }
    console.log(newGame.players.length >= 2 && newGame.status === 'open', newGame)
  }

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
    axios
      .post(`http://${process.env.REACT_APP_GO_HOST}/start`)
      .then(function () {
        console.log('Game has started!')
      })
      .catch(function (error: AxiosError) {
        console.log(error)
      })
  }

  const rolldice = () => {
    axios
      .get(`http://${process.env.REACT_APP_GO_HOST}/rolldice/${account}`)
      .then(function (response: AxiosResponse) {
        console.log('Rolling the dice');

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
    if (wsStatus !== 'open') {
      const ws = new WebSocket(`ws://${process.env.REACT_APP_GO_HOST}/events`)
      ws.onopen = () => {
        console.log('Socket connected')
        setWsStatus('open')
        updateStatus()
      }
      ws.onmessage = (evt: MessageEvent) => {
        updateStatus()
        if (evt.data) {
          console.log(evt.data)
          let message = evt.data
          switch (message) {
            case 'start':
              rolldice()
              break
            case 'callliar':
              if (game.players.length !== 1) {
                console.log('Game finished! Winner is ' + game.players[0])
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
        setWsStatus('closed')
        setTimeout(function () {
          setGame({
            status: 'open',
            round: 0,
            current_player: '',
            player_order: [],
            players: [],
          })
          connect()
        }, 1000)
      }

      ws.onerror = function (err) {
        console.error('Socket encountered error: ', err, 'Closing socket')
        ws.close()
        setWsStatus('close')
      }
    }
  }

  useEffect(() => {
    connect()
  }, [])

  const joinGame = () => {
    console.log('Joining game...')
    axios
      .post('http://localhost:3000/v1/game/join', {
        wallet: account,
      })
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

  return (
    <div
      style={{
        width: '100%',
        display: 'flex',
        justifyContent: 'start',
        alignItems: 'center',
        maxWidth: '100vw',
      }}
      id="mainRoom"
    >
      <SideBar joinGame={joinGame} />
      <GameTable />
    </div>
  )
}

export default MainRoom
