import Die from './icons/die'
import { useContext } from 'react'
import { GameContext } from '../contexts/gameContext'
import { DiceProps } from '../types/props.d'

const Dice = (props: DiceProps) => {
  const { diceNumber, isPlayerTurn } = props
  const dice: JSX.Element[] = []
  const { game } = useContext(GameContext)

  if (game.status === 'playing') {
    diceNumber.forEach((die, i) => {
      dice.push(
        <Die
          key={i}
          fill={isPlayerTurn ? 'var(--primary-color)' : 'var(--modals)'}
          dieNumber={die}
        ></Die>,
      )
    })
  }
  return <div className="dice">{dice}</div>
}

export default Dice
