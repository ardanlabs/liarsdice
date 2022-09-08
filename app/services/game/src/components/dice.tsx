import Die from './icons/die'
import { useContext } from 'react'
import { GameContext } from '../contexts/gameContext'
import { DiceProps } from '../types/props.d'

// Dice component renders 5 die, if the die is 0 or unknow will return a ? die.
function Dice(props: DiceProps) {
  // Extracts props.
  const { diceNumber, isPlayerTurn } = props

  // Initialize an empty JSX array.
  const dice: JSX.Element[] = []

  // Extracts game from useContext hook.
  const { game } = useContext(GameContext)

  // Adds each die to the dice array only if the game is playing.
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
