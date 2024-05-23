all: download vet fmt lint test

download:
	go mod download

vet:
	go vet ./...

lint:
	golangci-lint run

fmt:
	go fmt

test:
	go test ./... -v

cover:
	go test -cover -coverprofile=coverage.out -v ./...
	go tool cover -html=coverage.out

clean:
	git clean -fxd -e .idea
