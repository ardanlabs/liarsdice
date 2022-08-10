import React, { FC } from 'react'
import { claim } from '../types/index.d'
import Die from './icons/die'

interface ClaimProps {
  claim: claim
}

const Claim: FC<ClaimProps> = (ClaimProps) => {
  const { claim } = ClaimProps
  return (
    <>
      {`${claim.number} X `}
      <Die
        dieNumber={claim.suite}
        fill="var(--secondary-color)"
        width="59px"
        height="60px"
      />
    </>
  )
}

export default Claim
