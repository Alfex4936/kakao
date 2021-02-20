install:
	go install ./...

build:
	go build ./...

test:
	go test -v

fmt:
	gofmt -w *.go */*.go

tags:
	find ./ -name '*.go' -print0 | xargs -0 gotags > TAGS

push:
	git push -u origin main

push-tag:
	git push origin main --tags

.PHONY: tag update-pkg-cache

tag:
	git tag -a $(VERSION) -m $(MSG)

delete-tag:
	git tag -d $(VERSION)
	git push origin :$(VERSION)

update-pkg-cache:
    set GOPROXY=https://proxy.golang.org
	set GO111MODULE=on
	go get github.com/$(USER)/$(PACKAGE)@v$(VERSION)