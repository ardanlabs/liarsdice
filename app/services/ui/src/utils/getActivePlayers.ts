// Filters cups array

import { user } from '../types/index.d'

// Condition: item has length
export default function getActivePlayersLength(users: user[]) {
  return users.filter((player: user) => player.outs < 3).length
}
