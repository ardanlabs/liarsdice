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
  config: {
    ChainID: number
    ContractID: string
    Network: string
  }
}
