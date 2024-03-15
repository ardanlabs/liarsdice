var chainId;

function config() {
    $.ajax({
        type: "get",
        url: "http://0.0.0.0:3000/v1/game/config",
        success: function (res) {
            console.log(res);
            chainId = res.chainId;
            $("#config").text(JSON.stringify(res));
        },
        error: function (jqXHR, exception) {
            console.log(exception);
            $("#config").text(exception);
        },
    });
}

function connect() {
    ethereum
    .request({
        method: 'eth_requestAccounts',
        params: [],
    })
    .then((res) => {
        console.log('request accounts', res);
        $("#connect").text(JSON.stringify(res));
    })
    .catch((e) => {
        console.log('request accounts ERR', e)
        $("#connect").text(e);
    });
}

function accounts() {
    ethereum
    .request({
        method: 'eth_accounts',
        params: [],
    })
    .then((res) => {
        console.log('accounts', res);
        $("#sign").text(JSON.stringify(res));
    })
    .catch((e) => {
        console.log('accounts ERR', e)
        $("#sign").text(e);
    });
}
            
var signed;
var dateTime;
            
function personalSign() {
    dateTime = currentDateTime();
    const data = `{"address":"0x6327a38415c53ffb36c11db55ea74cc9cb4976fd","chainId":${chainId},"dateTime":"${dateTime}"}`;

    ethereum
    .request({
        method: 'personal_sign',
        params: [
            hexer(data),
            "0x6327a38415c53ffb36c11db55ea74cc9cb4976fd"
        ],
    })
    .then((res) => {
        signed=res;
        console.log('personal sign', res);
        $("#psign").text(JSON.stringify(res) + " " + data);
    })
    .catch((e) => {
        console.log('personal sign ERR', e)
        $("#psign").text(e + " " + data);
    });
}

function connectGE() {
    const data = `{"address":"0x6327a38415c53ffb36c11db55ea74cc9cb4976fd","chainId":${chainId},"dateTime":"${dateTime}","sig":"${signed}"}`;

    $.ajax({
        type: "post",
        url: "http://0.0.0.0:3000/v1/game/connect",
        data: data,
        success: function (res) {
            console.log(res);
            $("#cge").text(JSON.stringify(res));
        },
        error: function (jqXHR, exception) {
            console.log(exception);
            $("#cge").text(exception + " " + data);
        },
    });
}