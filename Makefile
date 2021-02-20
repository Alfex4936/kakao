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

.PHONY: tag

tag:
	git tag -a $(VERSION) -m $(MSG)

delete-tag:
	git tag -d $(VERSION)
	git push origin :$(VERSION)