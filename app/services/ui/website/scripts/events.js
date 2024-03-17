class Events {
    url;
    token;

    constructor(url) {
        this.url = url;
    }

    init() {
        // Make sure 'this' is the object and not the html element
        // when these methods are executed by the event listener.
        this.eventGameConnect = this.eventGameConnect.bind(this);
        this.eventGameTables = this.eventGameTables.bind(this);

        const gameConnect = document.getElementById("gameConnect");
        gameConnect.addEventListener(
            'click',
            this.eventGameConnect,
            false
        );

        const gameTables = document.getElementById("gameTables");
        gameTables.addEventListener(
            'click',
            this.eventGameTables,
            false
        );
    }

    async eventGameConnect() {
        const [token, err] = await App.gameConnect(this.url);
        if (err != null) {
            $("#error").text(err);
            return;
        }

        this.token = token;

        // For now display the token.
        $("#error").text(token);
    }

    async eventGameTables() {
        const [tables, err] = await Engine.queryGameTables(this.url, this.token);
        if (err != null) {
            $("#error").text(err);
            return;
        }

        $("#error").text(JSON.stringify(tables));
    }
}