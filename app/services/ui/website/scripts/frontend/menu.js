// This is the first screen of the game. Here the player will connect to wallet
// Find other games or start a new game
class MenuScene extends Phaser.Scene {
    constructor() {
        super({ key: 'menu' });
    }

    preload() {
        // Load assets for your menu UI (images, fonts)
        this.load.image('background_image', './assets/greenbackground.png');
        this.load.image('dice_image', './assets/Dice.png');
        this.load.image('connect_button', './assets/connectbtn.png');
        this.load.image('joingame_button', './assets/joinbtn.png');
        this.load.image('newgame_button', './assets/newbtn.png');
    }

    create() {
        // Create and position menu UI elements
        const backgroundImage = this.add.image(0, 0, 'background_image');
        backgroundImage.setScale(0.5); // Adjust image scale if needed

        const connectButton = this.add.image(400, 400, 'connect_button');
        connectButton.setInteractive();
        const joinButton = this.add.image(400, 400, 'joingame_button');
        joinButton.setInteractive();
        const newGameButton = this.add.image(400, 400, 'newgame_button');
        newGameButton.setInteractive();


        // Handle button click to start the game
        startButton.on('pointerdown', () => {
            this.scene.start('game'); // Start the game scene
        });
    }
}
  
export default MenuScene;