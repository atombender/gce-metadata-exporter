.PHONY: default
default: build

VERSION := 0.1.2
NAME := gce-metadata-exporter
ARCH := $(shell uname -m)

BUILD_DIR := $(PWD)/build
GOPATH := $(BUILD_DIR)/go
GO_PACKAGE_PATH := $(GOPATH)/src/github.com/atombender/gce-metadata-exporter
GO := env GOPATH=$(GOPATH) go

GO_SRC := $(shell find . -name '*.go' -type f | fgrep -v ./vendor/ | fgrep -v '${BUILD_DIR}')

.PHONY: build
build: $(BUILD_DIR)/gce-metadata-exporter

$(BUILD_DIR)/gce-metadata-exporter: $(GO_SRC)
	mkdir -p $(GOPATH)/src/github.com/atombender
	ln -sf $(PWD) $(GOPATH)/src/github.com/atombender/gce-metadata-exporter
	$(GO) build -o ${BUILD_DIR}/gce-metadata-exporter github.com/atombender/gce-metadata-exporter

.PHONY: docker-build
docker-build:
	docker build -t atombender/gce-metadata-exporter:$(VERSION) .
	docker tag atombender/gce-metadata-exporter:$(VERSION) atombender/gce-metadata-exporter:latest

.PHONY: docker-dist
docker-dist: docker-build
	docker push atombender/gce-metadata-exporter:$(VERSION)
	docker push atombender/gce-metadata-exporter:latest

.PHONY: dist
dist: build
	mkdir -p dist
	cp LICENSE README.md build/gce-metadata-exporter dist/
	tar cvfz $(NAME)-$(VERSION)-$(ARCH).tar.gz -C dist .
