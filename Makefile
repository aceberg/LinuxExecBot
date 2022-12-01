mod:
	rm go.mod || true && \
	rm go.sum || true && \
	go mod init github.com/aceberg/LinuxExecBot && \
	go mod tidy

run:
	cd cmd/LinuxExecBot/ && \
	go run .

fmt:
	go fmt ./...

lint:
	golangci-lint run
	golint ./...

go-build:
	cd cmd/LinuxExecBot/ && \
	CGO_ENABLED=0 go build -o ../../tmp/LinuxExecBot .
