var connect = {
    url: "http://0.0.0.0:3000",
    token: ""
};

window.onload = function () {
    wireEvents();
}

function wireEvents() {
    const gameConnect = document.getElementById("gameConnect");
    gameConnect.addEventListener(
        'click',
        function () { eventGameConnect(connect.url); },
        false
    );

    const gameTables = document.getElementById("gameTables");
    gameTables.addEventListener(
        'click',
        function () { eventGameTables(connect.url, connect.token); },
        false
    );
}

async function eventGameConnect(url) {
    const [token, err] = await gameConnect(url);
    if (err != null) {
        $("#error").text(err);
        return;
    }

    connect.token = token;

    // For now display the token.
    $("#error").text(token);
}

async function eventGameTables(url, token) {
    const [tables, err] = await getGameTables(url, token);
    if (err != null) {
        $("#error").text(err);
        return;
    }

    $("#error").text(JSON.stringify(tables));
}