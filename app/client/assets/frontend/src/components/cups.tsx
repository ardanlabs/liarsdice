import React, { FC, useContext } from 'react'
import { user } from '../types/index.d'
import Star from './icons/star'
import Claim from './claim'
import Cup from './cup'
import { GameContext } from '../gameContext'

interface CupsProps {}

const Cups: FC<CupsProps> = (CupsProps) => {
  const { game, setGame } = useContext(GameContext)
  const { players, current_player } = game
  const cupsElements: JSX.Element[] = []
  Array.from(players as user[]).forEach((player, i) => {
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
          key={player.wallet}
          className="player__ui"
        >
          <div className="d-flex">
            <Star fill="var(--primary-color)" />
            <Star fill="var(--primary-color)" />
            <Star />
          </div>
          <h2 className={current_player === player.wallet ? 'active' : ''}>{`Player ${i + 1}`}</h2>
          {/* <div className="claim">
            {player.claim.number ? 'Claim:' : '' }
            <Claim claim={player.claims[0]} dieWidth="27" dieHeight='27' fill='var(--modals)'/>
          </div> */}
          <Cup player={player} currentPlayerWallet={current_player}/>
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
