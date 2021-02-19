install:
	go install ./...

build:
	go build ./...

fmt:
	gofmt -w *.go */*.go

tags:
	find ./ -name '*.go' -print0 | xargs -0 gotags > TAGS

push:
	git push -u origin main

