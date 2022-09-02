export type dice = readonly die[] | []
export type die = 0 | 1 | 2 | 3 | 4 | 5 | 6
export interface user {
  account: string
  outs: number
  bet: bet
}

export interface bet {
  account: string
  number: number
  suite: die
}

export interface game {
  status: string
  last_out: string
  last_win: string
  current_player: string
  current_cup: number
  round: number
  cups: user[]
  player_order: string[]
  bets: bet[]
  ante_usd: number
}

// Props Interfaces
export interface CupProps {
  player: user
  playerDice: dice
}

export interface CupsProps {
  playerDice: dice
}

export interface appConfig {
  chain_id: number
  contract_id: string
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
