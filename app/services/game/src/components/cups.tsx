import React, { FC, useContext } from 'react'
import { claim, CupsProps, user } from '../types/index.d'
import Star from './icons/star'
import Claim from './claim'
import { GameContext } from '../gameContext'
import { shortenIfAddress } from '@usedapp/core'

const Cups: FC<CupsProps> = (CupsProps) => {
  const { game } = useContext(GameContext)
  const { cups, player_order, current_cup, status } = game
  const cupsElements: JSX.Element[] = []

  ;(cups as user[]).forEach((player: user, i: number) => {
    const claims = game.claims
      ? game.claims.filter((claim: claim) => claim.account === player.account)
      : []
    const isPlayerTurn =
      (player_order as string[])[current_cup] === player.account &&
      status === 'playing'
    const isPlayerActive = player.outs < 3
    const stars: JSX.Element[] = []
    for (let i = 1; i <= 3; i++) {
      stars.push(
        <Star
          key={Math.random()}
          fill={i > player.outs ? '#F0EAD6' : 'var(--primary-color)'}
        />,
      )
    }
    cupsElements.push(
      <div
        key={player.account}
        className={`player__ui ${isPlayerActive ? 'active' : ''}`}
        data-testid="player__ui"
      >
        <div className="d-flex">{stars}</div>
        <p
          style={{ fontWeight: '600' }}
          className={isPlayerTurn ? 'own_turn' : ''}
        >{`Player ${shortenIfAddress(player.account)}`}</p>
        {isPlayerActive ? (
          <div className={`claim`}>
            {claims[0] ? 'Claim: ' : ''}
            <Claim
              claim={claims[0]}
              dieWidth="27"
              dieHeight="27"
              fill="var(--modals)"
            />
          </div>
        ) : null}
      </div>,
    )
  })
  return (
    <div
      style={{
        display: 'flex',
        width: '100%',
        alignItems: 'center',
        justifyContent: 'center',
        minHeight: '414px',
        flexWrap: 'wrap',
      }}
      id="cupsContainer"
    >
      {cupsElements}
    </div>
  )
}

export default Cups
