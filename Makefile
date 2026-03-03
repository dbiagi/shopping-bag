GOEXEC = go

# Build Variables
GOBUILD=$(GOEXEC) build
GOBUILD_VARS = CGO_ENABLED=1 GOOS=linux
GOBUILD_DEBUG_FLAGS = -gcflags="all=-N -l"
GOBUILD_BUILD_TAG = $(shell git describe --tags --abbrev=0 --always)
GOBUILD_BUILD_TIME = $(shell date +%s)
GOBUILD_GO_VERSION = $(shell go version | cut -c 14-)
GOBUILD_PARAMS = -tags musl -mod vendor -ldflags="-X 'main.buildVersion=$(GOBUILD_BUILD_TAG)' -X 'main.buildTime=$(GOBUILD_BUILD_TIME)' -X 'main.goVersion=$(GOBUILD_GO_VERSION)'"
GOBUILD_BINARY_DIR = ./bin
GOBUILD_MAIN_PATH = cmd/main.go
GOBUILD_BINARY_NAME = app

COVERAGE_REPORT = coverage.out
TEST_REPORT = report.out
TEST_FILES = ./internal/...
TEST_FILES_INTEGRATION = ./tests/...
COMPOSE_FILE = ./docker/docker-compose.yaml
DOCKER_COMPOSE = docker compose -f "${COMPOSE_FILE}"
GOINSTALL = ${GOEXEC} install

.PHONY: all tests test-unit test-integration test-coverage serve-dev infra-up infra-down deps build tools pipeline

tests:
	make test-unit
#	make test-integration

test-unit:
	@echo "Running tests..."
	${GOEXEC} test -v -coverprofile="${COVERAGE_REPORT}" ${TEST_FILES}

test-coverage:
	make test-unit
	# make test-integration >> ${TEST_REPORT}

test-integration:
	@echo "Running integration tests..."
	make infra-up
	${GOEXEC} test -v ${TEST_FILES_INTEGRATION}
	make infra-down

serve-dev:
	@echo "Starting server..."
	${GOEXEC} run cmd/main.go serve --env=dev

infra-up:
	@echo "Starting infrastructure..."
	${DOCKER_COMPOSE} up -d

infra-down:
	@echo "Stopping infrastructure..."
	${DOCKER_COMPOSE} down

deps:
	@echo "Installing dependencies..."
	${GOEXEC} mod tidy
	${GOEXEC} mod vendor

build:
	@echo "Building..."
	mkdir -p $(GOBUILD_BINARY_DIR)
	$(GOBUILD_VARS) $(GOBUILD) $(GOBUILD_PARAMS) -a -o $(GOBUILD_BINARY_DIR)/$(GOBUILD_BINARY_NAME) $(GOBUILD_MAIN_PATH);

lint:
	@echo "Running linter..."
	golangci-lint run ./...

tools:
	${GOINSTALL} go.uber.org/mock/mockgen@latest
	${GOINSTALL} github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.8.0

pipeline: tools lint tests
