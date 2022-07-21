async function connect() {
  if (window.ethereum) {
    await window.ethereum.request({ method: "eth_requestAccounts" });
    window.web3 = new Web3(window.ethereum);
  } else {
    console.log("No wallet");
  }
}

async  function placeAnte() {
  // Web3.givenProvider is MetaMask (in my case)
  const web3 = new Web3(Web3.givenProvider || "ws://localhost:8545");

  const id = await web3.eth.net.getId();
  console.log("id:", id)

  let json = JSON.parse(JSON.stringify([{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"string","name":"value","type":"string"}],"name":"EventLog","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"string","name":"uuid","type":"string"}],"name":"EventNewGame","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"player","type":"address"},{"indexed":false,"internalType":"string","name":"uuid","type":"string"},{"indexed":false,"internalType":"uint256","name":"amount","type":"uint256"}],"name":"EventPlaceAnte","type":"event"},{"inputs":[{"internalType":"string","name":"uuid","type":"string"}],"name":"GameAnte","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"player","type":"address"},{"internalType":"string","name":"uuid","type":"string"}],"name":"GameEnd","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"uuid","type":"string"}],"name":"NewGame","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"Owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"uuid","type":"string"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"PlaceAnte","outputs":[],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"address","name":"player","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"deposit","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"","type":"string"}],"name":"games","outputs":[{"internalType":"uint256","name":"created_at","type":"uint256"},{"internalType":"bool","name":"finished","type":"bool"},{"internalType":"uint256","name":"pot","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"playerbalance","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"player","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"withdraw","outputs":[],"stateMutability":"nonpayable","type":"function"}]));

  const from = window.ethereum.selectedAddress;
  console.log("window.ethereum.selectedAddress:", from)

  const contract = new web3.eth.Contract(
    json,
    "0xaBeb7a32d3162642a9210D88459c00e54efb0776"
  );
  console.log("contract:", contract)

  // web3.eth.subscribe('logs', {} ,function(error, result){
  //     console.log("error", error);
  //     console.log("result", result);
  // });

  let wei;
  web3.eth.getGasPrice().then((result) => {
    console.log(web3.utils.fromWei(result, 'ether'))
    wei = web3.utils.fromWei(result, 'ether')
  });

  //contract.methods.PlaceAnte("liarsgame", 1, 1).estimateGas({gas: wei}, function(error, gasAmount){
  //  console.log(gasAmount);
  //});

  const result = await contract.methods.PlaceAnte("liarsgame", 1).send({from: from, gasPrice: wei, gas: 22000, value: web3.utils.toWei("1", "ether")});
  console.log("result:", result);

  // const result2 = await contract.methods.PlaceAnte("liarsgame", 1).call({from: from, gasPrice: wei, gas: 22000, value: web3.utils.toWei("1", "ether")});
  // console.log("result:", result2);


  // Send creates a transaction.
  // await contract.methods.PlaceAnte("liarsgame", 1).send({from: from, gasPrice: wei, gas: 22000, value: web3.utils.toWei("1", "ether")}).then(function(error, receipt){
  //   console.log(error);
  //   console.log(receipt);
  // });

  console.log("placeAnte done");
}

// "[ethjs-query] while formatting outputs from RPC '{"value":{"code":-32603,"data":{"code":-32000,"message":"intrinsic gas too low"}}}'"