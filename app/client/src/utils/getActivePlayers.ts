import { game } from '../types/index.d'

const getActivePlayersLength = (gameToFilter: game) => {
  return gameToFilter.player_order.filter((player: string) => player.length)
    .length
}

export default getActivePlayersLength
