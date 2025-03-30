# Makefile
GOEXEC = go
COVERAGE_REPORT = coverage.out
TEST_REPORT = report.out
TEST_FILES = ./internal/...
TEST_FILES_INTEGRATION = ./tests/...
COMPOSE_FILE = ./docker/docker-compose.yml
DOCKER_COMPOSE = docker compose -f "${COMPOSE_FILE}"
GOINSTALL = ${GOEXEC} install

tests:
	make test-unit
	make test-integration

test-unit:
	@echo "Running tests..."
	${GOEXEC} test -v ${TEST_FILES}

test-coverage:
	@echo "Running tests with coverage..."
	${GOEXEC} test -bench=. -json -coverprofile="${COVERAGE_REPORT}" ${TEST_FILES} > "${TEST_REPORT}"

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
	${GOEXEC} build -o bin/app cmd/main.go

tools:
	${GOINSTALL} go.uber.org/mock/mockgen@latest
