import React from 'react'
import { AppHeaderProps } from '../types/props.d'
import PlayerBalance from './playerBalance'

// Header for the aplication, shows the title and the playerBalance component
function AppHeader(AppHeaderProps: AppHeaderProps) {
  // Prop that controls if the header is shown
  const { show } = AppHeaderProps

  // Renders if show is true.
  return show ? (
    <header data-testid="app-header" className="App-header">
      <h1>Ardan's Liar's Dice</h1>
      <PlayerBalance />
    </header>
  ) : null
}
export default AppHeader
