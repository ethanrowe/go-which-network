arch?=amd64
today=$(shell date +%Y%m%d-%T)
commit=$(shell git rev-parse HEAD)
build_image?=golang:1.10
repo_owner?=ethanrowe
source_loc?=/go/src/github.com/$(repo_owner)
project_name?=which-network
build_path?=$(source_loc)/$(project_name)

.PHONY: clean version

build: build-mac build-linux

version:
	@echo $(today)-$(commit)

native-build:
	go build -ldflags='-X main.build_version=$(today)-$(commit)' -o which-network .

build-mac:
	docker run \
  --rm \
  -it \
  -e GOOS=darwin \
  -e GOARCH=$(arch) \
  -v $(shell pwd):$(build_path) \
  -w $(build_path) \
  $(build_image) \
  /bin/bash -c "go build -ldflags='-X main.build_version=$(today)-$(commit)' -o which-network-mac ."

build-linux:
	docker run \
  --rm \
  -it \
  -e GOOS=linux \
  -e GOARCH=$(arch) \
  -v $(shell pwd):$(build_path) \
  -w $(build_path) \
  $(build_image) \
  /bin/bash -c "go build -ldflags='-X main.build_version=$(today)-$(commit)' -o which-network-linux ."

clean:
	-rm -rf which-network-mac
	-rm -rf which-network-linux
	-rm -rf which-network

