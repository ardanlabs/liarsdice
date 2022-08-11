import { createContext } from "react";
import { game } from "./types/index.d";

export const GameContext = createContext(
  {
    game: {
      status: 'open',
      round: 0,
      current_player: '',
      player_order: [],
      players: [],
    } as game,
    setGame: (() => {}) as React.Dispatch<React.SetStateAction<game>>,
  }
)