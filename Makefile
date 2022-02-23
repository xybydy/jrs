TARGET=jrs

all: build

build:
	@go mod tidy
	@go build -o $(TARGET) .

clean:
	@rm -rf $(TARGET)
	@rm -rf build

install: build
	@mv $(TARGET) $(GOPATH)/bin/

golangci:
	@go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	@cd $(GOPATH)/src/github.com/golangci/golangci-lint/cmd/golangci-lint
	@go install -ldflags "-X 'main.version=$(git describe --tags)' -X 'main.commit=$(git rev-parse --short HEAD)' -X 'main.date=$(date)'"


