import React, { FC, useContext } from 'react'
import { GameContext } from '../gameContext'
import { user } from '../types/index.d'
import Dice from './dice'

interface CupProps {
  player: user
}

const Cup: FC<CupProps> = (CupProps) => {
  const { player } = CupProps
  const { game } = useContext(GameContext)
  const { current_player } = game

  return player.outs < 3 ? (
    <div className="player__cup active">
      <Dice
        isPlayerTurn={current_player === player.account}
        diceNumber={player.dice ? player.dice : [0, 0, 0, 0, 0]}
        playerAccount={player.account}
      />
    </div>
  ) : (
    <div className="player__cup"></div>
  )
}

export default Cup
