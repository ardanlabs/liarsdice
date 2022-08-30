import React, { FC } from 'react'
import Claim from './claim'
import { claim } from '../types/index.d'
import { shortenIfAddress } from '@usedapp/core'

interface CurrentClaimProps {
  currentClaim: claim
}

const CurrentClaim: FC<CurrentClaimProps> = (CurrentClaimProps) => {
  const { currentClaim } = CurrentClaimProps

  return (
    <div
      data-testid="current_claim_text_container"
      style={{
        display: 'flex',
        justifyContent: 'center',
        flexDirection: 'column',
        textAlign: 'center',
        alignItems: 'center',
        color: 'var(--modals)',
        borderRadius: '8px',
        fontSize: '28px',
        fontWeight: '500',
        height: '100%',
      }}
    >
      {/* Checks if there's a claim and display who maded it */}
      {currentClaim.account?.length ? (
        <span>
          Current claim by Player {shortenIfAddress(currentClaim.account)}
        </span>
      ) : (
        ''
      )}
      {/* Returns an empty box if there's no claim. Works with Claim Component logic. */}
      <div
        data-testid="current_claim_container"
        style={{
          color: 'var(--secondary-color)',
          fontSize: '28px',
          fontWeight: '500',
          borderRadius: '8px',
          height: '102px',
          width: '322px',
          backgroundColor: 'var(--modals)',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          marginTop: '38px',
          marginBottom: '20px',
        }}
      >
        <Claim claim={currentClaim} fill="var(--secondary-color)" />
      </div>
    </div>
  )
}

export default CurrentClaim
