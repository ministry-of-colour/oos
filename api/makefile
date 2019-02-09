all: test build

build: 
	$(MAKE) -C cmd/api-server

test:
	go vet ./...
	golint ./...
	go test ./...

run: test build
	$(MAKE) -C cmd/api-server run
