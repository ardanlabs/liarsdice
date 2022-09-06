import { dice } from './index.d'

export interface GameTableProps {
  timer: number
}

export interface AppHeaderProps {
  show?: boolean
}

export interface SignOutProps {
  disabled: boolean
}

export interface CounterProps {
  timer: number
  show: boolean
}

export interface DiceProps {
  // This type spec is to prevent user from passing an array bigger than 5
  diceNumber: dice
  isPlayerTurn: boolean
}

export interface MainRoomProps {
  timer: number
}
