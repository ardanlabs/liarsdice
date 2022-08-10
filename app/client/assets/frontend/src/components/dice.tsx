import Die from "./icons/die"
import { dice } from '../types/index.d'

interface DiceProps {
  // This type spec is to prevent user from passing an array bigger than 5
  diceNumber: dice
  isPlayerTurn: boolean
}

const Dice = (props: DiceProps) => {
  const { diceNumber, isPlayerTurn } = props
  const dice: JSX.Element[] = []
  if(diceNumber.length) {
    diceNumber.forEach((die) => {
      dice.push(<Die key={die} fill={isPlayerTurn ? 'var(--primary-color)' : 'var(--secondary-color)'} dieNumber={die}></Die>)
    });
  } else {
    for (let i = 0; i < 5; i++) {
      dice.push(<Die key={i} fill={isPlayerTurn ? 'var(--primary-color)' : 'var(--secondary-color)'} />)
    }
  }
  return (
    <div className="dice">
      { dice }
    </div>
  )
}

export default Dice
