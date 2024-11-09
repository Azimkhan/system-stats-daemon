test:
	go test -race ./internal/... --count=1

.PHONY: test