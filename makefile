# The environment has three accounts all using this same passkey (123).
# Geth is started with address 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd and is used as the coinbase address.
# The coinbase address is the account to pay mining rewards to.
# The coinbase address is given a LOT of money to start.
#
# Testing running system
# For testing a simple query on the system. Don't forget to `make seed` first.
# curl --user "admin@example.com:gophers" http://localhost:3000/v1/users/token
# export TOKEN="COPY TOKEN STRING FROM LAST CALL"
# curl -H "Authorization: Bearer ${TOKEN}" http://localhost:3000/v1/users/1/2
#
# Test debug endpoints
# curl http://localhost:4000/debug/liveness
# curl http://localhost:4000/debug/readiness
#
# Running pgcli client for database
# brew install pgcli
# pgcli postgresql://postgres:postgres@localhost

# ==============================================================================
# Install dependencies

dev.setup.mac:
	brew update
	brew list kind || brew install kind
	brew list kubectl || brew install kubectl
	brew list kustomize || brew install kustomize


# ==============================================================================
# Building containers

# $(shell git rev-parse --short HEAD)
VERSION := 1.0

build-engine:
	docker build \
		-f zarf/docker/dockerfile.engine \
		-t liars-game-engine:$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

# ==============================================================================
# Running from within k8s/kind

KIND_CLUSTER := ardan-engine-cluster

# Upgrade to latest Kind: brew upgrade kind
# For full Kind v0.14 release notes: https://github.com/kubernetes-sigs/kind/releases/tag/v0.14.0
# The image used below was copied by the above link and supports both amd64 and arm64.

kind-up:
	kind create cluster \
		--image kindest/node:v1.24.0@sha256:0866296e693efe1fed79d5e6c7af8df71fc73ae45e3679af05342239cdc5bc8e \
		--name $(KIND_CLUSTER) \
		--config zarf/k8s/kind/kind-config.yaml
	kubectl config set-context --current --namespace=engine-system

kind-down:
	kind delete cluster --name $(KIND_CLUSTER)

kind-load:
	cd zarf/k8s/kind/engine-pod; kustomize edit set image engine-image=liars-game-engine:$(VERSION)
	kind load docker-image liars-game-engine:$(VERSION) --name $(KIND_CLUSTER)

kind-apply:
	kustomize build zarf/k8s/kind/database-pod | kubectl apply -f -
	kubectl wait --namespace=database-system --timeout=120s --for=condition=Available deployment/database-pod
	kustomize build zarf/k8s/kind/engine-pod | kubectl apply -f -

kind-services-delete:
	kustomize build zarf/k8s/kind/engine-pod | kubectl delete -f -
	kustomize build zarf/k8s/kind/database-pod | kubectl delete -f -

kind-restart:
	kubectl rollout restart deployment engine-pod

kind-update: build-engine kind-load kind-restart

kind-update-apply: build-engine kind-load kind-apply

kind-logs:
	kubectl logs -l app=engine --all-containers=true -f --tail=100 | go run app/tooling/logfmt/main.go

kind-logs-engine:
	kubectl logs -l app=engine --all-containers=true -f --tail=100 | go run app/tooling/logfmt/main.go -service=engine-API

kind-logs-db:
	kubectl logs -l app=database --namespace=database-system --all-containers=true -f --tail=100

kind-status:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces

kind-status-engine:
	kubectl get pods -o wide --watch --namespace=engine-system

kind-status-db:
	kubectl get pods -o wide --watch --namespace=database-system

kind-describe:
	kubectl describe nodes
	kubectl describe svc
	kubectl describe pod -l app=engine

kind-describe-deployment:
	kubectl describe deployment engine-pod

kind-describe-replicaset:
	kubectl get rs
	kubectl describe rs -l app=engine

kind-events:
	kubectl get ev --sort-by metadata.creationTimestamp

kind-events-warn:
	kubectl get ev --field-selector type=Warning --sort-by metadata.creationTimestamp

kind-context-engine:
	kubectl config set-context --current --namespace=engine-system

kind-shell:
	kubectl exec -it $(shell kubectl get pods | grep engine | cut -c1-26) --container engine-api -- /bin/sh

kind-database:
	# ./admin --db-disable-tls=1 migrate
	# ./admin --db-disable-tls=1 seed

# ==============================================================================
# Administration

migrate:
	go run app/tooling/engine-admin/main.go migrate

seed: migrate
	go run app/tooling/engine-admin/main.go seed

# ==============================================================================
# Running tests within the local computer

test:
	go test ./... -count=1
	staticcheck -checks=all ./...

# ==============================================================================
# Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -v ./...
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache

list:
	go list -mod=mod all


# ==============================================================================
# These commands build and deploy basic smart contract.

# This will compile the smart contract and produce the binary code. Then with the
# abi and binary code, a Go source code file can be generated for Go API access.
contract-build:
	solc --abi contract/sol/src/contract.sol -o contract/sol/abi --overwrite
	solc --bin contract/sol/src/contract.sol -o contract/sol/abi --overwrite
	abigen --bin=contract/sol/abi/LiarsDice.bin --abi=contract/sol/abi/LiarsDice.abi --pkg=contract --out=contract/sol/go/contract.go

# This will deploy the smart contract to the locally running Ethereum environment.
contract-deploy: contract-build
	go run contract/cmd/deploy/main.go

contract-transfer:
	go run contract/cmd/transfer/main.go

contract-trancheck:
	go run contract/cmd/trancheck/main.go

contract-newgame:
	go run contract/cmd/start/main.go

contract-bet:
	go run contract/cmd/bet/main.go

contract-deposit:
	go run contract/cmd/deposit/main.go

contract-withdraw:
	go run contract/cmd/withdraw/main.go

contract-endgame:
	go run contract/cmd/end/main.go


# ==============================================================================
# These commands start the Ethereum node and provide examples of attaching
# directly with potential commands to try, and creating a new account if necessary.

# This is start Ethereum in developer mode. Only when a transaction is pending will
# Ethereum mine a block. It provides a minimal environment for development.
geth-up:
	geth --dev --ipcpath zarf/ethereum/geth.ipc --http --allow-insecure-unlock --rpc.allow-unprotected-txs --mine --miner.threads 1 --verbosity 5 --datadir "zarf/ethereum/" --unlock 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd --password zarf/ethereum/password

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
	curl -H 'Content-Type: application/json' --data '{"jsonrpc":"2.0","method":"eth_sendTransaction", "params": [{"from":"0x8e113078adf6888b7ba84967f299f29aece24c55", "to":"0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7", "value":"0x1000000000000000000"}], "id":1}' localhost:8545
	curl -H 'Content-Type: application/json' --data '{"jsonrpc":"2.0","method":"eth_sendTransaction", "params": [{"from":"0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd", "to":"0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7", "value":"0x1000000000000000000"}], "id":1}' localhost:8545
