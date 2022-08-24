export type dice = readonly die[] | []
export type die = 0 | 1 | 2 | 3 | 4 | 5 | 6
export interface user {
  account: string
  dice: dice
  outs: number
  claim: claim
}

export interface claim {
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
  claims: claim[]
  ante_usd: number
}

export type getExchangeRateResponse = {
  data: {
    amount: string
    base: 'ETH'
    currency: 'USD'
  }
}
