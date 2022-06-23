.PHONY: build-binary
build-binary:
	go build -o ./target/bin/markdown-converter ./cmd/cli/...

