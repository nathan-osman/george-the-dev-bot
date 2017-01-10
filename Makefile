CWD = $(shell pwd)
PKG = github.com/nathan-osman/george-the-dev-bot

all: dist/george

dist/george: dist
	docker run \
	    --rm \
	    -e CGO_ENABLED=0 \
	    -v ${CWD}:/go/src/${PKG} \
	    -v ${CWD}/dist:/go/bin \
	    -w /go/src/${PKG} \
	    golang:latest \
	    go get ./...

dist:
	@mkdir dist

clean:
	@rm -rf dist

.PHONY: clean
