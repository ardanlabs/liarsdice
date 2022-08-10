import { user } from '../types/index.d'

interface PlayersListProps {
  players: user[]
  title: string
}

const PlayersList = (props: PlayersListProps) => {
  const { players, title } = props

  const playersElements: JSX.Element[] = []
  players.forEach((player) => {
    return playersElements.push(
      <li style={{ textAlign: 'start' }}>
        {player.address.slice(0, 7)}...
        {player.address.slice(player.address.length - 7, player.address.length)}
      </li>,
    )
  })
  return (
    <div
      className="list_of__players"
      style={{ height: '50%', flexGrow: '1', textAlign: 'start' }}
    >
      <span>
        {title} ({players.length})
      </span>
      <ul>{playersElements}</ul>
    </div>
  )
}

export default PlayersList
