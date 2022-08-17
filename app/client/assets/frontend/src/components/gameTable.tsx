import { useEthers } from '@usedapp/core'
import React, { FC, useContext } from 'react'
import { GameContext } from '../gameContext'
import { dice } from '../types/index.d'
import Counter from './counter'
import Cups from './cups'
import CurrentClaim from './currentClaim'
import LiarsCall from './liarsCall'

interface GameTableProps {
  timer: number
  playerDice: dice
}

const GameTable: FC<GameTableProps> = (GameTableProps) => {
  const { timer, playerDice } = GameTableProps
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
      <Counter
        show={
          (game.player_order as string[])[game.current_cup] === account &&
          game.status === 'playing'
        }
        timer={timer}
      />
      <Cups playerDice={playerDice} />
      {game.status === 'playing' ? (
        <>
          <LiarsCall />
          <CurrentClaim
            currentClaim={
              game.claims[game.claims.length - 1]
                ? game.claims[game.claims.length - 1]
                : { account: '', number: 0, suite: 1 }
            }
          />
        </>
      ) : (
        ''
      )}
    </div>
  )
}

export default GameTable
