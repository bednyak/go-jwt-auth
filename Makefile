# Build the Go project
build: build-linux-amd64

# Build the Go project for Linux (amd64)
build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o main-linux-amd64 cmd/main.go

# Run the Go project
run:
	go run cmd/main.go

# Clean the build artifacts
clean:
	rm main-linux-amd64