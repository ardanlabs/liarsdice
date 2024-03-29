#!/bin/sh

GETH_ADDRESS=http://localhost:8545

if [ "${1}" = "force" ]; then
  kubectl -n liars-system delete secret contract-id
fi

kubectl -n liars-system get secret contract-id >/dev/null 2>&1
RET=${?}

if [ ${RET} -eq 0 ]; then
  echo "Contract already exists, not re-deploying"
  exit 0
fi

CONTRACT_OUTPUT=$(./admin contract deploy -n ${GETH_ADDRESS})
RET=${?}

if [ ${RET} -ne 0 ]; then
  echo "Error deploying contract:"
  echo ${CONTRACT_OUTPUT}
  exit 1
fi

export GAME_CONTRACT_ID=$(echo "${CONTRACT_OUTPUT}" | awk '/export GAME_CONTRACT_ID/ {split($2,a,"="); print a[2]}')
echo $GAME_CONTRACT_ID

kubectl -n liars-system create secret generic contract-id --from-literal=id=${GAME_CONTRACT_ID}

echo "Add balances to the test accounts"

curl -H 'Content-Type: application/json' --data '{"jsonrpc":"2.0","method":"eth_sendTransaction", "params": [{"from":"0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd", "to":"0x8E113078ADF6888B7ba84967F299F29AeCe24c55", "value":"0x1000000000000000000"}], "id":1}' ${GETH_ADDRESS}
curl -H 'Content-Type: application/json' --data '{"jsonrpc":"2.0","method":"eth_sendTransaction", "params": [{"from":"0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd", "to":"0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7", "value":"0x1000000000000000000"}], "id":1}' ${GETH_ADDRESS}

./admin contract -n ${GETH_ADDRESS} -a 0x8e113078adf6888b7ba84967f299f29aece24c55 -m 1000.00
./admin contract -n ${GETH_ADDRESS} -a 0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7 -m 1000.00
