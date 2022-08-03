.PHONY: build clean tool lint help

all: build
	@cd hello && protoc --go_out=plugins=grpc:. hello.proto