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

interface MainRoomProps {}
const MainRoom = (props: MainRoomProps) => {
  const navigate = useNavigate()
  const { state } = useLocation()
  let { game } = useContext(GameContext),
    roundInterval: NodeJS.Timer

  const // Timer time in seconds
    timeoutTime = 30,
    // Get the timer that's set inside the sessionStorage
    sessionTimer = window.sessionStorage.getItem('round_timer')
      ? parseInt(window.sessionStorage.getItem('round_timer') as string) - 1
      : timeoutTime,
    [timer, setTimer] = useState(sessionTimer),
    { account } = useEthersConnection(),
    { playerDice, setPlayerDice, addOut } = useGame()

  const resetTimer = () => {
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

  const { connect, wsStatus } = useWebSocket(resetTimer)

  // First render effect to connect the websocket, clear the round timer and set Player dice if needed.
  useEffect(() => {
    connect()
    wsStatus.current = 'attemptingConnection'
    setPlayerDice(
      JSON.parse(window.localStorage.getItem('playerDice') as string),
    )
    setTimer(parseInt(window.sessionStorage.getItem('round_timer') as string))
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  useEffect(() => {
    if (!account || !token() || !(state as appConfig).config) {
      navigate('/')
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [account, state])

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
        <GameTable playerDice={playerDice} timer={timer} />
      </div>
      <Footer />
    </div>
  )
}

export default MainRoom
