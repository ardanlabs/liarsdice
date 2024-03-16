async function getGameTables(url, token) {
    try {
        const tables = await $.ajax({
            type: "get",
            url: `${url}/v1/game/table`,
            beforeSend: function (xhr) {
                xhr.setRequestHeader ("Authorization", "Bearer " + token);
            },
        });

        return tables;
    }

    catch (e) {
        return e.responseJSON;
    }
}