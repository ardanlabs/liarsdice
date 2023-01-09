require('responses.d')
require('props.d')

export type dice = readonly die[] | []
export type die = 0 | 1 | 2 | 3 | 4 | 5 | 6
export interface user {
  account: string
  outs: number
  lastBet: bet
  dice: dice
}

export interface bet {
  account: string
  number: number
  suite: die
}

export interface game {
  status:
    | 'newgame'
    | 'playing'
    | 'roundover'
    | 'gameover'
    | 'reconciled'
    | 'nogame'
  lastOut: string
  lastWin: string
  currentPlayer: string
  currentCup: number
  round: number
  cups: user[]
  playerOrder: string[]
  currentID: string
  bets: bet[]
  anteUSD: number
  balances: string[]
}

// Props Interfaces
export interface CupProps {
  player: user
  playerDice: dice
}

export interface appConfig {
  chainId: number
  contractId: string
  network: string
}

export type Falsy = false | 0 | '' | null | undefined

export interface AddEthereumChainParameter {
  chainId: string // A 0x-prefixed hexadecimal string
  chainName: string
  nativeCurrency: {
    name: string
    symbol: string // 2-6 characters long
    decimals: 18
  }
  rpcUrls: string[]
  blockExplorerUrls?: string[]
  iconUrls?: string[] // Currently ignored.
}

export interface DiceConfigs {
  [key: number]: Phaser.Types.GameObjects.Sprite.SpriteConfig
}
