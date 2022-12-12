import { DEFAULT_WIDTH, DEFAULT_HEIGHT } from '../src/utils/config'

export const resize = (game: Phaser.Game) => {
  game.scale.resize(DEFAULT_WIDTH, DEFAULT_HEIGHT)
}
