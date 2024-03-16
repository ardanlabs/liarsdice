// gameConnect does everything to connect the browser to the wallet and
// to the game engine. If successful, a JWT is returned that is needed
// for other game engine API calls.
async function gameConnect(url) {

    // Get configuration information from the game engine.
    const cfg = await config(url);
    if (isError(cfg)) {
        return;
    }

    // Ask the user's wallet is talking to the same blockchain as
    // the game engine.
    const sw = await switchChain(cfg.chainId);
    if (isError(sw)) {

        // The blockchain does not exist in the user's wallet so
        // let's try to help them.
        const aec = await addEthereumChain(cfg);
        if (isError(aec)) {
            return;
        }

        // Try one more time to switch the wallet.
        sw = await switchChain(cfg.chainId);
        if (isError(sw)) {
            return;
        }
    }

    // Request permission to use the wallet. The user will select an
    // account to use.
    const rp = await requestPermissions();
    if (isError(rp)) {
        return;
    }

    // Capture the account that the user selected.
    if (rp.length != 1) {
        isError(newError("user didn't select one account"));
        return;
    }
    if (rp[0].caveats.length != 1) {
        isError(newError("user didn't select one account"));
        return;
    }
    if (rp[0].caveats[0].value.length != 1) {
        isError(newError("user didn't select one account"));
        return;
    }
    const address = rp[0].caveats[0].value[0];

    // Get the current time to sign data to send to the game engine.
    const dateTime = currentDateTime();

    // Sign the arbitrary data.
    const signature = await personalSign(address, cfg.chainId, dateTime);
    if (isError(signature)) {
        return;
    }
    
    // Connect to the game engine to get a token for game play.
    const cge = await connectGameEngine(url, address, cfg.chainId, dateTime, signature);
    if (isError(cge)) {
        return;
    }

    return cge.token;
}

async function config(url) {
    try {
        const result = await $.ajax({
            type: "get",
            url: `${url}/v1/game/config`
        });

        return result;
    }
    
    catch (e) {
        return e;
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
    
        return result;
    }

    catch (e) {
        return e;
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
    
        return result;
    }

    catch (e) {
        return e;
    }
}

async function requestPermissions() {
    try {
        const res = await ethereum.request({
            method: 'wallet_requestPermissions',
            params: [
                {
                    "eth_accounts": {}
                }
            ],
        });

        return res;
    }

    catch (e) {
        return e;
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

        return signature;
    }

    catch (e) {
        return e;
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

        return token;
    }

    catch (e) {
        return e.responseJSON;
    }
}
