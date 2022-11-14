import React from 'react'
import SidebarDetails from './sidebarDetails'
import NotificationCenter from './notificationCenter/notificationCenter'
import { SideBarProps } from '../types/props.d'

// SideBar component
function SideBar(props: SideBarProps) {
  // Extracts props.
  const { notificationCenterWidth } = props

  // Renders this markup.
  return (
    <aside
      id="side-bar"
      style={{
        display: 'flex',
        alignItems: 'flex-start',
        flexDirection: 'column',
        border: '1px inset var(--secondary-color)',
        right: '0px',
        flexGrow: '1',
        top: '95px',
        width: `${notificationCenterWidth}`,
      }}
    >
      <SidebarDetails />
      <NotificationCenter
        notificationCenterWidth={notificationCenterWidth}
        trigger={false}
        asideContainerStyle={{
          display: 'flex',
          flexDirection: 'column',
          height: 'calc(100% - 43px)',
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
