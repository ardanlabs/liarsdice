import React, { useContext } from 'react'
import SidebarDetails from './sidebarDetails'
import Players from './players'
import { GameContext } from '../gameContext'

interface SideBarProps {
  ante: number
  gamePot: number
}
const SideBar = (props: SideBarProps) => {
  const { ante, gamePot } = props
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
      <SidebarDetails ante={ante} round={round} pot={gamePot} />
      <Players />
    </div>
  )
}

export default SideBar
