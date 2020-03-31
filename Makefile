.DEFAULT_GOAL = all

CMD_DIR=main.go

OUT_DIR=bin

# Name of actual binary to create
BINARY = jsont

.PHONY: all
all: test build

 # Run the application after building it first
.PHONY: run
run: build
	go run $(CMD_DIR) ;

# Build simply builds the application
.PHONY: build
build:
	env go build -o $(OUT_DIR)/${BINARY} $(CMD_DIR) ;

# Run unit tests
.PHONY: test
test:
	go test -v ./...

# Generate a coverage report
.PHONY: cover
cover:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

# Remove coverage report, binary, and Docker image
.SILENT: clean
.PHONY: clean
clean:
	go clean $(CMD_DIR)
	@rm -f $(OUT_DIR)/${BINARY}
	@rm -f coverage.out