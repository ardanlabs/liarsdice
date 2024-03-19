class GameScene extends Phaser.Scene {
    constructor() {
        super({ key: 'game' });
    }
  
    preload() {
        // Load UI asset images, fonts, etc. (e.g., button image)
        this.load.image('background_image', './assets/greenbackground.png');
        this.load.image('one', './assets/one.png');
        this.load.image('two', './assets/two.png');
        this.load.image('three', './assets/three.png');
        this.load.image('four', './assets/four.png');
        this.load.image('five', './assets/five.png');
        this.load.image('six', './assets/six.png');
    }
  
    create() {
        // Create UI elements using Phaser objects
        const backgroundImage = this.add.image(0, 0, 'background_image');
        const connectButton = this.add.image(x, y, 'button_image');
        connectButton.setInteractive();
        connectButton.on('pointerdown', () => {
            // Call your connect function to interact with MetaMask
            connect();
            // Update UI based on response (e.g., display retrieved accounts)
        });

        // Add other UI elements here
    }
}
  
const config = {
    type: Phaser.AUTO,
    width: 800,
    height: 600,
    scene: [GameScene],
};

const game = new Phaser.Game(config);