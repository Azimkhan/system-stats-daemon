
build:
	go build -o bin/systemstats_server ./cmd/server/
	go build -o bin/systemstats_client ./cmd/client/

generate:
	mkdir -p gen/systemstats/pb
	protoc \
		--proto_path=api \
		--go_out=gen/systemstats/pb \
		--go-grpc_out=gen/systemstats/pb \
		api/*.proto

test-local:
	go test -race ./internal/... --count=1

test:
	@if docker compose -f deployment/docker-compose-test.yaml up --build --exit-code-from test; then \
		echo "Unit tests passed"; \
		docker compose -f deployment/docker-compose-test.yaml down; \
		exit 0; \
	else \
		echo "Unit tests failed"; \
		docker compose -f deployment/docker-compose-test.yaml down; \
		exit 1; \
	fi

integration-test:
	@if docker compose -f deployment/docker-compose-integration-test.yaml up --build --exit-code-from integration-test; then \
		echo "Integration tests passed"; \
		docker compose -f deployment/docker-compose-integration-test.yaml down; \
		exit 0; \
	else \
		echo "Integration tests failed"; \
		docker compose -f deployment/docker-compose-integration-test.yaml down; \
		exit 1; \
	fi

.PHONY: test test-docker build generate test-local integration-test