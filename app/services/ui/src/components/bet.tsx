import React from 'react'
import { BetProps } from '../types/props.d'
import Die from './icons/die'

// Bet component
function Bet(props: BetProps) {
  // Extracts props
  const { bet, dieWidth, dieHeight, fill } = props

  // If there's a bet renders this markup
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
