import React, { useContext } from 'react'
import Button from './button'
import { user } from '../types/index.d'
import PlayersList from './playersList'
import { useEthers } from '@usedapp/core'
import { GameContext } from '../gameContext'

interface PlayersProps {
  joinGame: Function
}
const Players = (props: PlayersProps) => {
  const { joinGame } = props
  const { account } = useEthers()
  const { game } = useContext(GameContext)
  const isUserPlaying = (game.players as user[]).filter((user) => {
    return user.wallet === account
  })

  return (
    <div
      className="players"
      style={{
        display: 'flex',
        alignItems: 'start',
        flexDirection: 'column',
        backgroundColor: 'var(--modals)',
        borderRadius: '0 8px 8px 0',
        margin: '42px 0 0 0',
        minHeight: '550px',
        aspectRatio: '2/4',
        position: 'relative',
      }}
    >
      <div
        className="players__list"
        style={{
          display: 'flex',
          alignItems: 'start',
          flexDirection: 'column',
          padding: '16px 10px',
          height: '100%',
          flexGrow: '1',
          width: '100%',
        }}
      >
        <PlayersList title="Active players" />
        {/* <PlayersList players={waitingPlayers} title="Waiting players" /> */}
      </div>
      <Button
        disabled={Boolean(isUserPlaying.length)}
        classes="join__buton"
        clickHandler={() => joinGame()}
      >
        <span>Join Game</span>
      </Button>
    </div>
  )
}

export default Players
