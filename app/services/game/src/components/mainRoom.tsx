import React, { useEffect, useContext } from 'react'
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
import { useTimer } from 'react-timer-hook'
import SideBar from './sidebar'

// MainRoom component
function MainRoom() {
  // Extracts navigate from useNavigate Hook.
  const navigate = useNavigate()

  // Extracts state (a prop send by the router) from useLocation Hook.
  const { state, key } = useLocation()

  // Extracts game from the gameContext using useContext Hook.
  let { game } = useContext(GameContext)

  // Extracts account from ethersConnection Hook.
  const { account } = useEthersConnection()

  // Extracts addOut and setPlayerDice from useGame Hook.
  const { addOut, setPlayerDice } = useGame()

  // Extracts function to connect to ws (connect) from useWebSocket Hook.
  const { connect } = useWebSocket(restartTimer)

  // Variable to set the notification center width.
  const notificationCenterWidth = '340px'

  const wsStatus = window.sessionStorage.getItem('wsStatus')

  // ===========================================================================

  // initUEFn connects the websocket, clears the round timer and
  // sets Player dice if needed.
  const initUEFn = () => {
    // Connects to websocket depending on status.
    function connectToWs() {
      connect().then(() => {
        window.sessionStorage.setItem('wsStatus', 'open')
      })
    }
    if (wsStatus !== 'open' && wsStatus !== 'attemptingConnection') {
      window.sessionStorage.setItem('wsStatus', 'attemptingConnection')
      connectToWs()
    }
    window.sessionStorage.setItem('wsStatus', 'close')
    // Sets the player dice with the localstore value
    setPlayerDice(
      JSON.parse(window.localStorage.getItem('playerDice') as string),
    )

    // Given that this is the first component that access the game,
    // we set the playerDice item on localStorage with it's zero value.
    if (!window.localStorage.getItem('playerDice')) {
      window.localStorage.setItem('playerDice', JSON.stringify([]))
    }
  }

  // An empty dependecies array triggers useEffect only on the first render of the component
  // We disable the next line so eslint doens't complain about missing dependencies.
  // eslint-disable-next-line
  useEffect(initUEFn, [])

  // ===========================================================================

  const authUEFn = () => {
    // Handles if the user is logged and has a token.
    // If not, we redirect it to the login page. (<Login />)
    function checkAuth() {
      if (!account || !token() || !(state as appConfig)) {
        navigate('/')
      }
    }

    checkAuth()
  }

  // eslint-disable-next-line
  useEffect(authUEFn, [account, state])

  // ===============================Timer=======================================

  // Timer duration in seconds
  const timerDuration = 30

  // getTimeoutTime returns a Date object with with seconds appart from now
  function getTimeoutTime(seconds: number) {
    const now = new Date()
    now.setSeconds(now.getSeconds() + seconds)
    return now
  }

  // Gets the timer that's set inside the sessionStorage
  // If not timer is set it returns the default time.
  const sessionTimer =
    window.sessionStorage.getItem('round_timer') !== '0'
      ? parseInt(window.sessionStorage.getItem('round_timer') as string)
      : timerDuration

  // restartTimer restarts the timer to 30 seconds.
  function restartTimer() {
    restart(getTimeoutTime(timerDuration), false)
  }

  // ===========================================================================
  const timerExpiredFn = () => {
    setTimeout(() => {
      pause()
      restartTimer()
      addOut()
    }, 500)
  }

  const { seconds, start, restart, pause, isRunning } = useTimer({
    expiryTimestamp: getTimeoutTime(sessionTimer),
    autoStart: false,
    onExpire: timerExpiredFn,
  })

  // ===========================================================================

  // If the timer updates we store it in the sessionStorage persist it when
  // refreshing the page.
  useEffect(() => {
    window.sessionStorage.setItem('round_timer', `${seconds}`)
  }, [seconds])

  // timerUEFn starts the timer if the conditions are met.
  const timerUEFn = () => {
    const isGamePlaying = game.status === 'playing'
    const isPlayerTurn = game.currentID === account

    if (isPlayerTurn && isGamePlaying && !isRunning) {
      start()
      return
    }
  }

  useEffect(timerUEFn, [game, account, isRunning, start])

  //
  // ===============================Finish Timer================================

  // Renders this final markup
  return (
    <div
      className="d-flex align-items-center justify-content-start px-0 flex-column"
      style={{ height: '100%', maxHeight: '100vh' }}
    >
      <AppHeader show={true} />
      <div className="d-flex" style={{ width: '100vw' }}>
        <section
          style={{
            width: `calc(100% - ${notificationCenterWidth})`,
            zIndex: '1',
          }}
          className="d-flex flex-column align-items-center justify-content-start"
        >
          <GameTable timer={seconds} />
          <Footer />
        </section>
        <SideBar notificationCenterWidth={notificationCenterWidth} />
      </div>
    </div>
  )
}

export default MainRoom
