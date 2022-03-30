.PHONY: build, clean

build:
	goreleaser build --snapshot --rm-dist

.PHONY: test
test:
	go test ./...

clean:
	rm -rf dist

dist: build
	goreleaser release --skip-publish --snapshot --rm-dist
