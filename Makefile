.SILENT:
.IGNORE:
.DEFAULT_GOAL := build

.PHONY: build
build:
	go build -o ./build/mkpkg ./cmd/mkpkg

.PHONY: install
install:
	mkdir -p /usr/local/bin
	mv ./build/mkpkg /usr/local/bin
