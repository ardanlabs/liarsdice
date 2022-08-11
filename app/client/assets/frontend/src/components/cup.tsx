import React, { FC } from 'react'
import { user } from '../types/index.d'
import Dice from './dice'

interface CupProps {
  player: user
  currentPlayerWallet: string
}

const Cup: FC<CupProps> = (CupProps) => {
  const { player, currentPlayerWallet } = CupProps
  return player.outs < 3 ? (
    <div className="player__cup active">
      <Dice
        isPlayerTurn={currentPlayerWallet === player.wallet}
        diceNumber={player.dice}
        playerWallet={player.wallet}
      />
    </div>
  ) : <div className="player__cup"></div>
}

export default Cup
