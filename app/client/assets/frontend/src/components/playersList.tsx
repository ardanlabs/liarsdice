import { shortenAddress } from '@usedapp/core'
import { useContext } from 'react'
import { GameContext } from '../gameContext'
import { user } from '../types/index.d'

interface PlayersListProps {
  title: string
}

const PlayersList = (props: PlayersListProps) => {
  const { title } = props
  const { game } = useContext(GameContext)
  const { cups, current_player } = game
  const playersElements: JSX.Element[] = []
  if ((cups as user[]).length) {
    Array.from(cups as user[]).forEach((player) => {
      playersElements.push(
        <li
          style={{ textAlign: 'start' }}
          className={current_player === player.account ? 'active' : ''}
          key={player.account}
        >
          {shortenAddress(player.account)}
        </li>,
      )
    })
  }
  return (
    <div
      className="list_of__players"
      style={{ height: '50%', flexGrow: '1', textAlign: 'start' }}
    >
      <span>
        {title} ({(cups as user[]).length ? (cups as user[]).length : 0})
      </span>
      <ul>{playersElements}</ul>
    </div>
  )
}

export default PlayersList
