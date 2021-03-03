SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: tag update-pkg-cache

install:
	go install ./...

build:
	go build ./...

test:
	go test -v

gcc:
	go build -gcflags "-m -m -l" $(file).go

bench:
	go test -benchmem -run=^$$ -bench $(name) -benchtime $(run)x github.com/Alfex4936/kakao

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
	git tag -d v$(VERSION)
	git push origin :v$(VERSION)

update-pkg-cache:
    GOPROXY=https://proxy.golang.org GO111MODULE=on go get github.com/$(USER)/$(PACKAGE)@v$(VERSION)