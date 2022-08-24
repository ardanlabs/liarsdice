/* ************useGameHook************

  This hook provides all functions needed to run the game.
  It only exports the methods needed to run the game from another component.
  Apart from the gameflow methods (line 36) we have a timer, and also the gamepot set in here.
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
  let { game, setGame } = useContext(GameContext),
    roundInterval: NodeJS.Timer
  const [playerDice, setPlayerDice] = useState([] as dice)
  const gamePot = useMemo(
    () => game.ante_usd * game.cups.length,
    [game.cups.length, game.ante_usd],
  )
  const setNewGame = (data: game) => {
      const newGame = assureGameType(data)
      setGame(newGame)
    },
    // Timer time in seconds
    timeoutTime = 30,
    // Get the timer that's set inside the sessionStorage
    sessionTimer = window.sessionStorage.getItem('round_timer')
      ? parseInt(window.sessionStorage.getItem('round_timer') as string) - 1
      : timeoutTime
  const [timer, setTimer] = useState(sessionTimer)

  const updateStatus = () => {
    axios
      .get(`http://${apiUrl}/status`, axiosConfig)
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
                game.last_win === account
              ) {
                axios
                  .get(`http://${apiUrl}/reconcile`, axiosConfig)
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

  const managePlayerDice = (dice: dice) => {
    setPlayerDice(dice)
  }

  const rolldice = () => {
    axios
      .get(`http://${apiUrl}/rolldice`, axiosConfig)
      .then(function (response: AxiosResponse) {
        if (response.data) {
          managePlayerDice(response.data.dice)
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
        window.sessionStorage.removeItem('round_timer')
        setTimer(timeoutTime)
        // We calculate how many players are active and if every but one timed out
        // if (getActivePlayersLength(response.data) - 1 === timeoutsCount) {
        //   newRound()
        // } else {
        updateStatus()
        // }
      })
      .catch(function (error: AxiosError) {
        // To be discused
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

  const resetTimer = () => {
    setTimer(timeoutTime)
  }
  // If the timer updates we store it in the sessionStorage in order to persits it when refreshing the page
  useEffect(() => {
    window.sessionStorage.setItem('round_timer', `${timer}`)
  }, [timer])
  // Effect to handle the timer.
  useEffect(() => {
    if (
      (game.player_order as string[])[game.current_cup] === account &&
      game.status === 'playing'
    ) {
      clearInterval(roundInterval)
      // eslint-disable-next-line react-hooks/exhaustive-deps
      roundInterval = setInterval(() => {
        if (timer > 0 && game.status === 'playing') {
          setTimer((prevState) => {
            return prevState - 1
          })
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

  return [
    rolldice,
    timer,
    resetTimer,
    gamePot,
    playerDice,
    managePlayerDice,
    updateStatus,
  ] as const
}

export default useGame
