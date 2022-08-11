import React, { FC } from 'react'
import { user } from '../types/index.d'
import Star from './icons/star'
import Claim from './claim'
import Cup from './cup'

interface CupsProps {
  activePlayers: Set<user>
  currentPlayerWallet: string
}

const Cups: FC<CupsProps> = (CupsProps) => {
  const { activePlayers, currentPlayerWallet } = CupsProps
  const cupsElements: JSX.Element[] = []
  Array.from(activePlayers).forEach((player, i) => {
    if (player.active) {
      cupsElements.push(
        <div
          style={{
            display: 'flex',
            flexDirection: 'column',
            height: '414px',
            width: '213px',
            justifyContent: 'start',
            alignItems: 'center',
          }}
          key={player.wallet}
          className="player__ui"
        >
          <div className="d-flex">
            <Star fill="var(--primary-color)" />
            <Star fill="var(--primary-color)" />
            <Star />
          </div>
          <h2 className={currentPlayerWallet === player.wallet ? 'active' : ''}>{`Player ${i + 1}`}</h2>
          <div className="claim">
            {player.claim.number ? 'Claim:' : '' }
            <Claim claim={player.claim} dieWidth="27" dieHeight='27' fill='var(--modals)'/>
          </div>
          <Cup player={player} currentPlayerWallet={currentPlayerWallet}/>
        </div>,
      )
    }
  })
  return (
    <div
      style={{
        display: 'flex',
        width: '100%',
        alignItems: 'start',
      }}
      id="cupsContainer"
    >
      {cupsElements}
    </div>
  )
}

export default Cups
