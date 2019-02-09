all: test build

test:
	go vet ./...
	golint ./...
	go test ./...

build:
	go build .
	ls -ltra .

run:
	./api-server

