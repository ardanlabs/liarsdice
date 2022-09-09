import React, { useContext } from 'react'
import axios, { AxiosError, AxiosResponse } from 'axios'
import Button from './button'
import { toast } from 'react-toastify'
import { capitalize } from '../utils/capitalize'
import { apiUrl, axiosConfig } from '../utils/axiosConfig'
import { GameContext } from '../contexts/gameContext'
import { user } from '../types/index.d'
import useEthersConnection from './hooks/useEthersConnection'
import assureGameType from '../utils/assureGameType'
import { JoinProps } from '../types/props.d'

// Join component
function Join(props: JoinProps) {
  // Extracts props.
  const { disabled } = props

  // Extracts game and setGame from useContext hook.
  const { game, setGame } = useContext(GameContext)

  // Extracts connected account from useEthersConnection hook.
  const { account } = useEthersConnection()

  // Checks if player is in game.
  function isPlayerInGame() {
    return Boolean(
      game.cups.filter((cup: user) => {
        return cup.account === account
      }).length,
    )
  }

  // Checks if button is disabled
  const isButtonDisabled =
    game.status === 'playing' ||
    (game.status === 'newgame' && isPlayerInGame()) ||
    disabled

  // ===========================================================================

  function createNewGame() {
    // Sets a new game in the gameContext.
    const createGameFn = (response: AxiosResponse) => {
      if (response.data) {
        const newGame = assureGameType(response.data)
        setGame(newGame)
      }
    }

    // Catches the error from the axios call.
    const createGameCatchFn = (error: AxiosError) => {
      let errorMessage = (error as any).response.data.error.replace(
        / \[.+\]/gm,
        '',
      )
      toast(
        <div style={{ textAlign: 'start' }}>{capitalize(errorMessage)}</div>,
      )
      console.group()
      console.error('Error:', (error as any).response.data.error)
      console.groupEnd()
    }

    axios
      .get(`http://${apiUrl}/new`, axiosConfig)
      .then(createGameFn)
      .catch(createGameCatchFn)
  }

  // joinGame calls to backend join endpoint.
  function joinGame() {
    toast.info('Joining game...')

    // catchFn catches the error
    const catchFn = (error: AxiosError) => {
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
    }

    axios
      .get(`http://${apiUrl}/join`, {
        headers: {
          authorization: window.sessionStorage.getItem('token') as string,
        },
      })
      .then(() => {
        toast.info('Welcome to the game')
      })
      .catch(catchFn)
  }

  // ===========================================================================
  function handleClick() {
    const handleClickAxiosFn = (response: AxiosResponse) => {
      if (
        response.data &&
        (game.status === 'nogame' || game.status === 'reconciled')
      ) {
        createNewGame()
        return
      }
      joinGame()
    }

    const handleClickAxiosErrorFn = (error: AxiosError) => {
      createNewGame()
      console.error((error as any).response.data.error)
    }

    axios
      .get(`http://${apiUrl}/status`, axiosConfig)
      .then(handleClickAxiosFn)
      .catch(handleClickAxiosErrorFn)
  }

  // Renders this markup
  return (
    <Button
      disabled={isButtonDisabled}
      classes="join__buton"
      clickHandler={() => handleClick()}
      style={{
        backgroundColor: `${
          isButtonDisabled ? 'grey' : 'var(--primary-color)'
        }`,
      }}
    >
      <span>
        {game.status === 'nogame' || game.status === 'reconciled'
          ? 'New Game'
          : 'Join Game'}
      </span>
    </Button>
  )
}

export default Join
