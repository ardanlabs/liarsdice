class Engine {
    url;
    token;

    // -------------------------------------------------------------------------

    constructor(url) {
        this.url = url;
    }

    // -------------------------------------------------------------------------

    async isConnected() {
        return (token != null) ? true : false;
    }

    async config() {
        try {
            const result = await $.ajax({
                type: "get",
                url: `${this.url}/v1/game/config`
            });

            return [result, null];
        }
        
        catch (e) {
            return [null, parseError(e)];
        }
    }

    async connect(address, chainId, dateTime, sigature) {
        try {
            this.token = null;

            const result = await $.ajax({
                type: "post",
                url: `${this.url}/v1/game/connect`,
                data: `{"address":"${address}","chainId":${chainId},"dateTime":"${dateTime}","sig":"${sigature}"}`
            });

            this.token = result.token;

            return null;
        }

        catch (e) {
            return [null, parseError(e)];
        }
    }

    async tables() {
        try {
            if (!this.isConnected) {
                return [null, "not connected to game engine"];
            }

            const tables = await $.ajax({
                type: "get",
                url: `${this.url}/v1/game/tables`,
                headers: { "Authorization": "Bearer " + this.token }
            });

            return [tables, null];
        }

        catch (e) {
            return [null, parseError(e)];
        }
    }
}
