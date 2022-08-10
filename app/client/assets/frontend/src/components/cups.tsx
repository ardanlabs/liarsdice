import React, { FC } from 'react'
import { user } from '../types/index.d'
import Star from './icons/star'
import Claim from './claim'
import Dice from './dice'

interface CupsProps {
  activePlayers: user[]
  currentPlayerAddress: string
}

const Cups: FC<CupsProps> = (CupsProps) => {
  const { activePlayers, currentPlayerAddress } = CupsProps
  const cupsElements: JSX.Element[] = []
  activePlayers.forEach((player, i) => {
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
          className="player__ui"
        >
          <div className="d-flex">
            <Star fill="var(--primary-color)" />
            <Star fill="var(--primary-color)" />
            <Star />
          </div>
          <h2>{`Player ${i + 1}`}</h2>
          <div className="claim">
            Claim:
            <Claim claim={player.claim} />
          </div>
          <div className="player__cup active">
            <Dice
              isPlayerTurn={currentPlayerAddress === player.address}
              diceNumber={player.dice}
            />
          </div>
        </div>,
      )
    } else {
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
          className="player__ui"
        >
          <div className="d-flex">
            <Star fill="var(--primary-color)" />
            <Star fill="var(--primary-color)" />
            <Star />
          </div>
          <h2>{`Player ${i + 1}`}</h2>
          <div className="claim">
            Claim:
            <Claim claim={player.claim} />
          </div>
          <div className="player__cup"></div>
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
