## GNU Make operational configuration

MAKEFLAGS := --warn-undefined-variables --no-builtin-rules
SHELL := bash
.SHELLFLAGS := -euc

.DELETE_ON_ERROR:
.ONESHELL:


## Recipe parameters

SPEC ?= ./...
FLAGS ?= -debug

## External recipe definitions

run: build
	./blocky $(FLAGS) image.png > image.svg

build:
	go build
.PHONY: build

test:
	go test $(SPEC)
.PHONY: test

sizes: build
	./blocky image.png > image_pfp.svg
	gzip -c9 image_pfp.svg > image_pfp.svg.gz
	du -b image*
	rm image_*
.PHONY: sizes
.SILENT: sizes
