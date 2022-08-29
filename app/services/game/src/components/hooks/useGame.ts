/* ************useGameHook************

  This hook provides all functions needed to run the game.
  It only exports the methods needed to run the game from another component.
  Apart from the gameflow methods (line 36) we have the gamepot set in here.
  Inside ~/src/components/mainRoom.tsx you can see an implementation of how all of this is working.
  Might be helpfull to see ~/src/hooks/useWebSocket.ts to understand the events that run the game notifications/updating system.
  
  **************************************** */
import { useEffect, useContext, useState, useMemo } from 'react'
import axios, { AxiosError, AxiosResponse } from 'axios'
import { GameContext } from '../../gameContext'
import { useEthers } from '@usedapp/core'
import { dice, game, user } from '../../types/index.d'
import assureGameType from '../../utils/assureGameType'
import { axiosConfig, apiUrl } from '../../utils/axiosConfig'
import getActivePlayersLength from '../../utils/getActivePlayers'

const useGame = () => {
  const { account } = useEthers()
  let { game, setGame } = useContext(GameContext)
  const [playerDice, setPlayerDice] = useState([] as dice)
  const gamePot = useMemo(
    () => game.ante_usd * game.cups.length,
    [game.cups.length, game.ante_usd],
  )
  const setNewGame = (data: game) => {
    const newGame = assureGameType(data)
    setGame(newGame)
  }

  const updateStatus = () => {
    axios
      .get(`http://${apiUrl}/status`)
      .then(function (response: AxiosResponse) {
        if (response.data) {
          setNewGame(response.data)
          switch (response.data.status) {
            // Keeping in case of edge cases
            // case 'roundover':
            //   newRound()
            //   break
            case 'newgame':
              if (getActivePlayersLength(response.data) >= 2) {
                startGame()
              }
              break
            case 'gameover':
              if (
                getActivePlayersLength(response.data) === 1 &&
                response.data.last_win === account
              ) {
                axios
                  .get(`http://${apiUrl}/reconcile`, axiosConfig)
                  .then((response: AxiosResponse) => {
                    updateStatus()
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
        console.error((error as any).response.data.error)
      })
  }

  // Game flow methods

  const startGame = () => {
    if (game.status === 'gameover') {
      axios
        .get(`http://${apiUrl}/start`, axiosConfig)
        .then(function () {})
        .catch(function (error: AxiosError) {
          console.error(error)
        })
    }
  }

  // Effect to persits players dice
  useEffect(() => {
    if (playerDice.length) {
      window.localStorage.setItem('playerDice', JSON.stringify(playerDice))
    }
  }, [playerDice])

  const rolldice = () => {
    axios
      .get(`http://${apiUrl}/rolldice`, axiosConfig)
      .then(function (response: AxiosResponse) {
        if (response.data) {
          setPlayerDice(response.data.dice)
        }
      })
      .catch(function (error: AxiosError) {
        console.error(error)
      })
  }

  const nextTurn = () => {
    axios
      .get(`http://${apiUrl}/next`, axiosConfig)
      .then(function (response: AxiosResponse) {
        updateStatus()
      })
      .catch(function (error: AxiosError) {
        console.error(error)
      })
  }
  // Takes an account address and adds an out to that account
  const addOut = (
    accountToOut = (game.player_order as string[])[game.current_cup],
  ) => {
    const player = game.cups.filter((player: user) => {
      return player.account === accountToOut
    })
    axios
      .get(`http://${apiUrl}/out/${player[0].outs + 1}`, axiosConfig)
      .then(function (response: AxiosResponse) {
        setNewGame(response.data)
        // If the game didn't stop we call next-turn
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

  return {
    addOut,
    rolldice,
    gamePot,
    playerDice,
    setPlayerDice,
    updateStatus,
  }
}

export default useGame
