import React, { FC } from 'react'
import { claim } from '../types/index.d'
import Die from './icons/die'

interface ClaimProps {
  claim: claim
  dieWidth?: string
  dieHeight?: string
  fill: string
}

const Claim: FC<ClaimProps> = (ClaimProps) => {
  const { claim, dieWidth, dieHeight, fill } = ClaimProps

  return claim ? (
    <>
      {`${claim.number} X `}
      <Die
        dieNumber={claim.suite}
        fill={fill}
        width={dieWidth}
        height={dieHeight}
      />
    </>
  ) : null
}

export default Claim
