import Engine from '../backend/engine.js';
import Wallet from '../backend/wallet.js';

class Backend {
    Engine;
    Wallet;

    // -------------------------------------------------------------------------

    constructor(url) {
        this.Engine = new Engine(url);
        this.Wallet = new Wallet();
    }

    // -------------------------------------------------------------------------

    // GameConnect does everything to connect the browser to the wallet and
    // to the game engine.
    async GameConnect() {
        // Get configuration information from the game engine.
        var [cfg, err] = await this.Engine.Config();
        if (err != null) {
            return err;
        }

        // Ask the user's wallet if it's talking to the same blockchain as
        // the game engine.
        var [_, err] = await this.Wallet.SwitchChain(cfg.chainId);
        if (err != null) {
            // The blockchain does not exist in the user's wallet so
            // let's try to help them.
            var [_, err] = await this.Wallet.AddEthereumChain(cfg.chainId, cfg.network);
            if (err != null) {
                return err;
            }

            // Try one more time to switch the wallet.
            var [_, err] = await this.Wallet.SwitchChain(cfg.chainId);
            if (err != null) {
                return err;
            }
        }

        // Request permission to use the wallet. The user will select an
        // account to use.
        var [rp, err] = await this.Wallet.RequestPermissions();
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
        var [sig, err] = await this.Wallet.PersonalSign(address, cfg.chainId, dateTime);
        if (err != null) {
            return err;
        }

        // Connect to the game engine to get a token for game play.
        var err = await this.Engine.Connect(address, cfg.chainId, dateTime, sig);
        if (err != null) {
            return err;
        }

        return null;
    }
}

export default Backend;

// =============================================================================

function currentDateTime() {
    const dt = new Date();

    const year = dt.getUTCFullYear();
    const month = String(dt.getUTCMonth() + 1).padStart(2, '0'); // Month (0-indexed)
    const day = String(dt.getUTCDate()).padStart(2, '0');
    const hours = String(dt.getUTCHours()).padStart(2, '0');
    const minutes = String(dt.getUTCMinutes()).padStart(2, '0');
    const seconds = String(dt.getUTCSeconds()).padStart(2, '0');

    return `${year}${month}${day}${hours}${minutes}${seconds}`;
}
