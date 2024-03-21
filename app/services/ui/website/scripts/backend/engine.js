class Engine {
    #url;
    #token;

    // -------------------------------------------------------------------------

    constructor(url) {
        this.#url = url;
    }

    // -------------------------------------------------------------------------

    async #isConnected() {
        return (this.#token != null) ? true : false;
    }

    async Config() {
        try {
            const result = await $.ajax({
                type: "get",
                url: `${this.#url}/v1/game/config`
            });

            return [result, null];
        }
        
        catch (e) {
            return [null, parseError(e)];
        }
    }

    async Connect(address, chainId, dateTime, sigature) {
        try {
            this.#token = null;

            const result = await $.ajax({
                type: "post",
                url: `${this.#url}/v1/game/connect`,
                data: `{"address":"${address}","chainId":${chainId},"dateTime":"${dateTime}","sig":"${sigature}"}`
            });

            this.#token = result.token;

            return null;
        }

        catch (e) {
            return [null, parseError(e)];
        }
    }

    async Tables() {
        try {
            if (!this.isConnected) {
                return [null, "not connected to game engine"];
            }

            const tables = await $.ajax({
                type: "get",
                url: `${this.#url}/v1/game/tables`,
                headers: { "Authorization": "Bearer " + this.#token }
            });

            return [tables, null];
        }

        catch (e) {
            return [null, parseError(e)];
        }
    }
}

export default Engine;

// =============================================================================

function parseError(e) {
    switch (true) {
        case ('responseJSON' in e):
            return e.responseJSON.error;
        case ('responseText' in e):
            return e.responseText;
    }

    return "no error field identified";
}
