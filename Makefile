GIT_COMMIT_SHA := $(shell git rev-list -1 HEAD)
BUILD_TAGS = rocksdb,builtin_static
BUILD_LD_FLAGS = "-X github.com/iotaledger/wasp/packages/wasp.VersionHash=$(GIT_COMMIT_SHA)"

#
# You can override these e.g. as
#     make test TEST_PKG=./packages/vm/core/testcore/ TEST_ARG="-v --run TestAccessNodes"
#
TEST_PKG=./...
TEST_ARG=

all: build-lint

build:
	go build -o . -tags $(BUILD_TAGS) -ldflags $(BUILD_LD_FLAGS) ./...

build-windows:
	go build -o . -tags $(BUILD_TAGS) -ldflags $(BUILD_LD_FLAGS) -buildmode=exe ./...

build-lint: build lint

test-full: install
	go test -tags $(BUILD_TAGS),runheavy ./... --timeout 60m --count 1 -failfast

test: install
	go test -tags $(BUILD_TAGS) $(TEST_PKG) --timeout 30m --count 1 -failfast $(TEST_ARG)

test-short:
	go test -tags $(BUILD_TAGS) --short --count 1 -failfast ./...

install:
	go install -tags $(BUILD_TAGS) -ldflags $(BUILD_LD_FLAGS) ./...

install-windows:
	go install -tags $(BUILD_TAGS) -ldflags $(BUILD_LD_FLAGS) -buildmode=exe ./...

lint:
	golangci-lint run

gofumpt-list:
	gofumpt -l ./

docker-build:
	docker build \
		--build-arg BUILD_TAGS=${BUILD_TAGS} \
		--build-arg BUILD_LD_FLAGS='${BUILD_LD_FLAGS}' \
		.

.PHONY: all build build-windows build-lint test test-short test-full install install-windows lint gofumpt-list docker-build

