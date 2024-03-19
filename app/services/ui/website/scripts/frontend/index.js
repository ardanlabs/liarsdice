import MenuScene from './menu.js';
import GameScene from './game.js';
// import EndScene from './end.js';

const config = {
    type: Phaser.AUTO,
    width: 800,
    height: 600,
    scene: [MenuScene, GameScene],
};

const game = new Phaser.Game(config);