MAKEFLAGS += --silent

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

PATAKA_DATABASE_URL="postgres://${PATAKA_DATABASE_USER}:${PATAKA_DATABASE_PASSWORD}@${PATAKA_DATABASE_HOST}:${PATAKA_DATABASE_PORT}/${PATAKA_DATABASE_NAME}?sslmode=${PATAKA_DATABASE_SSL_MODE}"

# Repository Manager
# Run server
.PHONY: develop
develop:
	go run main.go serve

mock: format
	rm -rf mocks
	mockery --all --keeptree --dir internal --output internal/mocks
	mockery --all --output mocks/google.golang.org/grpc/grpclog --srcpkg google.golang.org/grpc/grpclog

format:
	go fmt ./...

test: format
	gotestsum --format testname --junitfile junit.xml -- -coverprofile=coverage.lcov.info -covermode count ./...
	gocover-cobertura < coverage.lcov.info > coverage.xml
	gototal-cobertura < coverage.xml

# Prepare for development
.PHONY: prepare
prepare:
	go install github.com/vektra/mockery/v2@latest 1> /dev/null
	go install gotest.tools/gotestsum@latest 1> /dev/null
	go install github.com/boumenot/gocover-cobertura@latest 1> /dev/null
	go install github.com/ggere/gototal-cobertura@latest 1> /dev/null

# Download dependency
.PHONY: mod
mod:
	go mod tidy -compat=1.17
	go mod vendor

# Database Management
# Create database
.PHONY: createdb
createdb:
	createdb "${PATAKA_DATABASE_NAME}"

# Drop database
.PHONY: dropdb
dropdb:
	dropdb "${PATAKA_DATABASE_NAME}"

# Migrate database
.PHONY: migratedb
migratedb:
	migrate --path=db/migrations/ \
			--database ${PATAKA_DATABASE_URL} up

# Rollback database
.PHONY: rollbackdb
rollbackdb:
	echo "y" | migrate --path=db/migrations/ \
			--database ${PATAKA_DATABASE_URL} down

# Create database migration
migration:
	$(eval timestamp := $(shell date +%s))
	touch db/migrations/$(timestamp)_${name}.up.sql
	touch db/migrations/$(timestamp)_${name}.down.sql
