
##help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

## run/api: runs the cmd/api application
.PHONY: run/api
run/api:
	go run ./cmd/api