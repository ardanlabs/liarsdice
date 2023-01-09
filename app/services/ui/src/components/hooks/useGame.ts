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
import { dice, die, game, user } from '../../types/index.d'
import assureGameType from '../../utils/assureGameType'
import { apiUrl } from '../../utils/axiosConfig'
import getActivePlayersLength from '../../utils/getActivePlayers'
import useEthersConnection from './useEthersConnection'
import { connectResponse, defaultApiError } from '../../types/responses.d'
import { getAppConfig } from '../..'
import { toast } from 'react-toastify'
import { capitalize } from '../../utils/capitalize'
import { useNavigate } from 'react-router-dom'

// Create an axios instance to keep the token updated
const axiosInstance = axios.create({
  headers: {
    authorization: window.sessionStorage.getItem('token') as string,
  },
})

function useGame() {
  const navigate = useNavigate()

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
                .then(() => {
                  updateStatus()
                })
                .catch((error: AxiosError) => {
                  console.error(error)
                })
            }
            break
          case 'nogame':
            window.localStorage.removeItem('playerDice')
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
  function connectToGameEngine(data: {
    address: string
    dateTime: string
    sig: string
  }) {
    const getAppConfigFn = () => {
      navigate('/mainroom')
    }
    if (window.sessionStorage.getItem('token')) {
      window.localStorage.setItem('account', data.address)
      getAppConfig.then(getAppConfigFn)
    }
    const axiosConnectFn = (connectResponse: connectResponse) => {
      window.sessionStorage.setItem(
        'token',
        `Bearer ${connectResponse.data.token}`,
      )

      window.localStorage.setItem('account', data.address)

      getAppConfig.then(getAppConfigFn)
    }

    const axiosConnectErrorFn = (error: AxiosError) => {
      const errorMessage = (error as any).response.data.error.replace(
        / \[.+\]/gm,
        '',
      )

      console.group()
      console.error('Error:', errorMessage)
      console.groupEnd()
    }

    axiosInstance
      .post(`http://${apiUrl}/connect`, { ...data })
      .then(axiosConnectFn)
      .catch(axiosConnectErrorFn)
  }

  // ===========================================================================

  // Game flow methods

  // joinGame calls to backend join endpoint.
  function joinGame() {
    toast.info('Joining game...')

    // catchFn catches the error
    const catchFn = (error: defaultApiError) => {
      const errorMessage = error.response.data.error.replace(/\[[^\]]+\]/gm, '')

      console.log(errorMessage.replace(/\[[^\]]+\]/gm, ''))

      toast(capitalize(errorMessage))
      console.group()
      console.error('Error:', error.response.data.error)
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

  function createNewGame() {
    // Sets a new game in the gameContext.
    const createGameFn = (response: AxiosResponse) => {
      if (response.data) {
        const newGame = assureGameType(response.data)
        setGame(newGame)
      }
    }

    // Catches the error from the axios call.
    const createGameCatchFn = (error: defaultApiError) => {
      // Figure out regex
      console.log(error.response.data.error)

      let errorMessage = error.response.data.error.replace(/\[[^\]]+\]/gm, '')
      toast(capitalize(errorMessage))
      console.group()
      console.error('Error:', error.response.data.error)
      console.groupEnd()
    }

    axiosInstance
      .get(`http://${apiUrl}/new`)
      .then(createGameFn)
      .catch(createGameCatchFn)
  }

  // rolldice rolls the player dice.
  function rolldice(): void {
    toast(`Rolling dice's`)
    axiosInstance
      .get(`http://${apiUrl}/rolldice`)
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
      .then(function () {
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
    createNewGame,
    joinGame,
  }
}

export default useGame
