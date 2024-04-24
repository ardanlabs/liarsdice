// This is the first screen of the game. Here the player will connect to wallet
// Find other games or start a new game
class MenuScene extends Phaser.Scene {
    constructor() {
        super({key: 'menu'});
    }

    preload() {
        // Load assets for your menu UI (images, fonts)
        this.load.image('background_image', '/assets/greenbackground.png');
        this.load.image('title_image', '/assets/titleimage.png');
        this.load.image('dice_image', '/assets/Dice.png');
        this.load.image('connect_button', '/assets/connectbtn.png');
        this.load.image('joingame_button', '/assets/joinbtn.png');
        this.load.image('newgame_button', '/assets/newbtn.png');
    }

    create() {
        // Create and position menu UI elements

        const backgroundImage = this.add.image(450, 300, 'background_image');
        backgroundImage.setScale(1.5);

        const titleImage = this.add.image(400, 105, 'title_image');
        titleImage.setScale(1);

        const diceImage = this.add.image(90, 130, 'dice_image');
        diceImage.setScale(0.5);

        const connectButton = this.add.image(400, 300, 'connect_button');
        connectButton.setInteractive();

        const joinButton = this.add.image(250, 400, 'joingame_button');
        joinButton.setInteractive();

        const newGameButton = this.add.image(550, 400, 'newgame_button');
        newGameButton.setInteractive();

        connectButton.on('pointerdown', () => {
            // The logic for connectic the wallet should be here
            this.scene.start('game'); // Start the game scene
        });

        joinButton.on('pointerdown', () => {
            // The logic for finding available tables and/or selecting one
            // should be here
        });

        newGameButton.on('pointerdown', () => {
            // The logic for a new game should be here
        });
    }
}

export default MenuScene;
