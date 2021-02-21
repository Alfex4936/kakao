SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: tag update-pkg-cache

install:
	go install ./...

build:
	go build ./...

test:
	go test -v

bench:
	go test -bench=. -benchtime 100x

fmt:
	@gofmt -l -w $(SRC)

tags:
	find ./ -name '*.go' -print0 | xargs -0 gotags > TAGS

push:
	git push -u origin main

push-tag:
	git push origin main --tags

tag:
	git tag -a v$(VERSION) -m $(MSG)

delete-tag:
	git tag -d $(VERSION)
	git push origin :$(VERSION)

update-pkg-cache:
    GOPROXY=https://proxy.golang.org GO111MODULE=on go get github.com/$(USER)/$(PACKAGE)@v$(VERSION)