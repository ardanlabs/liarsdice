import React, { FC, useContext } from 'react'
import { claim, CupsProps, user } from '../types/index.d'
import Star from './icons/star'
import Claim from './claim'
import Cup from './cup'
import { GameContext } from '../gameContext'
import { shortenIfAddress } from '@usedapp/core'

const Cups: FC<CupsProps> = (CupsProps) => {
  const playerDice =
    JSON.parse(window.localStorage.getItem('playerDice') as string) ?? []
  const { game } = useContext(GameContext)
  const { cups, player_order, current_cup, status } = game
  const cupsElements: JSX.Element[] = []

  ;(cups as user[]).forEach((player: user, i: number) => {
    const claims = game.claims
      ? game.claims.filter((claim: claim) => claim.account === player.account)
      : []
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
        data-testid="player__ui"
      >
        <div className="d-flex">{stars}</div>
        <h2
          className={
            (player_order as string[])[current_cup] === player.account &&
            status === 'playing'
              ? 'active'
              : ''
          }
        >{`Player ${shortenIfAddress(player.account)}`}</h2>
        <div className="claim">
          {claims[0] ? 'Claim: ' : ''}
          <Claim
            claim={claims[0]}
            dieWidth="27"
            dieHeight="27"
            fill="var(--modals)"
          />
        </div>
        <Cup player={player} playerDice={playerDice} />
      </div>,
    )
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
