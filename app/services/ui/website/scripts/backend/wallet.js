const sdk = new MetaMaskSDK.MetaMaskSDK({
    dappMetadata: {
        name: 'Liars Dice',
        url: 'http://' + window.location.host,
    },
    logging: {
        sdk: true,
    },
});

class Wallet {
    async SwitchChain(chainId) {
        try {
            const result = await ethereum.request({
                method: 'wallet_switchEthereumChain',
                params: [
                    {
                        chainId: '0x' + chainId.toString(16),
                    },
                ],
            });

            return [result, null];
        } catch (e) {
            return [null, parseError(e)];
        }
    }

    async AddEthereumChain(chainId, network) {
        try {
            const result = await ethereum.request({
                method: 'wallet_addEthereumChain',
                params: [
                    {
                        chainId: '0x' + chainId.toString(16),
                        chainName: 'Liars Dice Local',
                        rpcUrls: [cfg.network],
                        iconUrls: [],
                        nativeCurrency: {
                            name: 'Ethereum',
                            symbol: 'ETH',
                            decimals: 18,
                        },
                        blockExplorerUrls: [network],
                    },
                ],
            });

            return [result, null];
        } catch (e) {
            return [null, parseError(e)];
        }
    }

    async RequestPermissions() {
        try {
            const result = await ethereum.request({
                method: 'wallet_requestPermissions',
                params: [
                    {
                        eth_accounts: {},
                    },
                ],
            });

            return [result, null];
        } catch (e) {
            return [null, parseError(e)];
        }
    }

    async PersonalSign(address, chainId, dateTime) {
        const data = `{"address":"${address}","chainId":${chainId},"dateTime":"${dateTime}"}`;

        try {
            const signature = await ethereum.request({
                method: 'personal_sign',
                params: [hexer(data), address],
            });

            return [signature, null];
        } catch (e) {
            return [null, parseError(e)];
        }
    }
}

export default Wallet;

// =============================================================================

function parseError(e) {
    switch (true) {
        case 'message' in e:
            return e.ReferenceError;
    }

    return 'no error field identified';
}

function hexer(input) {
    const utf8 = toUTF8Array(input);
    const hex = utf8.map((n) => n.toString(16));
    return '0x' + hex.join('');
}

function toUTF8Array(str) {
    var utf8 = [];

    for (var i = 0; i < str.length; i++) {
        var charcode = str.charCodeAt(i);

        if (charcode < 0x80) {
            utf8.push(charcode);
        } else if (charcode < 0x800) {
            utf8.push(0xc0 | (charcode >> 6), 0x80 | (charcode & 0x3f));
        } else if (charcode < 0xd800 || charcode >= 0xe000) {
            utf8.push(0xe0 | (charcode >> 12), 0x80 | ((charcode >> 6) & 0x3f), 0x80 | (charcode & 0x3f));
        } else {
            // Surrogate pair.
            i++;

            // UTF-16 encodes 0x10000-0x10FFFF by
            // subtracting 0x10000 and splitting the
            // 20 bits of 0x0-0xFFFFF into two halves
            charcode = 0x10000 + (((charcode & 0x3ff) << 10) | (str.charCodeAt(i) & 0x3ff));
            utf8.push(0xf0 | (charcode >> 18), 0x80 | ((charcode >> 12) & 0x3f), 0x80 | ((charcode >> 6) & 0x3f), 0x80 | (charcode & 0x3f));
        }
    }

    return utf8;
}
