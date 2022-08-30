import React, { FC } from 'react'
import { bet } from '../types/index.d'
import Die from './icons/die'

interface BetProps {
  bet: bet
  dieWidth?: string
  dieHeight?: string
  fill: string
}

const Bet: FC<BetProps> = (BetProps) => {
  const { bet, dieWidth, dieHeight, fill } = BetProps

  return bet ? (
    <>
      {`${bet.number} X `}
      <Die
        dieNumber={bet.suite}
        fill={fill}
        width={dieWidth}
        height={dieHeight}
        style={{ marginLeft: '20px' }}
      />
    </>
  ) : null
}

export default Bet
