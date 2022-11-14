import React, { useContext } from 'react'
import { GameContext } from '../contexts/gameContext'
import { CupProps } from '../types/index.d'
import Dice from './dice'

// Cup component
function Cup(props: CupProps) {
  // Extracts props.
  const { player, playerDice } = props

  // Extracts game from useContext hook.
  const { game } = useContext(GameContext)

  // Extracts properties from game.
  const { currentID, status } = game

  // Renders the cup depending on player status.
  return player.outs < 3 ? (
    <div data-testid="player__cup" className="player__cup active">
      <Dice
        isPlayerTurn={currentID === player.account && status === 'playing'}
        diceNumber={
          playerDice && status === 'playing' ? playerDice : [0, 0, 0, 0, 0]
        }
      />
    </div>
  ) : (
    <div data-testid="player__cup" className="player__cup"></div>
  )
}

export default Cup
