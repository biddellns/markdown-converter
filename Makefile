.PHONY: build-binary
build-binary:
	go build -o ./target/bin/markdown-converter ./cmd/cli/...

.PHONY: test
test:
	go test ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	go fmt ./...
