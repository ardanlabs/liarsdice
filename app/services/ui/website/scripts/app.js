class App {
    engine;

    // -------------------------------------------------------------------------

    constructor(url) {
        this.engine = new Engine(url);
    }

    // -------------------------------------------------------------------------

    init() {
        // Make sure 'this' is the object and not the html element
        // when these methods are executed by the event listener.
        this.handlerGameConnect = this.handlerGameConnect.bind(this);
        this.handlerGameTables = this.handlerGameTables.bind(this);

        $("#gameConnect").click(this.handlerGameConnect);
        $("#gameTables").click(this.handlerGameTables);
    }

    // -------------------------------------------------------------------------

    async handlerGameConnect() {
        const err = await this.gameConnect();
        if (err != null) {
            $("#error").text(err);
            return;
        }

        // For now display the token.
        $("#error").text(this.engine.token);
    }

    async handlerGameTables() {
        const [tables, err] = await this.engine.tables();
        if (err != null) {
            $("#error").text(err);
            return;
        }

        $("#error").text(JSON.stringify(tables));
    }

    // -------------------------------------------------------------------------

    // gameConnect does everything to connect the browser to the wallet and
    // to the game engine. If successful, a JWT is returned that is needed
    // for other game engine API calls.
    async gameConnect() {

        // Get configuration information from the game engine.
        var [cfg, err] = await this.engine.config();
        if (err != null) {
            return err;
        }

        // Ask the user's wallet is talking to the same blockchain as
        // the game engine.
        var [_, err] = await Wallet.switchChain(cfg.chainId);
        if (err != null) {

            // The blockchain does not exist in the user's wallet so
            // let's try to help them.
            var [_, err] = await Wallet.addEthereumChain(cfg);
            if (err != null) {
                return err;
            }

            // Try one more time to switch the wallet.
            var [_, err] = await Wallet.switchChain(cfg.chainId);
            if (err != null) {
                return err;
            }
        }

        // Request permission to use the wallet. The user will select an
        // account to use.
        var [rp, err] = await Wallet.requestPermissions();
        if (err != null) {
            return err;
        }

        // Capture the account that the user selected.
        if (rp.length != 1) {
            return [null, "user didn't select one account"];
        }
        if (rp[0].caveats.length != 1) {
            return [null, "user didn't select one account"];
        }
        if (rp[0].caveats[0].value.length != 1) {
            return [null, "user didn't select one account"];
        }
        const address = rp[0].caveats[0].value[0];

        // Get the current time to sign data to send to the game engine.
        const dateTime = currentDateTime();

        // Sign the arbitrary data.
        var [sig, err] = await Wallet.personalSign(address, cfg.chainId, dateTime);
        if (err != null) {
            return err;
        }

        // Connect to the game engine to get a token for game play.
        var err = await this.engine.connect(address, cfg.chainId, dateTime, sig);
        if (err != null) {
            return err;
        }

        return null;
    }
}

// =============================================================================

function currentDateTime() {
    const dt = new Date();
    
    const year    = dt.getUTCFullYear();
    const month   = String(dt.getUTCMonth() + 1).padStart(2, '0'); // Month (0-indexed)
    const day     = String(dt.getUTCDate()).padStart(2, '0');
    const hours   = String(dt.getUTCHours()).padStart(2, '0');
    const minutes = String(dt.getUTCMinutes()).padStart(2, '0');
    const seconds = String(dt.getUTCSeconds()).padStart(2, '0');

    return `${year}${month}${day}${hours}${minutes}${seconds}`;
}