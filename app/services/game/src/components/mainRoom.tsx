import React, { useEffect, useContext, useState } from 'react'
import GameTable from './gameTable'
import { GameContext } from '../contexts/gameContext'
import useGame from './hooks/useGame'
import useWebSocket from './hooks/useWebSocket'
import { token } from '../utils/axiosConfig'
import { useLocation, useNavigate } from 'react-router-dom'
import AppHeader from './appHeader'
import Footer from './footer'
import { appConfig } from '../types/index.d'
import useEthersConnection from './hooks/useEthersConnection'

// MainRoom component
function MainRoom() {
  // Extracts navigate from useNavigate Hook
  const navigate = useNavigate()

  // Extracts state (a prop send by the router) from useLocation Hook
  const { state } = useLocation()

  // Extracts game from the gameContext using useContext Hook
  let { game } = useContext(GameContext)

  // Extracts account from ethersConnection Hook
  const { account } = useEthersConnection()

  // Extracts addOut and setPlayerDice from useGame Hook
  const { addOut, setPlayerDice } = useGame()

  // Extracts function to connect to ws (connect)
  // and websocket current status from useWebSocket Hook
  const { connect, wsStatus } = useWebSocket(resetTimer)

  // ------------------Timer-------------------
  // Round Interval timer.
  let roundInterval: NodeJS.Timer

  // Timer time in seconds
  const timeoutTime = 30

  // Get the timer that's set inside the sessionStorage
  const sessionTimer = window.sessionStorage.getItem('round_timer')
    ? parseInt(window.sessionStorage.getItem('round_timer') as string) - 1
    : timeoutTime

  // Creates a state to handle the timer
  const [timer, setTimer] = useState(sessionTimer)

  // Sets timer to 0 and removes it from every place is stored at.
  function resetTimer(): void {
    window.sessionStorage.removeItem('round_timer')
    clearInterval(roundInterval)
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
      // eslint-disable-next-line react-hooks/exhaustive-deps
      roundInterval = setInterval(() => {
        if (timer > 0 && game.status === 'playing') {
          setTimer((prevState) => {
            return prevState - 1
          })
        } else {
          addOut()
          resetTimer()
        }
      }, 1000)
    } else {
      clearInterval(roundInterval)
    }
    return () => clearInterval(roundInterval)
  }, [timer, account, game.player_order, game.current_cup, game.status])

  // ---------------Finish timer----------------

  // Connects to websocket depending on status.
  function connectToWs() {
    if (
      wsStatus.current !== 'open' &&
      wsStatus.current !== 'attemptingConnection'
    ) {
      connect()
      wsStatus.current = 'attemptingConnection'
    }
  }
  // First render effect to connect the websocket, clear the round timer and set Player dice if needed.
  useEffect(() => {
    connectToWs()

    // Sets the player dice with the localstore value
    setPlayerDice(
      JSON.parse(window.localStorage.getItem('playerDice') as string),
    )

    // Given that this is the first component that access the game,
    // we set the playerDice item on localStorage with it's zero value.
    if (!window.localStorage.getItem('playerDice')) {
      window.localStorage.setItem('playerDice', JSON.stringify([]))
    }

    // We set the timer with the sessionStorage value.
    // This is to persit the value on refresh.
    setTimer(parseInt(window.sessionStorage.getItem('round_timer') as string))

    // An empty dependecies array triggers useEffect only on the first render of the component
    // We disable the next line so eslint doens't complain about missing dependencies.
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  useEffect(() => {
    // Handles if the user is logged and has a token.
    // If not, we redirect it to the login page. (<Login />)
    function checkAuth() {
      if (!account || !token() || !(state as appConfig)) {
        navigate('/')
      }
    }

    checkAuth()

    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [account, state])

  // Render
  return (
    <div
      className="container-fluid d-flex align-items-center justify-content-start px-0 flex-column"
      style={{ height: '100%', maxHeight: '100vh' }}
    >
      <AppHeader show={true} />
      <div
        style={{
          width: '100%',
          display: 'flex',
          justifyContent: 'start',
          alignItems: 'start',
          maxWidth: '100vw',
          marginTop: '15px',
          height: 'calc(100vh - 181px)',
        }}
        id="mainRoom"
      >
        <GameTable timer={timer} />
      </div>
      <Footer />
    </div>
  )
}

export default MainRoom
