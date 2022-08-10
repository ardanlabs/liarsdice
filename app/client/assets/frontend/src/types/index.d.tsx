export type dice = readonly die[]
export type die = 1 | 2 | 3 | 4 | 5 | 6
export interface user {
  address: string,
  active: boolean,
  dice: dice,
  out: number,
  claim: claim,
}

export interface claim {
  number: number,
  suite: die
}