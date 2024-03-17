async function config(url) {
    try {
        const result = await $.ajax({
            type: "get",
            url: `${url}/v1/game/config`
        });

        return [result, null];
    }
    
    catch (e) {
        return [null, e.responseJSON];
    }
}

async function connectGameEngine(url, address, chainId, dateTime, sigature) {
    const data = `{"address":"${address}","chainId":${chainId},"dateTime":"${dateTime}","sig":"${sigature}"}`;

    try {
        const token = await $.ajax({
            type: "post",
            url: `${url}/v1/game/connect`,
            data: data
        });

        return [token, null];
    }

    catch (e) {
        return [null, e.responseJSON];
    }
}

async function getGameTables(url, token) {
    try {
        const tables = await $.ajax({
            type: "get",
            url: `${url}/v1/game/table`,
            beforeSend: function (xhr) {
                xhr.setRequestHeader ("Authorization", "Bearer " + token);
            },
        });

        return [tables, null];
    }

    catch (e) {
        return [null, e.responseJSON];
    }
}
