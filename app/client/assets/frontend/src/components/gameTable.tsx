import React, { FC } from 'react'
import { claim, user } from '../types/index.d'
import Counter from './counter'
import Cups from './cups'
import CurrentClaim from './currentClaim'
import LiarsCall from './liarsCall'

interface GameTableProps {
  activePlayers: user[]
  currentPlayerAddress: string
  currentClaim: { address: string; claim: claim }
}

const GameTable: FC<GameTableProps> = (GameTableProps) => {
  const { activePlayers, currentPlayerAddress, currentClaim } = GameTableProps
  return (
    <div
      style={{
        display: 'flex',
        width: '100%',
        justifyContent: 'start',
        alignItems: 'center',
        flexDirection: 'column',
      }}
    >
      <Counter />
      <Cups
        activePlayers={activePlayers}
        currentPlayerAddress={currentPlayerAddress}
      />
      <LiarsCall />
      <CurrentClaim currentClaim={currentClaim} />
    </div>
  )
}

export default GameTable
