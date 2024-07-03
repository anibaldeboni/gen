.PHONY: build install_deps install clean

default: all

all: install_deps build

build: install_deps
	@go build -o bin/gen -ldflags "-s -w" ./main.go

install: build
	@cp bin/* $(HOME)/.local/bin/

install_deps:
	@go get -v ./...

clean:
	@rm -rf bin
