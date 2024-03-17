class Engine {
    url;
    token;

    // -------------------------------------------------------------------------

    constructor(url) {
        this.url = url;
    }

    // -------------------------------------------------------------------------

    async config() {
        try {
            const result = await $.ajax({
                type: "get",
                url: `${this.url}/v1/game/config`
            });

            return [result, null];
        }
        
        catch (e) {
            if ('responseJSON' in e) {
                return [null, e.responseJSON];
            }
            return [null, e.responseText];
        }
    }

    async connect(address, chainId, dateTime, sigature) {
        const data = `{"address":"${address}","chainId":${chainId},"dateTime":"${dateTime}","sig":"${sigature}"}`;

        try {
            const result = await $.ajax({
                type: "post",
                url: `${this.url}/v1/game/connect`,
                data: data
            });

            this.token = result.token;

            return null;
        }

        catch (e) {
            if ('responseJSON' in e) {
                return [null, e.responseJSON];
            }
            return e.responseText;
        }
    }

    static async queryTables() {
        try {
            const tables = await $.ajax({
                type: "get",
                url: `${this.url}/v1/game/table`,
                beforeSend: function(xhr) {
                    xhr.setRequestHeader ("Authorization", "Bearer " + this.token);
                },
            });

            return [tables, null];
        }

        catch (e) {
            if ('responseJSON' in e) {
                return [null, e.responseJSON];
            }
            return [null, e.responseText];
        }
    }
}