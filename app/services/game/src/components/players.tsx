import React, { useContext } from 'react'
import { user } from '../types/index.d'
import PlayersList from './playersList'

import { GameContext } from '../contexts/gameContext'
import Join from './join'
import useEthersConnection from './hooks/useEthersConnection'

// Players
function Players() {
  // Extracts account from useEthersConnection hook.
  const { account } = useEthersConnection()

  // Extracts game from useContext hook.
  const { game } = useContext(GameContext)

  // Renders this markup
  return (
    <div
      className="players"
      style={{
        display: 'flex',
        alignItems: 'start',
        flexDirection: 'column',
        backgroundColor: 'var(--modals)',
        borderRadius: '0 8px 8px 0',
        position: 'relative',
        width: '100%',
        flexGrow: '1',
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
      <Join />
    </div>
  )
}

export default Players
