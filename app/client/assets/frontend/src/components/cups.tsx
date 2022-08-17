import React, { FC, useContext } from 'react'
import { user } from '../types/index.d'
import Star from './icons/star'
import Claim from './claim'
import Cup from './cup'
import { GameContext } from '../gameContext'

interface CupsProps {}

const Cups: FC<CupsProps> = (CupsProps) => {
  const { game } = useContext(GameContext)
  const { cups, current_player } = game
  const cupsElements: JSX.Element[] = []

  Array.from(cups as user[]).forEach((player, i) => {
    const stars: JSX.Element[] = []
    for (let i = 1; i <= 3; i++) {
      stars.push(
        <Star fill={i < player.outs ? 'var(--primary-color)' : '#F0EAD6'} />,
      )
    }
    if (player.outs < 3) {
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
          key={player.account}
          className="player__ui"
        >
          <div className="d-flex">{stars}</div>
          <h2
            className={current_player === player.account ? 'active' : ''}
          >{`Player ${i + 1}`}</h2>
          <div className="claim">
            {/* {player.claim.number ? 'Claim: ' : ''}
            <Claim
              claim={player.claim}
              dieWidth="27"
              dieHeight="27"
              fill="var(--modals)"
            /> */}
          </div>
          <Cup player={player} />
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
        minHeight: '414px',
      }}
      id="cupsContainer"
    >
      {cupsElements}
    </div>
  )
}

export default Cups
