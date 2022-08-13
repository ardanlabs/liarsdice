import { useContext } from 'react'
import { GameContext } from '../gameContext'
import { user } from '../types/index.d'

interface PlayersListProps {
  title: string
}

const PlayersList = (props: PlayersListProps) => {
  const { title } = props
  const { game } = useContext(GameContext)
  const { players, current_player } = game
  const playersElements: JSX.Element[] = []
  if ((players as user[]).length) {
    Array.from(players as user[]).forEach((player) => {
      playersElements.push(
        <li
          style={{ textAlign: 'start' }}
          className={current_player === player.wallet ? 'active' : ''}
          key={player.wallet}
        >
          {player.wallet.slice(0, 7)}...
          {player.wallet.slice(player.wallet.length - 7, player.wallet.length)}
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
        {title} ({(players as user[]).length ? (players as user[]).length : 0})
      </span>
      <ul>{playersElements}</ul>
    </div>
  )
}

export default PlayersList
