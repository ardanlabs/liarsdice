// local holds the web3 connection with the local eth network.
const local = new Web3(Web3.givenProvider || "ws://localhost:8545");

// ABI is the JSON representation of the contract.

const json = [{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"string","name":"value","type":"string"}],"name":"EventLog","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"string","name":"uuid","type":"string"}],"name":"EventNewGame","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"player","type":"address"},{"indexed":false,"internalType":"string","name":"uuid","type":"string"},{"indexed":false,"internalType":"uint256","name":"amount","type":"uint256"}],"name":"EventPlaceAnte","type":"event"},{"inputs":[],"name":"Deposit","outputs":[],"stateMutability":"payable","type":"function"},{"inputs":[],"name":"Game","outputs":[{"internalType":"uint256","name":"created_at","type":"uint256"},{"internalType":"bool","name":"finished","type":"bool"},{"internalType":"uint256","name":"pot","type":"uint256"},{"internalType":"uint256","name":"ante","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"player","type":"address"}],"name":"GameEnd","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"GamePot","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"NewGame","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"Owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"PlaceAnte","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"PlayerBalance","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"Withdraw","outputs":[],"stateMutability":"payable","type":"function"}]; 
const abi = JSON.parse(JSON.stringify(json));

const contract_address = "0x354A70cb44af82EB856F9498e2CA0D4bA248D590"
const owner_address = "0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd"

function getContract() {
  return new local.eth.Contract(
    abi,
    contract_address
  );
}

async function connect() {
  if (window.ethereum) {
    await window.ethereum.request({ method: "eth_requestAccounts" });
    window.web3 = new Web3(window.ethereum);
  } else {
    console.log("No wallet");
  }
}

async function ante() {
  var tx = {
    from: owner_address,
    value: local.utils.toWei("3914402.20", "gwei")
  }

  const result = await getContract().methods.PlaceAnte().send(tx);

  console.log("ante:", result);
}

async function deposit() {
  const from = window.ethereum.selectedAddress;

  // deposit in dollars, convert to gwei.
  var tx = {
    from: from,
    value: local.utils.toWei("3914402.20", "gwei")
  }

  const result = await getContract().methods.Deposit().send(tx);

  console.log("deposit:", result);
}

async function withdraw() {
  const from = window.ethereum.selectedAddress;

  var tx = {
    from: from,
  }

  const result = await getContract().methods.Withdraw().send(tx);

  console.log("withdraw:", result);
}

async function end() {
  // This will be the winner.
  const from = window.ethereum.selectedAddress;

  var tx = {
    from: from,
  }

  const result = await getContract().methods.EndGame().send(tx);

  console.log("end game:", result);
}

async function pot() {
  var tx = {
    from: owner_address,
  }

  const result = await getContract().methods.GamePot().call(tx);

  console.log("game pot:", result);
}