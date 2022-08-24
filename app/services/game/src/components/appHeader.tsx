import React, { FC } from 'react'
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
    <header className="App-header">
      <h1>Ardan's Liar's Dice</h1>
      <PlayerBalance />
    </header>
  )
}
export default AppHeader
