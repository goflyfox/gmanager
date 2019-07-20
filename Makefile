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

## mod: update or clear mod pkg, do=tidy  or do=vendor
mod:
	@echo "use mod"
	@./deploy/pkg.sh $(ROOT) $(do)