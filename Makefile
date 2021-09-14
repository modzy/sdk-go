.PHONY: all
all: fmt lint vet test

.PHONY: tools
tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.40.0

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run ./...	

.PHONY: test
test:
	go test -covermode=count -coverprofile=combined.coverprofile ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: build
build:
	go build ./...

.PHONY: clean
clean:
	find . -name \*.coverprofile -delete

.PHONY: testing
testing:
	@cd internal/testing && go run main.go