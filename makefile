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
# Create new accounts (use 123 as passphrase)
#     geth account new --keystore zarf/ethereum/keystore
#
# For more ethereum calls look at the smart contract project
#
# go get github.com/ethereum/go-ethereum@master
# make deps-upgrade
# Look at go.mod for go-ethereum and cherry pick deps

GOLANG          := golang:1.23
NODE            := node:16
ALPINE          := alpine:3.20
CADDY           := caddy:2.8-alpine
KIND            := kindest/node:v1.31.0
BUSYBOX         := busybox:stable
GETH            := ethereum/client-go:stable
POSTGRES        := postgres:16.4

KIND_CLUSTER    := liars-game-cluster
NAMESPACE       := liars-system
APP             := engine
BASE_IMAGE_NAME := ardanlabs/liars
SERVICE_NAME    := engine
VERSION         := 0.0.1
SERVICE_IMAGE   := $(BASE_IMAGE_NAME)/$(SERVICE_NAME):$(VERSION)

# ==============================================================================
# Install dependencies

dev-setup:
	brew update
	brew list kind || brew install kind
	brew list kubectl || brew install kubectl
	brew list kustomize || brew install kustomize
	brew list ethereum || brew install ethereum
	brew list solidity || brew install solidity
	brew list pgcli || brew install pgcli
	brew list watch || brew instal watch

dev-docker:
	docker pull $(GOLANG) & \
	docker pull $(NODE) & \
	docker pull $(ALPINE) & \
	docker pull $(CADDY) & \
	docker pull $(KIND) & \
	docker pull $(BUSYBOX) & \
	docker pull $(GETH) & \
	docker pull $(POSTGRES) & \
	wait;

# ==============================================================================
# Game UI

game-tui1:
	go run app/cli/liars/main.go -a 0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7

game-tui2:
	go run app/cli/liars/main.go -a 0x8e113078adf6888b7ba84967f299f29aece24c55

game-tuio:
	go run app/cli/liars/main.go -a 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd

game-ui:
	go run app/services/ui/main.go

# ==============================================================================
# Building containers

build: game-engine

game-engine:
	docker build \
		-f zarf/docker/dockerfile.engine \
		-t $(SERVICE_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

# ==============================================================================
# Building contract

SOLC_EVM_VERSION := cancun

contract-build:
	solc --evm-version $(SOLC_EVM_VERSION) --abi business/contract/src/bank/bank.sol -o business/contract/abi/bank --overwrite
	solc --evm-version $(SOLC_EVM_VERSION) --bin business/contract/src/bank/bank.sol -o business/contract/abi/bank --overwrite
	abigen --bin=business/contract/abi/bank/Bank.bin --abi=business/contract/abi/bank/Bank.abi --pkg=bank --out=business/contract/go/bank/bank.go

# ==============================================================================
# Running from within k8s/kind
#
# To start the system for the first time, run these two commands:
#     make dev-up
#     make dev-update-apply

dev-up:
	kind create cluster \
		--image $(KIND) \
		--name $(KIND_CLUSTER) \
		--config zarf/k8s/dev/kind-config.yaml
	kubectl wait --timeout=120s --namespace=local-path-storage --for=condition=Available deployment/local-path-provisioner
	
	kind load docker-image $(BUSYBOX) --name $(KIND_CLUSTER)
	kind load docker-image $(GETH) --name $(KIND_CLUSTER)
	kind load docker-image $(POSTGRES) --name $(KIND_CLUSTER)

dev-down:
	kind delete cluster --name $(KIND_CLUSTER)
	rm -f /tmp/credentials.json

dev-load:
	kind load docker-image $(SERVICE_IMAGE) --name $(KIND_CLUSTER)

dev-deploy:
	@zarf/k8s/dev/geth/setup-contract-k8s

dev-deploy-force:
	@zarf/k8s/dev/geth/setup-contract-k8s force	

dev-apply:
	go build -o admin app/tooling/admin/main.go
 	
	kustomize build zarf/k8s/dev/geth | kubectl apply -f -
	kubectl wait --timeout=120s --namespace=$(NAMESPACE) --for=condition=Available deployment/geth

	@zarf/k8s/dev/geth/setup-contract-k8s.sh

	kustomize build zarf/k8s/dev/database | kubectl apply -f -
	kubectl rollout status --namespace=$(NAMESPACE) --watch --timeout=120s sts/database

	kustomize build zarf/k8s/dev/engine | kubectl apply -f -
	kubectl wait --timeout=120s --namespace=$(NAMESPACE) --for=condition=Available deployment/engine

dev-restart:
	kubectl rollout restart deployment engine --namespace=liars-system

dev-update: build dev-load dev-restart

dev-update-apply: build dev-load dev-apply

dev-logs-init:
	kubectl logs --namespace=$(NAMESPACE) -l app=$(APP) -f --tail=100 -c init-ge-migrate

dev-logs:
	kubectl logs --namespace=$(NAMESPACE) -l app=$(APP) --all-containers=true -f --tail=100 | go run app/tooling/logfmt/main.go

dev-logs-geth:
	kubectl logs --namespace=$(NAMESPACE) -l app=geth --all-containers=true -f --tail=1000

dev-status-all:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces

dev-status:
	watch -n 2 kubectl get pods -o wide --all-namespaces

dev-describe:
	kubectl describe nodes
	kubectl describe svc

dev-describe-deployment-engine:
	kubectl describe deployment --namespace=$(NAMESPACE) $(APP)

dev-describe-engine:
	kubectl describe pod --namespace=$(NAMESPACE) -l app=$(APP)

# ==============================================================================
# Administration

migrate:
	export SALES_DB_HOST=localhost; go run app/tooling/sales-admin/main.go migrate

pgcli:
	pgcli postgresql://postgres:postgres@localhost

# ==============================================================================
# Running tests within the local computer
# go install honnef.co/go/tools/cmd/staticcheck@latest
# go install golang.org/x/vuln/cmd/govulncheck@latest

test:
	go test -count=1 ./...
	go vet ./...
	staticcheck -checks=all ./...
	govulncheck ./...

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