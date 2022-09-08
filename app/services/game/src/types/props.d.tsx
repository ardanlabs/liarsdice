import { bet, dice } from './index.d'

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

export interface transactionProps {
  buttonText: string
  action: 'Deposit' | 'Withdraw'
  updateBalance: Function
}

export interface BetProps {
  bet: bet
  dieWidth?: string
  dieHeight?: string
  fill: string
}

export interface ButtonProps {
  clickHandler: Function
  classes?: string
  id?: string
  disabled?: boolean
  children: JSX.Element[] | JSX.Element | string
  style?: React.CSSProperties
  tooltip?: string
}

export interface CurrentBetProps {
  currentBet: bet
}

export interface JoinProps {
  disabled: boolean
}

export interface PlayersListProps {
  title: string
}

export interface SideBarProps {
  ante: number
  gamePot: number
  notificationCenterWidth: string
}

export interface SidebarDetailsProps {
  round: number
  ante?: number
  pot?: number
}
