import { shortenIfAddress } from '@usedapp/core'
import React, { FC, useContext } from 'react'
import { GameContext } from '../gameContext'

interface LiarsCallProps {}

const LiarsCall: FC<LiarsCallProps> = (LiarsCallProps) => {
  const { game } = useContext(GameContext)
  return game.last_win && game.last_out ? (
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
        Player {shortenIfAddress(game.last_win)} called Player{' '}
        {shortenIfAddress(game.last_out)} a liar and got striked
      </span>
    </div>
  ) : null
}

export default LiarsCall
