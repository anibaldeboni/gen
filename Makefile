INSTALL_PATH = $(HOME)/.local/bin
BUILD_PATH = ./bin
APP_NAME = gen
ARTIFACTS = $(BUILD_PATH)/$(APP_NAME)

.PHONY: build install_deps install clean

default: all

all: build

build: install_deps
	@go build -o $(ARTIFACTS) -ldflags "-s -w" ./main.go
	$(info Built to $(BUILD_PATH))

install: build
	@cp $(ARTIFACTS) $(INSTALL_PATH)
	$(info Installed to $(INSTALL_PATH)/$(APP_NAME))

install_deps:
	@go get -v ./...

clean:
	@rm -rf $(BUILD_PATH)
