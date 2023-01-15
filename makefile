# To start the system for the first time, run these two commands:
#     make dev-up
#     make dev-update-apply
# Expect the building of the FE to take a wee bit of time :(
#
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
# ./admin contract -a 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd -m 1000.00
# ./admin contract -a 0x8e113078adf6888b7ba84967f299f29aece24c55 -m 1000.00
# ./admin contract -a 0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7 -m 1000.00
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
TELEPRESENCE := docker.io/datawire/tel2:2.10.1

dev.setup.mac.common:
	brew update
	brew list kind || brew install kind
	brew list kubectl || brew install kubectl
	brew list kustomize || brew install kustomize
	brew list ethereum || brew install ethereum
	brew list solidity || brew install solidity
	brew list vault || brew install vault

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
# Game UI

game-tui1:
	go run app/cli/liars/main.go -a 0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7

game-tui2:
	go run app/cli/liars/main.go -a 0x8e113078adf6888b7ba84967f299f29aece24c55

game-tuio:
	go run app/cli/liars/main.go -a 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd

react-install:
	yarn --cwd app/services/ui/ install

app-ui: react-install
	yarn --cwd app/services/ui/ start

# ==============================================================================
# Building containers
#
# The new docker buildx build system is required for these docker build commands On systems other than Docker Desktop
# buildx is not the default build system. You will need to enable it with: docker buildx install

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
	docker build \
		-f zarf/docker/dockerfile.ui \
		-t liarsdice-game-ui:$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

# ==============================================================================
# Running from within k8s/kind
#
# To start the system for the first time, run these two commands:
#     make dev-up
#     make dev-update-apply
# Expect the building of the FE to take a wee bit of time :(

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
	telepresence quit -s
	kind delete cluster --name $(KIND_CLUSTER)
	rm -f /tmp/credentials.json

dev-load:
	kind load docker-image liarsdice-game-engine:$(VERSION) --name $(KIND_CLUSTER)
	kind load docker-image liarsdice-game-ui:$(VERSION) --name $(KIND_CLUSTER)

dev-deploy:
	@zarf/k8s/dev/geth/setup-contract-k8s

dev-deploy-force:
	@zarf/k8s/dev/geth/setup-contract-k8s force	

dev-apply:
	go build -o admin app/tooling/admin/main.go

	kustomize build zarf/k8s/dev/vault | kubectl apply -f -

	@zarf/k8s/dev/vault/initialize-vault.sh

	kustomize build zarf/k8s/dev/geth | kubectl apply -f -
	kubectl wait --timeout=120s --namespace=liars-system --for=condition=Available deployment/geth

	@zarf/k8s/dev/geth/setup-contract-k8s.sh

	kustomize build zarf/k8s/dev/engine | kubectl apply -f -
	kubectl wait --timeout=120s --namespace=liars-system --for=condition=Available deployment/engine

	kustomize build zarf/k8s/dev/ui | kubectl apply -f -
	kubectl wait --timeout=120s --namespace=liars-system --for=condition=Available deployment/ui

dev-restart:
	kubectl rollout restart deployment engine --namespace=liars-system

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
# Running tests within the local computer
# go install honnef.co/go/tools/cmd/staticcheck@latest
# go install golang.org/x/vuln/cmd/govulncheck@latest

test:
	go test -count=1 ./...
	go vet ./...
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