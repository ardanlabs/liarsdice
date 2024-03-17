// gameConnect does everything to connect the browser to the wallet and
// to the game engine. If successful, a JWT is returned that is needed
// for other game engine API calls.
async function gameConnect(url) {

    // Get configuration information from the game engine.
    var [cfg, err] = await config(url);
    if (err != null) {
        return [null, err];
    }

    // Ask the user's wallet is talking to the same blockchain as
    // the game engine.
    var [_, err] = await switchChain(cfg.chainId);
    if (err != null) {

        // The blockchain does not exist in the user's wallet so
        // let's try to help them.
        var [_, err] = await addEthereumChain(cfg);
        if (err != null) {
            return [null, err];
        }

        // Try one more time to switch the wallet.
        var [_, err] = await switchChain(cfg.chainId);
        if (err != null) {
            return [null, err];
        }
    }

    // Request permission to use the wallet. The user will select an
    // account to use.
    var [rp, err] = await requestPermissions();
    if (err != null) {
        return [null, err];
    }

    // Capture the account that the user selected.
    if (rp.length != 1) {
        return [null, "user didn't select one account"];
    }
    if (rp[0].caveats.length != 1) {
        return [null, "user didn't select one account"];
    }
    if (rp[0].caveats[0].value.length != 1) {
        return [null, "user didn't select one account"];
    }
    const address = rp[0].caveats[0].value[0];

    // Get the current time to sign data to send to the game engine.
    const dateTime = currentDateTime();

    // Sign the arbitrary data.
    var [sig, err] = await personalSign(address, cfg.chainId, dateTime);
    if (err != null) {
        return [null, err];
    }
    
    // Connect to the game engine to get a token for game play.
    var [cge, err] = await connectGameEngine(url, address, cfg.chainId, dateTime, sig);
    if (err != null) {
        return [null, err];
    }

    return [cge.token, null];
}

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

async function switchChain(chainId) {
    try {
        const result = await ethereum.request({
            method: 'wallet_switchEthereumChain',
            params: [
                {
                    "chainId": '0x' + chainId.toString(16)
                }
            ],
        });
    
        return [result, null];
    }

    catch (e) {
        return [null, e.message];
    }
}

async function addEthereumChain(cfg) {
    try {
        const result = await ethereum.request({
            method: 'wallet_addEthereumChain',
            params: [
                {
                    chainId: '0x' + cfg.chainId.toString(16),
                    chainName: "Liars Dice Local",
                    rpcUrls: [
                        cfg.network,
                    ],
                    iconUrls: [],
                    nativeCurrency: {
                        "name": "Ethereum",
                        "symbol": "ETH",
                        "decimals": 18
                    },
                    blockExplorerUrls: [
                        cfg.network
                    ]
                }
            ],
        });
    
        return [result, null];
    }

    catch (e) {
        return [null, e.message];
    }
}

async function requestPermissions() {
    try {
        const result = await ethereum.request({
            method: 'wallet_requestPermissions',
            params: [
                {
                    "eth_accounts": {}
                }
            ],
        });

        return [result, null];
    }

    catch (e) {
        return [null, e.message];
    }
}

async function personalSign(address, chainId, dateTime) {
    const data = `{"address":"${address}","chainId":${chainId},"dateTime":"${dateTime}"}`;

    try {
        const signature = await ethereum.request({
            method: 'personal_sign',
            params: [
                hexer(data),
                address
            ],
        });

        return [signature, null];
    }

    catch (e) {
        return [null, e.message];
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
