BINARY := CanBook-backend-go
OUTPUT_DIR := build
MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

GOARCH = amd64

build:
	CGO_ENABLED=0 go build -o build/${BINARY}

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -o build/${BINARY}-linux-${GOARCH} .

mod:
	go mod download

lint:
	if ! [ -x $(GOPATH)/bin/golangci-lint ]; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.45.2; \
	fi
	golangci-lint run --concurrency 2

vet:
	go vet ./...

re:
	test -f build/${BINARY} && rm build/${BINARY} || true
	test -f build/${BINARY}-linux-${GOARCH} && rm build/${BINARY}-linux-${GOARCH} || true
	make build

.PHONY:	build mod test lint vet clean re
