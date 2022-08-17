import { useEthers } from '@usedapp/core'
import React, { FC, useContext } from 'react'
import { GameContext } from '../gameContext'
import Counter from './counter'
import Cups from './cups'
import CurrentClaim from './currentClaim'
import LiarsCall from './liarsCall'

interface GameTableProps {
  timer: number
}

const GameTable: FC<GameTableProps> = (GameTableProps) => {
  const { timer } = GameTableProps
  const { game } = useContext(GameContext)
  const { account } = useEthers()

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
      <Counter show={game.current_player === account} timer={timer} />
      <Cups />
      <LiarsCall />
      <CurrentClaim
        currentClaim={{ account: '', claim: { number: 1, suite: 2 } }}
      />
    </div>
  )
}

export default GameTable
