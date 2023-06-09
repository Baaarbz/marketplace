acceptance-test:
	go test -v ./tests/acceptance/...

contract-test:
	go generate ./...
	go test -v ./tests/contract/...

unit-test:
	go generate ./...
	go test -v ./internal/...
	go test -v ./cmd/...
	go test -v ./pkg/...