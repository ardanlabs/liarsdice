# The environment has three accounts all using this same passkey (123).
# Geth is started with address 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd and is used as the coinbase address.
# The coinbase address is the account to pay mining rewards to.
# The coinbase address is given a LOT of money to start.
#
# These are examples of what you can do in the attach JS environment.
# 	eth.getBalance("0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd") or eth.getBalance(eth.coinbase)
# 	eth.getBalance("0x8e113078adf6888b7ba84967f299f29aece24c55")
# 	eth.getBalance("0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7")
#   eth.sendTransaction({from:eth.coinbase, to:"0x8e113078adf6888b7ba84967f299f29aece24c55", value: web3.toWei(0.05, "ether")})
#   eth.sendTransaction({from:eth.coinbase, to:"0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7", value: web3.toWei(0.05, "ether")})
#   eth.blockNumber
#   eth.getBlockByNumber(8)
#   eth.getTransaction("0xaea41e7c13a7ea627169c74ade4d5ea86664ff1f740cd90e499f3f842656d4ad")
#
# make geth-deposit
# export GAME_CONTRACT_ID=0xeB380D740eC33ADf803abe0D6B14Ee29Ae6194a9
# ./admin -a 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd -m 1000.00
# ./admin -a 0x8e113078adf6888b7ba84967f299f29aece24c55 -m 1000.00
# ./admin -a 0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7 -m 1000.00
#
# Web3 API
# https://web3js.readthedocs.io/en/v1.7.4/

# ==============================================================================
# Install dependencies
# https://geth.ethereum.org/docs/install-and-build/installing-geth
# https://docs.soliditylang.org/en/v0.8.11/installing-solidity.html

GOLANG       := golang:1.19
NODE         := node:16
ALPINE       := alpine:3.16
CADDY        := caddy:2.6-alpine
KIND         := kindest/node:v1.25.3
GETH         := ethereum/client-go:stable
TELEPRESENCE := docker.io/datawire/tel2:2.9.2

dev.setup.mac.common:
	brew update
	brew list kind || brew install kind
	brew list kubectl || brew install kubectl
	brew list kustomize || brew install kustomize
	brew list ethereum || brew install ethereum
	brew list solidity || brew install solidity

dev.setup.mac: dev.setup.mac.common
	brew datawire/blackbird/telepresence || brew install datawire/blackbird/telepresence

dev.setup.mac.arm64: dev.setup.mac.common
	brew datawire/blackbird/telepresence-arm64 || brew install datawire/blackbird/telepresence-arm64

dev.docker:
	docker pull $(GOLANG)
	docker pull $(NODE)
	docker pull $(ALPINE)
	docker pull $(CADDY)
	docker pull $(KIND)
	docker pull $(GETH)
	docker pull $(TELEPRESENCE)

# ==============================================================================
# Game Engine and UI

game-up:
	CGO_ENABLED=0 go run app/services/engine/main.go | go run app/tooling/logfmt/main.go

game-tui1:
	CGO_ENABLED=0 go run app/cli/liars/main.go -a 0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7

game-tui2:
	CGO_ENABLED=0 go run app/cli/liars/main.go -a 0x8e113078adf6888b7ba84967f299f29aece24c55

game-tuio:
	CGO_ENABLED=0 go run app/cli/liars/main.go -a 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd

react-install:
	yarn --cwd app/services/ui/ install

app-ui: react-install
	yarn --cwd app/services/ui/ start

# ==============================================================================
# These commands build and deploy basic smart contract.

# This will compile the smart contract and produce the binary code. Then with the
# abi and binary code, a Go source code file can be generated for Go API access.
contract-build:
	solc --abi business/contract/src/bank/bank.sol -o business/contract/abi/bank --overwrite
	solc --bin business/contract/src/bank/bank.sol -o business/contract/abi/bank --overwrite
	abigen --bin=business/contract/abi/bank/Bank.bin --abi=business/contract/abi/bank/Bank.abi --pkg=bank --out=business/contract/go/bank/bank.go

# This will deploy the smart contract to the locally running Ethereum environment.
admin-build:
	CGO_ENABLED=0 go build -o admin app/tooling/admin/main.go

contract-deploy: contract-build admin-build
	./admin -d

# ==============================================================================
# These commands start the Ethereum node and provide examples of attaching
# directly with potential commands to try, and creating a new account if necessary.

# This is start Ethereum in developer mode. Only when a transaction is pending will
# Ethereum mine a block. It provides a minimal environment for development.
geth-up:
	geth --dev --ipcpath zarf/ethereum/geth.ipc --http.corsdomain '*' --http --allow-insecure-unlock --rpc.allow-unprotected-txs --http.vhosts=* --mine --miner.threads 1 --verbosity 5 --datadir "zarf/ethereum/" --unlock 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd --password zarf/ethereum/password

# This will signal Ethereum to shutdown.
geth-down:
	kill -INT $(shell ps | grep "geth " | grep -v grep | sed -n 1,1p | cut -c1-5)

# This will remove the local blockchain and let you start new.
geth-reset:
	rm -rf zarf/ethereum/geth/

# This is a JS console environment for making geth related API calls.
geth-attach:
	geth attach --datadir zarf/ethereum/

# This will add a new account to the keystore. The account will have a zero
# balance until you give it some money.
geth-new-account:
	geth --datadir zarf/ethereum/ account new

# This will deposit 1 ETH into the two extra accounts from the coinbase account.
# Do this if you delete the geth folder and start over or if the accounts need money.
geth-deposit:
	curl -H 'Content-Type: application/json' --data '{"jsonrpc":"2.0","method":"eth_sendTransaction", "params": [{"from":"0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd", "to":"0x8E113078ADF6888B7ba84967F299F29AeCe24c55", "value":"0x1000000000000000000"}], "id":1}' localhost:8545
	curl -H 'Content-Type: application/json' --data '{"jsonrpc":"2.0","method":"eth_sendTransaction", "params": [{"from":"0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd", "to":"0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7", "value":"0x1000000000000000000"}], "id":1}' localhost:8545
	./admin -a 0x8e113078adf6888b7ba84967f299f29aece24c55 -m 1000.00
	./admin -a 0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7 -m 1000.00

# ==============================================================================
# Running tests within the local computer
# go install honnef.co/go/tools/cmd/staticcheck@latest
# go install golang.org/x/vuln/cmd/govulncheck@latest

test:
	CGO_ENABLED=0 go test -count=1 ./...
	CGO_ENABLED=0 go vet ./...
	staticcheck -checks=all ./...
	govulncheck ./...

test-gui:
	yarn --cwd app/services/game/ test

# ==============================================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -v ./...
	go mod tidy
	go mod vendor

list:
	go list -mod=mod all

# ==============================================================================
# Building containers

# $(shell git rev-parse --short HEAD)
VERSION := 1.0

all: game-engine ui

game-engine:
	docker build \
		-f zarf/docker/dockerfile.engine \
		-t liarsdice-game-engine:$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

ui:
	docker -l debug build \
		-f zarf/docker/dockerfile.ui \
		-t liarsdice-game-ui:$(VERSION) \
		--progress=plain \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

# ==============================================================================
# Running from within k8s/kind

KIND_CLUSTER := liars-game-cluster

dev-up:
	kind create cluster \
		--image kindest/node:v1.25.3@sha256:f52781bc0d7a19fb6c405c2af83abfeb311f130707a0e219175677e366cc45d1 \
		--name $(KIND_CLUSTER) \
		--config zarf/k8s/dev/kind-config.yaml
	kubectl wait --timeout=120s --namespace=local-path-storage --for=condition=Available deployment/local-path-provisioner
	
	kind load docker-image $(TELEPRESENCE) --name $(KIND_CLUSTER)
	kind load docker-image $(GETH) --name $(KIND_CLUSTER)
	
	telepresence --context=kind-$(KIND_CLUSTER) helm install
	telepresence --context=kind-$(KIND_CLUSTER) connect

dev-down:
	telepresence quit -r -u
	kind delete cluster --name $(KIND_CLUSTER)

dev-load:
	kind load docker-image liarsdice-game-engine:$(VERSION) --name $(KIND_CLUSTER)
	kind load docker-image liarsdice-game-ui:$(VERSION) --name $(KIND_CLUSTER)

dev-apply:
	kustomize build zarf/k8s/dev/geth | kubectl apply -f -
	kubectl wait --timeout=120s --namespace=liars-system --for=condition=Available deployment/geth
		
	kustomize build zarf/k8s/dev/engine | kubectl apply -f -
	kubectl wait --timeout=120s --namespace=liars-system --for=condition=Available deployment/engine

	kustomize build zarf/k8s/dev/ui | kubectl apply -f -
	kubectl wait --timeout=120s --namespace=liars-system --for=condition=Available deployment/ui

dev-restart:
	kubectl rollout restart deployment sales --namespace=liars-system

dev-update: all dev-load dev-restart

dev-update-apply: all dev-load dev-apply

dev-logs:
	kubectl logs --namespace=liars-system --all-containers=true -f --tail=100 | go run app/tooling/logfmt/main.go

dev-logs-engine:
	kubectl logs --namespace=liars-system -l app=engine --all-containers=true -f --tail=100 | go run app/tooling/logfmt/main.go

dev-logs-ui:
	kubectl logs --namespace=liars-system -l app=ui --all-containers=true -f --tail=100 | go run app/tooling/logfmt/main.go

dev-status:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces

dev-describe:
	kubectl describe nodes
	kubectl describe svc

dev-describe-deployment-engine:
	kubectl describe deployment --namespace=liars-system engine

dev-describe-engine:
	kubectl describe pod --namespace=liars-system -l app=engine

dev-describe-deployment-ui:
	kubectl describe deployment --namespace=liars-system ui

dev-describe-ui:
	kubectl describe pod --namespace=liars-system -l app=ui

# ==============================================================================
# Docker Compose

fe-up:
	docker compose -f zarf/docker/compose.yml up

fe-down:
	docker compose -f zarf/docker/compose.yml down

fe-logs:
	docker compose -f zarf/docker/compose.yml logs