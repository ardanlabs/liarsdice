import React, { useContext } from 'react'
import { shortenIfAddress } from '../utils/address'
import { GameContext } from '../contexts/gameContext'

// LiarsCall component
function LiarsCall() {
  // Extracts game from useContext hook.
  const { game } = useContext(GameContext)

  // If there's a lastWin and lastOut it renders the last liar call.
  return game.lastWin && game.lastOut ? (
    <div
      style={{
        display: 'flex',
        height: 'auto',
        width: '60%',
        justifyContent: 'center',
        textAlign: 'center',
        alignItems: 'center',
        color: 'var(--secondary-color)',
        backgroundColor: 'var(--modals)',
        borderRadius: '8px',
        fontSize: '28px',
        fontWeight: '500',
        padding: '8px',
      }}
    >
      <span>
        Player {shortenIfAddress(game.lastWin)} called Player{' '}
        {shortenIfAddress(game.lastOut)} a liar and got striked
      </span>
    </div>
  ) : null
}

export default LiarsCall
