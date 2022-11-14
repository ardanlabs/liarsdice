import { shortenIfAddress } from '../utils/address'
import { useContext } from 'react'
import { GameContext } from '../contexts/gameContext'
import { user } from '../types/index.d'
import { PlayersListProps } from '../types/props.d'
import useEthersConnection from './hooks/useEthersConnection'

// PlayersList component
function PlayersList(props: PlayersListProps) {
  // Extracts props.
  const { title } = props

  // Extracts account from ethersConnection hook.
  const { account } = useEthersConnection()

  // Extracts game from useContext hook.
  const { game } = useContext(GameContext)

  // Extracts game properties.
  const { cups, currentID, status } = game

  // Creates an empty array of players ui elements.
  const playersElements: JSX.Element[] = []

  // If there's any cups it iterates them.
  if ((cups as user[]).length) {
    Array.from(cups as user[]).forEach((player) => {
      const className =
        currentID === player.account && status === 'playing' ? 'active' : ''

      playersElements.push(
        <li
          style={{ textAlign: 'start' }}
          className={className}
          key={player.account}
        >
          {`${
            account === player.account ? 'Me' : shortenIfAddress(player.account)
          }`}
        </li>,
      )
    })
  }

  // Renders this markup
  return (
    <div
      className="list_of__players"
      style={{
        height: '50%',
        flexGrow: '1',
        textAlign: 'start',
        flexDirection: 'column',
        display: 'flex',
        fontWeight: '500',
      }}
    >
      <span>
        {title} ({(cups as user[]).length ? (cups as user[]).length : 0})
      </span>
      <ul
        style={{
          listStyle: 'none',
          padding: '0',
        }}
      >
        {playersElements}
      </ul>
    </div>
  )
}

export default PlayersList
