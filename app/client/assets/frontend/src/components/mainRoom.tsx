import React, { useEffect, useContext } from 'react'
import { claim } from '../types/index.d'
import SideBar from './sidebar'
import GameTable from './gameTable'
import axios, { AxiosError, AxiosResponse } from 'axios'
import { useEthers } from '@usedapp/core'
import { GameContext } from '../gameContext'

interface MainRoomProps {}
const MainRoom = (props: MainRoomProps) => {
  const { account } = useEthers()
  const { game, setGame } = useContext(GameContext)
  
  const connect = () => {
    const ws = new WebSocket(`ws://${process.env.REACT_APP_GO_HOST}/events`)
    ws.onopen = () => {
      console.log('Socket connected')
      try {
        axios
        .get(`http://${process.env.REACT_APP_GO_HOST}/status`)
        .then(function (response: AxiosResponse) {
          if (response.data) {
            let newGame = response.data
            newGame = newGame.players?.length ? newGame : {...newGame, players: []}
            newGame = newGame.player_order ? newGame : {...newGame, player_order: []}
            setGame(newGame)
          }
        })
        .catch(function (error: AxiosError) {
          console.log(error)
        })
      } catch (error) {
        console.error(error)
      }
    }
    ws.onmessage = (evt: MessageEvent) => {
      console.log(evt)
      if (evt.data) {
        let response = evt.data
      }
      return
    }
    ws.onclose = (evt: CloseEvent) => {
      console.log(
        'Socket is closed. Reconnect will be attempted in 1 second.',
        evt.reason,
      )
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
    }
  }

  const updateStatus = () => {
    axios
      .get(`http://${process.env.REACT_APP_GO_HOST}/status`)
      .then(function (response: AxiosResponse) {
        if (response.data) {
          let newGame = response.data
          newGame = newGame.players?.length
            ? newGame
            : { ...newGame, players: [] }
          newGame = newGame.player_order
            ? newGame
            : { ...newGame, player_order: [] }
          setGame(newGame)
        }
      })
      .catch(function (error: AxiosError) {
        console.log(error)
      })
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
