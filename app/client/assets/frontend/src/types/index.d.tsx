export type dice = readonly die[] | []
export type die = 0 | 1 | 2 | 3 | 4 | 5 | 6
export interface user {
  wallet: string,
  active: boolean,
  dice: dice,
  outs: number,
  claim: claim,
}

export interface claim {
  number: number,
  suite: die
}