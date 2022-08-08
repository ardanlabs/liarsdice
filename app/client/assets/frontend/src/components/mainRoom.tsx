import React, { useState } from 'react'
import SideBar from './sidebar'
import { user } from '../../types/index.d'

interface MainRoomProps {
}
const MainRoom = (props: MainRoomProps) => {
  const activePlayers: user[] = []
  const waitingPlayers: user[] = []

  return (
    <div
      style={{
        height: '100%',
        width: '100%',
        display: 'flex',
        justifyContent: 'start',
        alignItems: 'center',
      }}
    >
      <SideBar activePlayers={activePlayers} waitingPlayers={waitingPlayers} />
      <div
        style={{
          height: '100%',
          width: 'fit-content',
          backgroundColor: 'var(--primary-color)',
          display: 'flex',
          justifyContent: 'start',
          alignItems: 'center',
        }}
      >
      </div>
    </div>
  )
}

export default MainRoom
