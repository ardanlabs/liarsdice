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
    const token = await gameConnect(url);
    if (isError(token)) {
        return;
    }

    connect.token = token;

    // For now display the token.
    $("#error").text(token);
}

async function eventGameTables(url, token) {
    const tables = await getGameTables(url, token);
    if (isError(tables)) {
        return;
    }

    $("#error").text(JSON.stringify(tables));
}