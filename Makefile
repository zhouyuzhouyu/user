.PHONY: build

all: build

build:
	@sh -c "'$(CURDIR)/scripts/build.sh'"

generate:
	go generate ./...

upx:
	@upx zy-apilabs

.PHONY: clean
## clean: Clean build files
clean:
	@rm -rf zy-apilabs apiGrpc dist.tar.gz *.upx
.NOTPARALLEL:

help:
