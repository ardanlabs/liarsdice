import React from 'react'
import { user } from '../types/index.d'
import SidebarDetails from './sidebarDetails'
import Players from './players'

interface MainRoomProps {
  activePlayers: user[],
  waitingPlayers: user[],
}
const MainRoom = (props: MainRoomProps) => {
  const { activePlayers, waitingPlayers } = props

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
      <SidebarDetails />
      <Players activePlayers={activePlayers} waitingPlayers={waitingPlayers}/>
    </div>
  )
}

export default MainRoom
