import React, { FC } from 'react'
import Counter from './counter'
import Cups from './cups'
import CurrentClaim from './currentClaim'
import LiarsCall from './liarsCall'

interface GameTableProps {}

const GameTable: FC<GameTableProps> = (GameTableProps) => {

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
      <Cups />
      <LiarsCall />
      <CurrentClaim currentClaim={{wallet: '', claim: {number: 1, suite: 2}}} />
    </div>
  )
}

export default GameTable
