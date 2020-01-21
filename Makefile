## GNU Make operational configuration

MAKEFLAGS := --warn-undefined-variables --no-builtin-rules
SHELL := bash
.SHELLFLAGS := -euc

.DELETE_ON_ERROR:
.ONESHELL:


## Recipe parameters

SPEC ?= ./...


## External recipe definitions

run: build
	./blocky

build:
	go build
.PHONY: build

test:
	go test $(SPEC)
.PHONY: test
