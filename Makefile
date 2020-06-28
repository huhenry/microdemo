GO := go
COMMAND=microdemo
BIN_PATH=bin

DOCKER := docker
DOCKER_SUPPORTED_VERSIONS ?= 17|18|19

REGISTRY_PREFIX ?= henryhucn/microdemo
OS=linux
ARCH=amd64



ifeq ($(origin VERSION), undefined)
VERSION := $(shell git describe --dirty="-dev" --always --tags | sed 's/-/./2' | sed 's/-/./2' )
endif
export VERSION

GIT_SHA=$(shell git rev-parse HEAD)
DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
BUILD_INFO_IMPORT_PATH=github.com/huhenry/microdemo/pkg/version
BUILD_INFO=-ldflags "-X $(BUILD_INFO_IMPORT_PATH).commitSHA=$(GIT_SHA) -X $(BUILD_INFO_IMPORT_PATH).latestVersion=$(VERSION) -X $(BUILD_INFO_IMPORT_PATH).date=$(DATE)"


.PHONY: build
build:
	@echo "===========> Building binary $(COMMAND) $(VERSION) for $(OS) $(ARCH) $(BUILD_INFO)"
	@mkdir -p bin/$(OS)/$(ARCH)
	@CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) $(GO) build -o bin/$(COMMAND) $(BUILD_INFO) cmd/$(COMMAND)/main.go 

.PHONY: docker
docker: build
	@echo "===========> Building microdemo $(VERSION) docker image"
	@$(DOCKER) build --pull -t $(REGISTRY_PREFIX)/microdemo:$(VERSION) .
	@echo "===========> Pushing microdemo $(VERSION) image to $(REGISTRY_PREFIX)"
	@$(DOCKER) push $(REGISTRY_PREFIX)/microdemo:$(VERSION)
   
