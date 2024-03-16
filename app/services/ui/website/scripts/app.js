function isError(res) {
    if ((res != null) && (typeof res === 'object')) {
        if ('message' in res) {
            $("#error").text(res.message);
            return true;
        }
    
        if ('error' in res) {
            $("#error").text(res.error);
            return true;
        }
    }

    return false;
}

function newError(msg) {
    return { error: msg };
}