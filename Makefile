
all: lint test

lint:
	golint ./...
	go vet ./...
	gofmt -w -s *.go */*.go

test:
	go test ./...
	misspell README.md *.go */*.go

clean:
	rm -f *~ */*~
	go clean ./...
