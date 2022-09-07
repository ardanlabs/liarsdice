import React from 'react'
import { bet, game } from '../types/index.d'

export interface gameContextInterface {
  game: game
  setGame: React.Dispatch<React.SetStateAction<game>>
}

export const GameContext = React.createContext({
  game: {
    status: 'nogame',
    lastOut: '',
    lastWin: '',
    currentPlayer: '',
    currentCup: 0,
    round: 1,
    cups: [],
    balances: [],
    playerOrder: [],
    bets: [] as bet[],
    currentID: '',
    anteUSD: 0,
  } as game,
  setGame: (() => {}) as React.Dispatch<React.SetStateAction<game>>,
})
