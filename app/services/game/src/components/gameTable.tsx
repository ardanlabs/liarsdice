import { useEthers } from '@usedapp/core'
import React, { FC, useContext } from 'react'
import { GameContext } from '../gameContext'
import { dice } from '../types/index.d'
import Counter from './counter'
import Cups from './cups'
import CurrentBet from './currentBet'
import Dice from './dice'
import LiarsCall from './liarsCall'
import NotificationCenter from './notificationCenter/notificationCenter'

interface GameTableProps {
  timer: number
  playerDice: dice
}

const GameTable: FC<GameTableProps> = (GameTableProps) => {
  const { timer, playerDice } = GameTableProps
  const { game } = useContext(GameContext)
  const { account } = useEthers()
  const notificationCenterWidth = '340px'

  return (
    <div
      style={{
        display: 'flex',
        flexGrow: '1',
      }}
    >
      <div
        style={{
          display: 'flex',
          justifyContent: 'start',
          alignItems: 'center',
          flexDirection: 'column',
          width: `calc(100% - ${notificationCenterWidth})`,
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
            <Dice
              isPlayerTurn={
                (game.player_order as string[])[game.current_cup] === account &&
                game.status === 'playing'
              }
              diceNumber={playerDice}
            />
            <CurrentBet
              currentBet={
                game.bets[game.bets.length - 1]
                  ? game.bets[game.bets.length - 1]
                  : { account: '', number: 0, suite: 1 }
              }
            />
            <LiarsCall />
          </>
        ) : (
          ''
        )}
      </div>
      <NotificationCenter
        trigger={false}
        asideContainerStyle={{
          height: '100%',
          display: 'flex',
          flexDirection: 'column',
        }}
        mainContainerStyle={{
          border: '1px inset var(--secondary-color)',
          height: 'calc(100% - 165px)',
          right: '0px',
          position: 'fixed',
          top: '95px',
          width: `${notificationCenterWidth}`,
        }}
      />
    </div>
  )
}

export default GameTable
