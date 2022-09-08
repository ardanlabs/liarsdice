import React from 'react'
import Bet from './bet'
import { shortenIfAddress } from '../utils/address'
import { CurrentBetProps } from '../types/props.d'

// CurrentBet component
function CurrentBet(props: CurrentBetProps) {
  // Extracts props.
  const { currentBet } = props

  // Renders this markup.
  return (
    <div
      data-testid="current_bet_text_container"
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
      {/* Checks if there's a bet and display who maded it */}
      {currentBet.account?.length ? (
        <span>
          Current bet by Player {shortenIfAddress(currentBet.account)}
        </span>
      ) : (
        ''
      )}
      {/* Returns an empty box if there's no bet. Works with Bet Component logic. */}
      <div
        data-testid="current_bet_container"
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
        <Bet bet={currentBet} fill="var(--secondary-color)" />
      </div>
    </div>
  )
}

export default CurrentBet
