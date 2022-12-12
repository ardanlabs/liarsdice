import 'phaser'
import { config } from '../src/utils/config'
import { resize } from './resize'

window.addEventListener('load', () => {
  let game = new Phaser.Game(config)

  window.addEventListener('resize', () => {
    resize(game)
  })
  resize(game)
})
