build: ## build
	go build .

lint:  ## basic lints
	golint ./...
	go vet ./...
	gofmt -w -s *.go */*.go

generate:  ## regenerate mapping
	go generate .

test:  ## tests
	go test ./...

clean:  ## clean up
	go clean ./...
	git gc

ci: build lint test  ## do what travis-ci does

# https://www.client9.com/self-documenting-makefiles/
help:
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
        printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF \
        }' $(MAKEFILE_LIST)
.DEFAULT_GOAL=help
.PHONY=help
