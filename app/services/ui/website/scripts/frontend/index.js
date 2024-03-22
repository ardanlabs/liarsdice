import Backend from './backend.js';
import MenuScene from './menu.js';
import GameScene from './game.js';
// import EndScene from './end.js';

// Construct the backend to have access to the backend API.
const backend = new Backend('http://0.0.0.0:3000');

// =============================================================================

const config = {
    type: Phaser.AUTO,
    width: 800,
    height: 600,
    scene: [MenuScene, GameScene],
};

const game = new Phaser.Game(config);

// =============================================================================

// Throw Away Code

window.onload = function () {
    $('#gameConnect').click(handlerGameConnect);
    $('#gameTables').click(handlerGameTables);
};

async function handlerGameConnect() {
    const err = await backend.GameConnect();
    if (err != null) {
        $('#error').text(err);
        return;
    }

    // For now display the token.
    $('#error').text(backend.Engine.Token());
}

async function handlerGameTables() {
    const [tables, err] = await backend.Engine.Tables();
    if (err != null) {
        $('#error').text(err);
        return;
    }

    $('#error').text(JSON.stringify(tables));
}
