## GNU Make operational configuration

MAKEFLAGS := --warn-undefined-variables --no-builtin-rules
SHELL := bash
.SHELLFLAGS := -euc

.DELETE_ON_ERROR:
.ONESHELL:


## Recipe parameters

SPEC ?= ./...
FLAGS ?= -debug
IMAGE ?= image.png

## External recipe definitions

run: build
	./blocky $(FLAGS) $(IMAGE) > image.svg

build:
	go build
.PHONY: build

test:
	go test $(SPEC)
.PHONY: test

sizes: build
	./blocky -keepInvisible $(IMAGE) > _image_pixels.svg
	gzip -c9 _image_pixels.svg > _image_pixels.svg.gz
	./blocky $(IMAGE) > _image_pixels_excludes.svg
	gzip -c9 _image_pixels_excludes.svg > _image_pixels_excludes.svg.gz
	./blocky -optimize rects $(IMAGE) > _image_rects_excludes.svg
	gzip -c9 _image_rects_excludes.svg > _image_rects_excludes.svg.gz
	du -b $(IMAGE) _image*
	rm _image*
.PHONY: sizes
.SILENT: sizes
