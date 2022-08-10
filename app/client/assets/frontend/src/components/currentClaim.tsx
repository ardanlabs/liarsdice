import React, { FC } from 'react'
import Claim from './claim'
import { claim } from '../types/index.d'

interface CurrentClaimProps {
  currentClaim: { address: string; claim: claim }
}

const CurrentClaim: FC<CurrentClaimProps> = (CurrentClaimProps) => {
  const { currentClaim } = CurrentClaimProps
  return (
    <div
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
        marginTop: '38px',
        height: '100%',
      }}
    >
      <span>
        Current claim by Player {currentClaim?.address.slice(0, 7)}...
        {currentClaim?.address.slice(
          currentClaim?.address.length - 7,
          currentClaim?.address.length,
        )}
      </span>
      <div
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
        <Claim claim={currentClaim?.claim} />
      </div>
    </div>
  )
}

export default CurrentClaim
