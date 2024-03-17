class Wallet {
    static async switchChain(chainId) {
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

    static async addEthereumChain(cfg) {
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

    static async requestPermissions() {
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

    static async personalSign(address, chainId, dateTime) {
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
}