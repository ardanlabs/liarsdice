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