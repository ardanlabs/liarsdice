import React, { useContext } from 'react'
import { GameContext } from '../contexts/gameContext'
import { AppHeaderProps } from '../types/props.d'
import useEthersConnection from './hooks/useEthersConnection'
import PlayerBalance from './playerBalance'

// Header for the aplication, shows the title and the playerBalance component
function AppHeader(AppHeaderProps: AppHeaderProps) {
  // Prop that controls if the header is shown
  const { show } = AppHeaderProps

  // Extracts game from useContext hook.
  const { game } = useContext(GameContext)

  // Deconstruct properties from game.
  const { currentID, status } = game

  // Extracts account from useEthersConnection hook.
  const { account } = useEthersConnection()

  const isPlayerTurn = account === currentID && status === 'playing'

  // Renders if show is true.
  return show ? (
    <header
      data-testid="app-header"
      className={isPlayerTurn ? 'active app-header' : 'app-header'}
    >
      <h1>Ardan's Liar's Dice</h1>
      <PlayerBalance />
    </header>
  ) : null
}
export default AppHeader
