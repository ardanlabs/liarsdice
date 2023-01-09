export interface AppHeaderProps {
  show?: boolean
}

export interface SignOutProps {
  disabled: boolean
}

export interface TransactionProps {
  buttonText: string
  action: 'Deposit' | 'Withdraw'
  updateBalance: Function
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

export interface JoinProps {
  disabled?: boolean
}

export interface PhaserTestProps {}
