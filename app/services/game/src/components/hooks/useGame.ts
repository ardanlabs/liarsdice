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
import { axiosConfig, apiUrl } from '../../utils/axiosConfig'
import getActivePlayersLength from '../../utils/getActivePlayers'
import useEthersConnection from './useEthersConnection'
import { toast } from 'react-toastify'
import { shortenIfAddress } from '../../utils/address'
import { connectResponse } from '../../types/responses.d'
import { useNavigate } from 'react-router-dom'
import { getAppConfig } from '../..'

const useGame = () => {
  const { account } = useEthersConnection()
  let { game, setGame } = useContext(GameContext)
  const [playerDice, setPlayerDice] = useState([] as dice)
  const gamePot = useMemo(
    () => game.anteUsd * game.cups.length,
    [game.cups.length, game.anteUsd],
  )
  const navigate = useNavigate()
  const setNewGame = (data: game) => {
    const newGame = assureGameType(data)
    if (newGame.cups.length) {
      setPlayerDice(newGame.cups[0].dice)
    }
    setGame(newGame)
  }

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
              if (getActivePlayersLength(response.data.playerOrder) >= 2) {
                startGame()
              }
              break
            case 'gameover':
              if (
                getActivePlayersLength(response.data.playerOrder) === 1 &&
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
        console.error(error as any)
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
    if (playerDice?.length)
      window.localStorage.setItem('playerDice', JSON.stringify(playerDice))
  }, [playerDice])

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
        navigate('/mainRoom', { state: { ...getConfigResponse } })
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

    axios
      .post('http://localhost:3000/v1/game/connect', data)
      .then(axiosFn)
      .catch(axiosErrorFn)
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
    accountToOut = (game.playerOrder as string[])[game.currentCup],
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

  function sendBet(number: number, suite: die) {
    axios
      .get(`http://${apiUrl}/bet/${number}/${suite}`, axiosConfig)
      .then(function (response: AxiosResponse) {
        if (response.data) {
          setNewGame(response.data)
        }
      })
      .catch(function (error: AxiosError) {
        console.error(error)
      })
  }
  function callLiar() {
    axios
      .get(`http://${apiUrl}/liar`, axiosConfig)
      .then(function (response: AxiosResponse) {
        if (getActivePlayersLength(game.playerOrder) === 1) {
          toast(
            `Game finished! Winner is ${shortenIfAddress(
              response.data.cups[0].account,
            )}`,
          )
        }
      })
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
    connectToGameEngine,
  }
}

export default useGame
