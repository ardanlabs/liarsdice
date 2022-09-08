import React, { useContext } from 'react'
import SidebarDetails from './sidebarDetails'
import { GameContext } from '../contexts/gameContext'
import NotificationCenter from './notificationCenter/notificationCenter'
import { SideBarProps } from '../types/props.d'

// SideBar component
function SideBar(props: SideBarProps) {
  // Extracts props.
  const { ante, gamePot, notificationCenterWidth } = props

  // Extracts game from useContext hook.
  const { game } = useContext(GameContext)

  // Extracts round from game.
  const { round } = game

  // Renders this markup.
  return (
    <aside
      id="side-bar"
      style={{
        display: 'flex',
        alignItems: 'flex-start',
        flexDirection: 'column',
        border: '1px inset var(--secondary-color)',
        height: 'calc(100% - 165px)',
        right: '0px',
        position: 'fixed',
        top: '95px',
        width: `${notificationCenterWidth}`,
        zIndex: '3',
      }}
    >
      <SidebarDetails ante={ante} round={round} pot={gamePot} />
      <NotificationCenter
        trigger={false}
        asideContainerStyle={{
          display: 'flex',
          flexDirection: 'column',
          height: '100%',
        }}
        mainContainerStyle={{
          height: '100%',
          width: '100%',
          maxHeight: 'calc(100% - 262px)',
          borderTop: '1px inset var(--secondary-color)',
        }}
      />
    </aside>
  )
}

export default SideBar
