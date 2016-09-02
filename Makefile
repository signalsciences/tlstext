
all: lint test

lint:
	golint ./...
	go vet ./...
	gofmt -w -s *.go */*.go

# generate will regenerate the map between hex value and string
# it only rarely needs to be run
generate:
	go generate .

test:
	go test ./...
	misspell README.md *.go */*.go

clean:
	go clean ./...
	git gc

ci: lint test

docker-ci:
	docker run --rm \
		-v $(PWD):/go/src/github.com/client9/tlstext \
		-w /go/src/github.com/client9/tlstext \
		nickg/golang-dev-docker \
		make ci

.PHONY: ci docker-ci
