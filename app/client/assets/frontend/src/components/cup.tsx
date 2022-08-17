import React, { FC, useContext } from 'react'
import { GameContext } from '../gameContext'
import { dice, user } from '../types/index.d'
import Dice from './dice'

interface CupProps {
  player: user
  playerDice: dice
}

const Cup: FC<CupProps> = (CupProps) => {
  const { player, playerDice } = CupProps
  const { game } = useContext(GameContext)
  const { current_cup, player_order, status } = game

  return player.outs < 3 ? (
    <div className="player__cup active">
      <Dice
        isPlayerTurn={
          (player_order as string[])[current_cup] === player.account &&
          status === 'playing'
        }
        diceNumber={
          playerDice && status === 'playing' ? playerDice : [0, 0, 0, 0, 0]
        }
        playerAccount={player.account}
      />
    </div>
  ) : (
    <div className="player__cup"></div>
  )
}

export default Cup
