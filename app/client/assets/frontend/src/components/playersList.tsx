import { user } from '../types/index.d'

interface PlayersListProps {
  players: Set<user> | string[]
  title: string
  currentPlayer: string
}

const PlayersList = (props: PlayersListProps) => {
  const { players, title, currentPlayer } = props
  const playersIsSet = typeof players === 'object'
  const playersLength = playersIsSet ? (players as Set<user>).size : (players as string[]).length
  console.log(players, title, 'playersList')
  
  const playersElements: JSX.Element[] = []
  if (playersIsSet && playersLength) {
    Array.from(players as Set<user>).forEach((player) => {
      playersElements.push(
        <li style={{ textAlign: 'start' }} className={currentPlayer === player.wallet ? 'active' : ''} key={player.wallet}>
          {player.wallet.slice(0, 7)}...
          {player.wallet.slice(player.wallet.length - 7, player.wallet.length)}
        </li>,
      )
    })
  } else if (playersLength){
    (players as string[]).forEach((player) => {
      playersElements.push(
        <li style={{ textAlign: 'start' }} key={player}>
          {player.slice(0, 7)}...
          {player.slice(player.length - 7, player.length)}
        </li>,
      )
    });
  }
  return (
    <div
      className="list_of__players"
      style={{ height: '50%', flexGrow: '1', textAlign: 'start' }}
    >
      <span>
        {title} ({playersLength ? playersLength : 0})
      </span>
      <ul>{playersElements}</ul>
    </div>
  )
}

export default PlayersList
