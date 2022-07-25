.PHONY: build clean tool lint help

all: build

vet:
	go vet ./...; true

# go install mvdan.cc/gofumpt@latest
fmt:
	gofumpt -l -w ./**/*.go

upgrade-go:
	@wget https://go.dev/dl/go1.18.4.linux-amd64.tar.gz
	@sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.18.4.linux-amd64.tar.gz
	rm -rf go1.18.4.linux-amd64.tar.gz

help:
	@echo "make vet: run specified go vet"
	@echo "make fmt: gofumpt -l -w ./pkg/"