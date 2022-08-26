import React, { useContext } from 'react'
import { useEthers } from '@usedapp/core'
import axios, { AxiosError, AxiosResponse } from 'axios'
import Button from './button'
import { toast } from 'react-toastify'
import { capitalize } from '../utils/capitalize'
import { axiosConfig } from '../utils/axiosConfig'
import { GameContext } from '../gameContext'
import { game, user } from '../types/index.d'

interface JoinProps {
  disabled: boolean
}

const Join = (props: JoinProps) => {
  const { game, setGame } = useContext(GameContext)
  const { account } = useEthers()
  const apiUrl = process.env.REACT_APP_GO_HOST
    ? process.env.REACT_APP_GO_HOST
    : 'localhost:3000/v1/game'

  const setNewGame = (data: game) => {
    let newGame = data
    newGame = newGame.claims ? newGame : { ...newGame, claims: [] }
    newGame = newGame.cups ? newGame : { ...newGame, cups: [] }
    newGame = newGame.player_order ? newGame : { ...newGame, player_order: [] }
    setGame(newGame)
  }

  // Saving for the future when we have multiple rooms
  const createNewGame = () => {
    axios
      .get(`http://${apiUrl}/new`, axiosConfig)
      .then(function (response: AxiosResponse) {
        if (response.data) {
          setNewGame(response.data)
        }
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

  const joinGame = () => {
    toast.info('Joining game...')
    axios
      .get('http://localhost:3000/v1/game/join', {
        headers: {
          authorization: window.sessionStorage.getItem('token') as string,
        },
      })
      .then((response) => {
        toast.info('Welcome to the game')
      })
      .catch((error: AxiosError) => {
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

  const handleClick = () => {
    axios
      .get(`http://${apiUrl}/status`, axiosConfig)
      .then(function (response: AxiosResponse) {
        if (response.data) {
          if (game.status === 'newGame' || game.status === 'gameover') {
            createNewGame()
          } else {
            joinGame()
          }
        }
      })
      .catch(function (error: AxiosError) {
        createNewGame()
        console.error((error as any).response.data.error)
      })
  }
  const getButtonText = () => {
    return game.status === 'gameover' ? 'New Game' : 'Join Game'
  }

  const isPlayerInGame = () => {
    return Boolean(
      game.cups.filter((cup: user) => {
        return cup.account === account
      }).length,
    )
  }

  return (
    <Button
      disabled={game.status === 'playing' || isPlayerInGame()}
      classes="join__buton"
      clickHandler={() => handleClick()}
    >
      <span>{getButtonText()}</span>
    </Button>
  )
}

export default Join
