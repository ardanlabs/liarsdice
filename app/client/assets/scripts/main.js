async function connect() {
    if (window.ethereum) {
        await window.ethereum.request({ method: "eth_requestAccounts" });
        
        window.web3 = new Web3(window.ethereum);

        // Brutal copy/paste of the LiarsDice.abi content. I don't know how to
        // open it here.
        let json = JSON.parse(JSON.stringify([
            {
              "inputs": [],
              "stateMutability": "nonpayable",
              "type": "constructor"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "string",
                  "name": "value",
                  "type": "string"
                }
              ],
              "name": "EventLog",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "string",
                  "name": "uuid",
                  "type": "string"
                }
              ],
              "name": "EventNewGame",
              "type": "event"
            },
            {
              "anonymous": false,
              "inputs": [
                {
                  "indexed": false,
                  "internalType": "address",
                  "name": "player",
                  "type": "address"
                },
                {
                  "indexed": false,
                  "internalType": "string",
                  "name": "uuid",
                  "type": "string"
                },
                {
                  "indexed": false,
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "name": "EventPlaceAnte",
              "type": "event"
            },
            {
              "inputs": [
                {
                  "internalType": "string",
                  "name": "uuid",
                  "type": "string"
                }
              ],
              "name": "GameAnte",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "player",
                  "type": "address"
                },
                {
                  "internalType": "string",
                  "name": "uuid",
                  "type": "string"
                }
              ],
              "name": "GameEnd",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "string",
                  "name": "uuid",
                  "type": "string"
                }
              ],
              "name": "NewGame",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [],
              "name": "Owner",
              "outputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "string",
                  "name": "uuid",
                  "type": "string"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                },
                {
                  "internalType": "uint256",
                  "name": "minimum",
                  "type": "uint256"
                }
              ],
              "name": "PlaceAnte",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "player",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "name": "deposit",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "string",
                  "name": "",
                  "type": "string"
                }
              ],
              "name": "games",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "created_at",
                  "type": "uint256"
                },
                {
                  "internalType": "bool",
                  "name": "finished",
                  "type": "bool"
                },
                {
                  "internalType": "uint256",
                  "name": "pot",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "",
                  "type": "address"
                }
              ],
              "name": "playerbalance",
              "outputs": [
                {
                  "internalType": "uint256",
                  "name": "",
                  "type": "uint256"
                }
              ],
              "stateMutability": "view",
              "type": "function"
            },
            {
              "inputs": [
                {
                  "internalType": "address",
                  "name": "player",
                  "type": "address"
                },
                {
                  "internalType": "uint256",
                  "name": "amount",
                  "type": "uint256"
                }
              ],
              "name": "withdraw",
              "outputs": [],
              "stateMutability": "nonpayable",
              "type": "function"
            }
          ]));

        var from = window.ethereum.selectedAddress;

        var contract = new window.web3.eth.Contract(json, "0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd");

        await contract.methods.PlaceAnte("liarsgame", 1, 1).send({from: from, gasPrice: "2200000000000", gas: "2200000", value: web3.utils.toWei("1", "ether")}).then(function(receipt){
            console.log(receipt);
        });

    } else {
        console.log("No wallet");
    }
}