import React from 'react'
import { bet, game } from '../types/index.d'

export interface gameContextInterface {
  game: game
  setGame: React.Dispatch<React.SetStateAction<game>>
}

export const GameContext = React.createContext({
  game: {
    status: 'gameover',
    lastOut: '',
    lastWin: '',
    currentPlayer: '',
    currentCup: 0,
    round: 1,
    cups: [],
    playerOrder: [],
    bets: [] as bet[],
    anteUsd: 0,
  } as game,
  setGame: (() => {}) as React.Dispatch<React.SetStateAction<game>>,
})
