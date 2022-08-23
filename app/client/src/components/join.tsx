import React, { useContext } from 'react'
import { utils } from 'ethers'
import { useEthers } from '@usedapp/core'
import axios, { AxiosError, AxiosResponse } from 'axios'
import Button from './button'
import { toast } from 'react-toastify'
import { capitalize } from '../utils/capitalize'
import { axiosConfig } from '../utils/axiosConfig'
import { GameContext } from '../gameContext'
import { game } from '../types/index.d'

interface JoinProps {
  disabled: boolean
}

const Join = (props: JoinProps) => {
  const { disabled } = props
  const { game, setGame } = useContext(GameContext)
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
  const createNewGame = (ante: number = 10) => {
    axios
      .get(`http://${apiUrl}/new`)
      .then(function (response: AxiosResponse) {
        if (response.data) {
          setNewGame(response.data)
          joinGame()
        }
      })
      .catch(function (error: AxiosError) {
        console.error(error)
      })
  }

  const joinGame = () => {
    toast.info('Joining game...')
    axios
      .get('http://localhost:3000/v1/game/join', axiosConfig)
      .then((response) => {
        toast.info('Welcome to the game')
        window.sessionStorage.setItem('token', `bearer ${response.data.token}`)
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
        createNewGame()
      })
  }
  return (
    <Button
      disabled={disabled}
      classes="join__buton"
      clickHandler={() => joinGame()}
    >
      <span>{game ? 'Join Game' : 'New Game'}</span>
    </Button>
  )
}

export default Join
