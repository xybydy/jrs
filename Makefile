TARGET=jrs

all: build

deps: godep golangci

build: deps
	@dep ensure
	@go build -o $(TARGET) .

clean:
	@rm -rf $(TARGET)
	@rm -rf build

install: build
	@mv $(TARGET) $(GOPATH)/bin/

godep:
	@go get -u github.com/golang/dep/...

golangci:
	@go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	@cd $(GOPATH)/src/github.com/golangci/golangci-lint/cmd/golangci-lint
	@go install -ldflags "-X 'main.version=$(git describe --tags)' -X 'main.commit=$(git rev-parse --short HEAD)' -X 'main.date=$(date)'"

