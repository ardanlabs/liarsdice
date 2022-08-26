import React, { useEffect, useContext, useState } from 'react'
import SideBar from './sidebar'
import GameTable from './gameTable'
import { GameContext } from '../gameContext'
import useGame from './hooks/useGame'
import useWebSocket from './hooks/useWebSocket'
import { useEthers } from '@usedapp/core'

interface MainRoomProps {}
const MainRoom = (props: MainRoomProps) => {
  let { game } = useContext(GameContext),
    roundInterval: NodeJS.Timer

  const // Timer time in seconds
    timeoutTime = 30,
    // Get the timer that's set inside the sessionStorage
    sessionTimer = window.sessionStorage.getItem('round_timer')
      ? parseInt(window.sessionStorage.getItem('round_timer') as string) - 1
      : timeoutTime,
    [timer, setTimer] = useState(sessionTimer),
    { account } = useEthers(),
    { gamePot, playerDice, setPlayerDice, addOut } = useGame()

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

  return (
    <div
      style={{
        width: '100%',
        display: 'flex',
        justifyContent: 'start',
        alignItems: 'center',
        maxWidth: '100vw',
        marginTop: '15px',
      }}
      id="mainRoom"
    >
      <SideBar ante={game.ante_usd} gamePot={gamePot} />
      <GameTable playerDice={playerDice} timer={timer} />
    </div>
  )
}

export default MainRoom
