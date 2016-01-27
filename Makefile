
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
	git gc

ci: lint test

docker-ci:
	docker run --rm \
		-e COVERALLS_REPO_TOKEN=$(COVERALLS_REPO_TOKEN) \
		-v $(PWD):/go/src/github.com/client9/tlstext \
		-w /go/src/github.com/client9/tlstext \
		nickg/golang-dev-docker \
		make ci

.PHONY: ci docker-ci
