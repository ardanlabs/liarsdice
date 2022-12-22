import colors from './colors'
import MainScene from '../scenes/mainScene'

export const DEFAULT_WIDTH = document.documentElement.clientWidth ?? 1280
export const DEFAULT_HEIGHT = document.documentElement.clientHeight ?? 720
export const DIE_PER_PLAYER = 5

export const config: Phaser.Types.Core.GameConfig = {
  type: Phaser.AUTO,
  backgroundColor: colors.borders,
  width: DEFAULT_WIDTH,
  height: DEFAULT_HEIGHT,
  scale: {
    parent: 'gameContainer', // this has to match the div id in index.html
    // mode: Phaser.Scale.ScaleModes.FIT,
    autoCenter: Phaser.Scale.CENTER_BOTH,
  },
  dom: {
    createContainer: true,
  },
  physics: {
    default: 'arcade',
    arcade: {
      gravity: { y: 300 },
      debug: true,
      debugShowBody: true,
    },
  },
  scene: MainScene,
}
