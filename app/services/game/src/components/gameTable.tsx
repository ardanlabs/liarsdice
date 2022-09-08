import React, { useContext } from 'react'
import { GameContext } from '../contexts/gameContext'
import { dice, die } from '../types/index.d'
import { GameTableProps } from '../types/props.d'
import Counter from './counter'
import Cups from './cups'
import CurrentBet from './currentBet'
import Dice from './dice'
import useEthersConnection from './hooks/useEthersConnection'
import useGame from './hooks/useGame'
import LiarsCall from './liarsCall'
import SideBar from './sidebar'

// GameTable Component.
function GameTable(GameTableProps: GameTableProps) {
  // Deconstructs the props.
  const { timer, notificationCenterWidth } = GameTableProps

  // Extracts the game using the useContext Hook.
  const { game } = useContext(GameContext)

  // Extracts the account from the useEthersConnection Hook.
  const { account } = useEthersConnection()

  const isGamePlaying = game.status === 'playing'

  const isPlayerTurn = game.currentID === account

  const currentBet = game.bets[game.bets.length - 1]
    ? game.bets[game.bets.length - 1]
    : { account: '', number: 0, suite: 1 as die }

  // Gets the playerDice from the localStorage.
  const playerDice = JSON.parse(
    window.localStorage.getItem('playerDice') as string,
  ) as dice

  // ===========================================================================

  // Renders the full game table.
  return (
    <div
      style={{
        width: '100%',
        display: 'flex',
        justifyContent: 'start',
        alignItems: 'center',
        flexDirection: 'column',
        maxWidth: '100vw',
        paddingTop: '20px',
        height: 'calc(100vh - 165px)',
      }}
      id="gameTable"
    >
      {/* <div
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
        > */}
      <Counter show={isPlayerTurn && isGamePlaying} timer={timer} />
      <Cups />
      {isGamePlaying ? (
        <>
          <Dice
            isPlayerTurn={isPlayerTurn && isGamePlaying}
            diceNumber={playerDice}
          />
          <CurrentBet currentBet={currentBet} />
          <LiarsCall />
        </>
      ) : (
        ''
      )}
      {/* </div> */}
      {/* </div> */}
    </div>
  )
}

export default GameTable
