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

GOLANG       := golang:1.21.6
NODE         := node:16
ALPINE       := alpine:3.19
CADDY        := caddy:2.6-alpine
KIND         := kindest/node:v1.29.0@sha256:eaa1450915475849a73a9227b8f201df25e55e268e5d619312131292e324d570
GETH         := ethereum/client-go:stable

KIND_CLUSTER := liars-game-cluster
VERSION      := 1.0

# ==============================================================================
# Install dependencies

dev-setup:
	brew update
	brew list kind || brew install kind
	brew list kubectl || brew install kubectl
	brew list kustomize || brew install kustomize
	brew list ethereum || brew install ethereum
	brew list solidity || brew install solidity

dev-docker:
	docker pull $(GOLANG)
	docker pull $(NODE)
	docker pull $(ALPINE)
	docker pull $(CADDY)
	docker pull $(KIND)
	docker pull $(GETH)

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

# all: game-engine ui
all: game-engine

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

dev-up:
	kind create cluster \
		--image $(KIND) \
		--name $(KIND_CLUSTER) \
		--config zarf/k8s/dev/kind-config.yaml
	kubectl wait --timeout=120s --namespace=local-path-storage --for=condition=Available deployment/local-path-provisioner
	
	kind load docker-image $(GETH) --name $(KIND_CLUSTER)

dev-down:
	kind delete cluster --name $(KIND_CLUSTER)
	rm -f /tmp/credentials.json

dev-load:
	kind load docker-image liarsdice-game-engine:$(VERSION) --name $(KIND_CLUSTER)
#	kind load docker-image liarsdice-game-ui:$(VERSION) --name $(KIND_CLUSTER)

dev-deploy:
	@zarf/k8s/dev/geth/setup-contract-k8s

dev-deploy-force:
	@zarf/k8s/dev/geth/setup-contract-k8s force	

dev-apply:
	go build -o admin app/tooling/admin/main.go

	kustomize build zarf/k8s/dev/geth | kubectl apply -f -
	kubectl wait --timeout=120s --namespace=liars-system --for=condition=Available deployment/geth

	@zarf/k8s/dev/geth/setup-contract-k8s.sh

	kustomize build zarf/k8s/dev/engine | kubectl apply -f -
	kubectl wait --timeout=120s --namespace=liars-system --for=condition=Available deployment/engine

# kustomize build zarf/k8s/dev/ui | kubectl apply -f -
# kubectl wait --timeout=120s --namespace=liars-system --for=condition=Available deployment/ui

dev-restart:
	kubectl rollout restart deployment engine --namespace=liars-system

dev-update: all dev-load dev-restart

dev-update-apply: all dev-load dev-apply

dev-logs:
	kubectl logs --namespace=liars-system -l app=engine --all-containers=true -f --tail=100 | go run app/tooling/logfmt/main.go

dev-logs-ui:
	kubectl logs --namespace=liars-system -l app=ui --all-containers=true -f --tail=100 | go run app/tooling/logfmt/main.go

dev-logs-geth:
	kubectl logs --namespace=liars-system -l app=geth --all-containers=true -f --tail=1000

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