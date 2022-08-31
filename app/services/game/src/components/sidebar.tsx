import React, { useContext } from 'react'
import SidebarDetails from './sidebarDetails'
import { GameContext } from '../gameContext'
import NotificationCenter from './notificationCenter/notificationCenter'

interface SideBarProps {
  ante: number
  gamePot: number
  notificationCenterWidth: string
}
const SideBar = (props: SideBarProps) => {
  const { ante, gamePot, notificationCenterWidth } = props
  const { game } = useContext(GameContext)
  const { round } = game

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
