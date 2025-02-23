BINARY_NAME=myapp

# Definition Go
GO=go
GOCMD=$(GO)
.PHONY: test #biar tidak dianggap comant

# Flag Go
GO_FLAGS=-v

# Running binary
build:
	$(GOCMD) build $(GO_FLAGS) -o $(BINARY_NAME) ./cmd

# running main apps
run:
	$(GOCMD) run ./cmd

# running all unit test
test:
	$(GOCMD) test -v ./...

# Format kode Go
fmt:
	$(GOCMD) fmt ./...

seed:
	SEED_DATA=true $(GOCMD) run ./cmd/main.go
