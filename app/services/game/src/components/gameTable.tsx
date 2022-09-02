import React, { FC, useContext } from 'react'
import { GameContext } from '../contexts/gameContext'
import { dice } from '../types/index.d'
import Counter from './counter'
import Cups from './cups'
import CurrentBet from './currentBet'
import Dice from './dice'
import useEthersConnection from './hooks/useEthersConnection'
import useGame from './hooks/useGame'
import LiarsCall from './liarsCall'
import SideBar from './sidebar'

interface GameTableProps {
  timer: number
  playerDice: dice
}

const GameTable: FC<GameTableProps> = (GameTableProps) => {
  const { timer, playerDice } = GameTableProps
  const { game } = useContext(GameContext)
  const { account } = useEthersConnection()
  const { gamePot } = useGame()
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
      <SideBar
        ante={game.ante_usd}
        gamePot={gamePot}
        notificationCenterWidth={notificationCenterWidth}
      />
    </div>
  )
}

export default GameTable
