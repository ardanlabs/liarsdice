import Die from './icons/die'
import { dice } from '../types/index.d'
import { useEthers } from '@usedapp/core'
import { useContext } from 'react'
import { GameContext } from '../gameContext'

interface DiceProps {
  // This type spec is to prevent user from passing an array bigger than 5
  diceNumber: dice
  isPlayerTurn: boolean
  playerAccount: string
}

const Dice = (props: DiceProps) => {
  const { diceNumber, isPlayerTurn, playerAccount } = props
  const { account } = useEthers()
  const dice: JSX.Element[] = []
  const { game } = useContext(GameContext)

  if (
    diceNumber.length &&
    account === playerAccount &&
    game.status === 'playing'
  ) {
    diceNumber.forEach((die, i) => {
      dice.push(
        <Die
          key={i}
          fill={
            isPlayerTurn ? 'var(--primary-color)' : 'var(--secondary-color)'
          }
          dieNumber={die}
        ></Die>,
      )
    })
  } else {
    for (let i = 0; i < 5; i++) {
      dice.push(
        <Die
          key={i}
          fill={
            isPlayerTurn ? 'var(--primary-color)' : 'var(--secondary-color)'
          }
        />,
      )
    }
  }
  return <div className="dice">{dice}</div>
}

export default Dice