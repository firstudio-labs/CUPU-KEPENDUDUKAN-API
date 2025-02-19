BINARY_NAME=myapp

# Definition Go
GO=go
GOCMD=$(GO)

# Flag Go
GO_FLAGS=-v

# Running binary
build:
	$(GOCMD) build $(GO_FLAGS) -o $(BINARY_NAME) ./cmd

run:
	$(GOCMD) run ./cmd

# Jalankan semua unit test
test:
	$(GOCMD) test -v ./...

# Format kode Go
fmt:
	$(GOCMD) fmt ./...