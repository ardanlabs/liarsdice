import React from 'react'
import { user } from '../types/index.d'
import SidebarDetails from './sidebarDetails'
import Players from './players'

interface MainRoomProps {
  activePlayers: Set<user>,
  waitingPlayers?: string[],
  joinGame: Function,
  currentGameStatus: any,
}
const MainRoom = (props: MainRoomProps) => {
  const { activePlayers, waitingPlayers, joinGame, currentGameStatus } = props
  const { round, current_player } = currentGameStatus

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
      <SidebarDetails round={round} ante={5} pot={10} diceAmount={activePlayers.size} />
      <Players activePlayers={activePlayers} currentPlayer={current_player} waitingPlayers={waitingPlayers} joinGame={joinGame}/>
    </div>
  )
}

export default MainRoom
