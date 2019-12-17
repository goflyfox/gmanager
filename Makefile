# note: call scripts from /deploy

# project name
PROJECTNAME=$(shell basename "$(PWD)")

# project path
ROOT=$(shell pwd)

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## vendor: go mod vendor
## tidy: go mod tidy
## build: go build -mod=vendor
## run: go run main.go
## mod: update or clear mod pkg, do=tidy  or do=vendor
mod:
	@echo "use mod"
	@./deploy/pkg.sh $(ROOT) $(do)

vendor:
	@echo "use mod vendor"
	@export GO111MODULE=on
	@export GOPROXY=https://goproxy.io
	@go mod vendor

tidy:
	@echo "use mod tidy"
	@export GO111MODULE=on
	@export GOPROXY=https://goproxy.io
	@go mod tidy

#go build -ldflags "-s -w"
build:
	@go build -mod=vendor

run:
	@echo "go run main.go"
	@go run main.go
