export type dice = readonly die[] | []
export type die = 0 | 1 | 2 | 3 | 4 | 5 | 6
export interface user {
  wallet: string,
  dice: dice,
  outs: number,
}

export interface claim {
  number: number,
  suite: die
}