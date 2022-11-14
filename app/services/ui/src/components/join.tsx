import React, { useContext } from 'react'
import axios, { AxiosError, AxiosResponse } from 'axios'
import Button from './button'
import { apiUrl, axiosConfig } from '../utils/axiosConfig'
import { GameContext } from '../contexts/gameContext'
import { user } from '../types/index.d'
import useEthersConnection from './hooks/useEthersConnection'
import { JoinProps } from '../types/props.d'
import useGame from './hooks/useGame'

// Join component
function Join(props: JoinProps) {
  // Extracts props.
  const { disabled } = props

  // Extracts game and setGame from useContext hook.
  const { game } = useContext(GameContext)

  // Extracts joinGame and createNewGame from useGame hook.
  const { joinGame, createNewGame } = useGame()

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
