import React, { useContext } from 'react'
import SidebarDetails from './sidebarDetails'
import Players from './players'
import { GameContext } from '../gameContext'

interface MainRoomProps {
  joinGame: Function
}
const MainRoom = (props: MainRoomProps) => {
  const { joinGame } = props
  const { game } = useContext(GameContext)
  const { round } = game

  return (
    <div
      id="side-bar"
      style={{
        display: 'flex',
        alignItems: 'start',
        flexDirection: 'column',
        height: '100%',
      }}
    >
      <SidebarDetails round={round} />
      <Players joinGame={joinGame} />
    </div>
  )
}

export default MainRoom
