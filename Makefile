PKGS = $(shell go list ./... | grep -v /test)

build-server:
	CGO_ENABLED=0 go build -o ./build/server ./cmd/server
.PHONY: build-server

build: build-server
.PHONY: build

lint:
	golint $(PKGS) 
.PHONY: lint

test-unit:
	go test --race --cover -v $(PKGS)
.PHONY: test-unit

test: test-unit
.PHONY: test

test-short:
	go test --race --short $(PKGS)
.PHONY: test-short