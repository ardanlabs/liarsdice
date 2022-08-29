import React, { FC } from 'react'
import NotificationCenter from './notificationCenter/notificationCenter'
import PlayerBalance from './playerBalance'

interface AppHeaderProps {
  show?: boolean
}

const AppHeader: FC<AppHeaderProps> = (AppHeaderProps) => {
  const { show } = AppHeaderProps
  if (!show) {
    return null
  }
  return (
    <header data-testid="app-header" className="App-header">
      <NotificationCenter />
      <h1>Ardan's Liar's Dice</h1>
      <PlayerBalance />
    </header>
  )
}
export default AppHeader
