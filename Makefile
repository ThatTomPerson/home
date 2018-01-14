PACKER := $(GOPATH)/bin/gokr-packer


test: generate
	go test ./...
.PHONY: test

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

pkg/*:
	go run $@/*.go
.PHONY: pkg/*
