import Die from './icons/die'
import { dice } from '../types/index.d'
import { useContext } from 'react'
import { GameContext } from '../contexts/gameContext'

interface DiceProps {
  // This type spec is to prevent user from passing an array bigger than 5
  diceNumber: dice
  isPlayerTurn: boolean
}

const Dice = (props: DiceProps) => {
  const { diceNumber, isPlayerTurn } = props
  const dice: JSX.Element[] = []
  const { game } = useContext(GameContext)

  if (diceNumber.length && game.status === 'playing') {
    diceNumber.forEach((die, i) => {
      dice.push(
        <Die
          key={i}
          fill={isPlayerTurn ? 'var(--primary-color)' : 'var(--modals)'}
          dieNumber={die}
        ></Die>,
      )
    })
  } else {
    for (let i = 0; i < 5; i++) {
      dice.push(
        <Die
          key={i}
          fill={isPlayerTurn ? 'var(--primary-color)' : 'var(--modals)'}
        />,
      )
    }
  }
  return <div className="dice">{dice}</div>
}

export default Dice
