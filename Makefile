
test-local:
	go test -race ./internal/... --count=1

test:
	@if docker compose -f deployment/docker-compose-test.yaml up --exit-code-from test; then \
		echo "Unit tests passed"; \
		docker compose -f deployment/docker-compose-test.yaml down; \
		exit 0; \
	else \
		echo "Unit tests failed"; \
		docker compose -f deployment/docker-compose-test.yaml down; \
		exit 1; \
	fi

.PHONY: test test-docker