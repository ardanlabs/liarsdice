#!/usr/bin/env bash

get_pod_status () {
  kubectl -n liars-system get pods vault-0 > /dev/null 2>&1
  RET=${?}

  if [ "${RET}" -ne 0 ]; then
    echo "Not running"
    exit
  fi

  kubectl -n liars-system get pods -o jsonpath='{.items[0].status.phase}'
}

echo "Waiting for vault pod to enter the Running state"
while [ "$(get_pod_status)" != "Running" ]; do
  sleep 1
done

echo "Initializing vault"
./admin vault init
RET=${?}

if [ "${RET}" -ne 0 ]; then
  exit "${RET}"
fi

echo "Adding pem keys"
./admin vault add-keys
