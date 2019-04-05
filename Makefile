# Project related variables
PROJECTNAME=$(shell basename "$(PWD)")
M = $(shell printf "\033[34;1m▶\033[0m")
DONE="\n  $(M)  done ✨"

# Go related variables
GOBASE=$(shell pwd)
GOPATH=$(GOBASE)/vendor:$(GOBASE)
GOBIN=$(GOBASE)/bin
GO111MODULE=on
export GO111MODULE

.PHONY: help
help: Makefile
	@echo "\n Choose a command run in "$(PROJECTNAME)":\n"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## install: Install missing dependencies. Builds binary in ./bin
.PHONY: install
install:
	@echo "  $(M)  Checking if there is any missing dependencies...\n"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get $(get) ./...
	@echo $(DONE)

## build: Creates a docker image of the app
.PHONY: build
build:
	@echo "  $(M)  Building the 🐳 image...\n"
	docker build -t=$(PROJECTNAME) .
	@echo $(DONE)

## run: Runs the current docker image on port 8080
.PHONY: run
run:
	@echo "  $(M)  Running the 🐳 image...\n"
	docker run -it -p 8080:8080 --rm --name app -t $(PROJECTNAME)
	@echo $(DONE)

## clean: Clean build files. Runs `go clean` internally
.PHONY: clean
clean:
	@echo "  $(M)  Cleaning build cache..."
	go clean ./...
	rm -rf bin
	rm -rf cp.out
	docker image rm $(PROJECTNAME)
	@echo $(DONE)

## fmt: Runs gofmt on all source files
.PHONY: fmt
fmt:
	@echo "  $(M) 🏃 gofmt..."
	@ret=0 && for d in $$(go list -f '{{.Dir}}' ./...); do \
		gofmt -l -w $$d/*.go || ret=$$? ; \
	 done ; exit $$ret
	@echo $(DONE)

## test: Runs all the tests.
.PHONY: test
test:
	@echo "  $(M)  🏃 all the tests...\n"
	cd "$$GOBASE"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test ./...
	@echo $(DONE)

## coverage: Tests code coverage
.PHONY: coverage
coverage:
	@echo "  $(M)  👀 testing code coverage...\n"
	cd "$$GOBASE"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test ./... -coverprofile cp.out
	@echo $(DONE)

## missing: Displays lines of code missing from coverage
.PHONY: missing
missing:
	@echo "  $(M)  👀 missing coverage...\n"
	grep -v -e " 1$" ./cp.out

