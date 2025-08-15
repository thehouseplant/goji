# Thanks to https://www.mohitkhare.com/blog/go-makefile/ for this!

export GO111MODULE=on

APP=goji
APP_EXECUTABLE="./out/$(APP)"

## Quality
# Runs a full quality check
check-quality:
	make lint
	mke fmt
	make vet

# Runs default go linting
lint:
	golangci-lint run --enable-all

# Checks for badly formatted code
vet:
	go vet ./...

# Format checker
fmt:
	go fmt ./...

# Fix dependencies in go.mod
tidy:
	go mod tidy

# Run tests and generate coverage report
test:
	make tidy
	make vendor
	go test -v -timeout 10m ./... -coverprofile=coverage.out -json > report.json

# Displays test coverage in HTML
coverage:
	make test
	go tool cover -html=coverage.out

## build
# Build the application
build:
	mkdir -p /out
	GOARCH=amd64 GOOS=darwin go build -o ./out/${APP_EXECUTABLE}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ./out/${APP_EXECUTABLE}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ./out/${APP_EXECUTABLE}-windows main.go

# Runs the go binary
run:
	make build
	chmod +x $(APP_EXECUTABLE)
	$(APP_EXECUTABLE)

# Clean binary and other generated files
clean:
	go clean
	rm -rf out/
	rm -f coverage*.out

# Put required packages to support builds and tests into /vendor
vendor:
	go mod vendor

## All
# Run all of the commands
all:
	make check-quality
	make test
	make build
