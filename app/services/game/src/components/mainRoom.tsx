import React, { useEffect, useContext } from 'react'
import SideBar from './sidebar'
import GameTable from './gameTable'
import { GameContext } from '../gameContext'
import useGame from './hooks/useGame'
import useWebSocket from './hooks/useWebSocket'

interface MainRoomProps {}
const MainRoom = (props: MainRoomProps) => {
  let { game } = useContext(GameContext)

  const { timer, gamePot, playerDice, managePlayerDice } = useGame()

  const { connect, wsStatus } = useWebSocket()

  // Effect to persits players dice
  useEffect(() => {
    window.sessionStorage.setItem('playerDice', JSON.stringify(playerDice))
  }, [playerDice])
  // Timer time in seconds

  // Effect to update the state of player's dice
  useEffect(() => {
    managePlayerDice(
      JSON.parse(window.sessionStorage.getItem('playerDice') as string),
    )
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  // First render effect to connect the websocket and to clear the round timer just in case.
  useEffect(() => {
    connect()
    window.sessionStorage.removeItem('round_timer')
    wsStatus.current = 'attemptingConnection'
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
