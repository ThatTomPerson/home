PACKER := $(GOPATH)/bin/gokr-packer

all: build
.PHONY:all

test:
	go test ./...
.PHONY: test

install:
	go get -t ./...
.PHONY: install

clean:
	rm -rf internal/api
.PHONY: clean

$(PACKER):
	go get -u github.com/gokrazy/tools/cmd/gokr-packer

generate: clean
	protoc --go_out=plugins=micro:./internal ./api/**/*.proto
	go generate ./...
.PHONY: generate

update: generate $(PACKER)
	gokr-packer -update="yes" -hostname="krazy" ./pkg/*
.PHONY: update

build: test pkg/*
.PHONY: build

pkg/*: bin = $(word 2,$(subst /, ,$@))
pkg/*:
	go build -o ./dist/$(bin) ./$@
.PHONY: pkg/*
