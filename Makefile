.PHONY: build

build_tags=$(shell test -f /usr/lib/libpigpio.so && echo "-tags pigpio")

build:
	go build $(build_tags)