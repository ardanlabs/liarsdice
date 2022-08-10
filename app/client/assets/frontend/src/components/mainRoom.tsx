import React from 'react'
import { user, claim } from '../types/index.d'
import SideBar from './sidebar'
import GameTable from './gameTable'

interface MainRoomProps {}
const MainRoom = (props: MainRoomProps) => {
  const currentClaim: { address: string; claim: claim } = {
    address: '0x39249126d90671284cd06495d19C04DD0e54d33',
    claim: { number: 1, suite: 4 },
  }
  const currentPlayerAddress: string =
    '0x39249126d90671284cd06495d19C04DD0e54d33'
  const activePlayers: user[] = [
    {
      address: '0x39249126d90671284cd06495d19C04DD0e54d33',
      active: true,
      dice: [1, 2, 3, 4, 5],
      out: 3,
      claim: { number: 1, suite: 4 },
    },
    {
      address: '0x39249126d90671284cd06495d19C04DD0e54d371',
      active: true,
      dice: [1, 2, 3, 4, 5],
      out: 2,
      claim: { number: 2, suite: 5 },
    },
  ]
  const waitingPlayers: user[] = []

  return (
    <div
      style={{
        width: '100%',
        display: 'flex',
        justifyContent: 'start',
        alignItems: 'center',
        maxWidth: '100vw',
      }}
      id="mainRoom"
    >
      <SideBar activePlayers={activePlayers} waitingPlayers={waitingPlayers} />
      <GameTable activePlayers={activePlayers} currentPlayerAddress={currentPlayerAddress} currentClaim={currentClaim} />
    </div>
  )
}

export default MainRoom
