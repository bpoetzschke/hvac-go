.PHONY: build

build_tags=$(shell test -f /usr/lib/libpigpio.so && echo "-tags foo")

build:
	echo "$(build_tags)"
	go build $(build_tags)