/* ************useGameHook************

  This hook provides all functions needed to run the game.
  It only exports the methods needed to run the game from another component.
  Apart from the gameflow methods (line 36) we have the gamepot set in here.
  Inside ~/src/components/mainRoom.tsx you can see an implementation of how all of this is working.
  Might be helpfull to see ~/src/hooks/useWebSocket.ts to understand the events that run the game notifications/updating system.
  
  **************************************** */
import { useEffect, useContext, useState, useMemo } from 'react'
import axios, { AxiosError, AxiosResponse } from 'axios'
import { GameContext } from '../../contexts/gameContext'
import { appConfig, dice, die, game, user } from '../../types/index.d'
import assureGameType from '../../utils/assureGameType'
import { apiUrl } from '../../utils/axiosConfig'
import getActivePlayersLength from '../../utils/getActivePlayers'
import useEthersConnection from './useEthersConnection'
import { connectResponse } from '../../types/responses.d'
import { useNavigate } from 'react-router-dom'
import { getAppConfig } from '../..'
import { toast } from 'react-toastify'

// Create an axios instance to keep the token updated
const axiosInstance = axios.create({
  headers: {
    authorization: window.sessionStorage.getItem('token') as string,
  },
})

function useGame() {
  // Extracts account from useEthersConnection hook
  const { account } = useEthersConnection()

  // Extracts game, setGame from useContext hook
  let { game, setGame } = useContext(GameContext)

  // Sets playerDice local state.
  const [playerDice, setPlayerDice] = useState([] as dice)

  // Stores a memoized value of the gamePot that only updates when dependencies change.
  const gamePot = useMemo(
    () => game.anteUSD * game.cups.length,
    [game.cups.length, game.anteUSD],
  )

  // Effect to persits players dice
  useEffect(() => {
    if (playerDice?.length)
      window.localStorage.setItem('playerDice', JSON.stringify(playerDice))
  }, [playerDice])

  // Extracts navigation functionality from useNavigate hook.
  const navigate = useNavigate()

  // ===========================================================================

  // SetNewGame updates the instance of the game in the local state.
  // Also sets the player dice.
  function setNewGame(data: game) {
    const newGame = assureGameType(data)
    if (newGame.cups.length) {
      // We filter the connected player
      const player = newGame.cups.filter((cup) => {
        return cup.account === account
      })
      if (player.length) {
        setPlayerDice(player[0].dice)
      }
    }
    setGame(newGame)
    return newGame
  }

  // updateStatus calls to the status endpoint and updates the local state.
  function updateStatus() {
    // updatesStatusAxiosFn handles the backend answer.
    const updateStatusAxiosFn = (response: AxiosResponse) => {
      if (response.data) {
        const parsedGame = setNewGame(response.data)
        switch (parsedGame.status) {
          case 'newgame':
            if (getActivePlayersLength(parsedGame.cups) >= 2) {
              startGame()
            }
            break
          case 'gameover':
            if (
              getActivePlayersLength(parsedGame.cups) >= 1 &&
              parsedGame.lastWin === account
            ) {
              axiosInstance
                .get(`http://${apiUrl}/reconcile`)
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
    }

    axiosInstance
      .get(`http://${apiUrl}/status`)
      .then(updateStatusAxiosFn)
      .catch(function (error: AxiosError) {
        console.error(error as any)
      })
  }

  // connectToGameEngine connects to the game engine, and stores the token
  // in the sessionStorage. Takes an object with the following type:
  // { dateTime: string; sig: string }
  function connectToGameEngine(data: { dateTime: string; sig: string }) {
    const axiosFn = (connectResponse: connectResponse) => {
      window.sessionStorage.setItem(
        'token',
        `bearer ${connectResponse.data.token}`,
      )
      const getAppConfigFn = (getConfigResponse: appConfig) => {
        window.location.reload()
      }
      getAppConfig.then(getAppConfigFn)
    }

    const axiosErrorFn = (error: AxiosError) => {
      const errorMessage = (error as any).response.data.error.replace(
        / \[.+\]/gm,
        '',
      )

      console.group()
      console.error('Error:', errorMessage)
      console.groupEnd()
    }

    axiosInstance
      .post(`http://${apiUrl}/connect`, data)
      .then(axiosFn)
      .catch(axiosErrorFn)
  }

  // ===========================================================================

  // Game flow methods

  // rolldice rolls the player dice.
  function rolldice(): void {
    axiosInstance
      .get(`http://${apiUrl}/rolldice`)
      .then(function (response: AxiosResponse) {
        toast(`Rolling dice's`)
      })
      .catch(function (error: AxiosError) {
        console.error(error)
      })
  }

  // startGame starts the game
  function startGame() {
    axiosInstance
      .get(`http://${apiUrl}/start`)
      .then(function () {})
      .catch(function (error: AxiosError) {
        console.error(error)
      })
  }

  // nextTurn calls to nextTurn and then updates the status.
  function nextTurn() {
    axiosInstance
      .get(`http://${apiUrl}/next`)
      .then(function (response: AxiosResponse) {
        updateStatus()
      })
      .catch(function (error: AxiosError) {
        console.error(error)
      })
  }

  // Takes an account address and adds an out to that account
  function addOut(accountToOut = game.currentID) {
    const player = game.cups.filter((player: user) => {
      return player.account === accountToOut
    })

    const addOutAxiosFn = (response: AxiosResponse) => {
      setNewGame(response.data)
      // If the game didn't stop we call next-turn
      if (response.data.status === 'playing') {
        nextTurn()
      }
    }

    axiosInstance
      .get(`http://${apiUrl}/out/${player[0].outs + 1}`)
      .then(addOutAxiosFn)
      .catch(function (error: AxiosError) {
        console.group('Something went wrong, try again.')
        console.error(error)
        console.groupEnd()
      })
  }

  // sendBet sends the player bet to the backend.
  function sendBet(number: number, suite: die) {
    axiosInstance
      .get(`http://${apiUrl}/bet/${number}/${suite}`)
      .then()
      .catch(function (error: AxiosError) {
        console.error(error)
      })
  }

  // callLiar triggers the callLiar endpoint
  function callLiar() {
    axiosInstance
      .get(`http://${apiUrl}/liar`)
      .then(function (response: AxiosResponse) {})
      .catch(function (error: AxiosError) {
        console.error(error)
      })
  }

  return {
    addOut,
    gamePot,
    playerDice,
    setPlayerDice,
    updateStatus,
    sendBet,
    callLiar,
    rolldice,
    connectToGameEngine,
  }
}

export default useGame
