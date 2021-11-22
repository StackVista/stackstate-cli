.PHONY: build, clean

build:
	docker run --rm --privileged \
		-v ${PWD}:/go/src/github.com/user/repo \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-w /go/src/github.com/user/repo \
		goreleaser/goreleaser build --snapshot --rm-dist

.PHONY: test
test:
	go test ./...

clean:
	rm -rf ${TARGET_DIR}/
	rm -rf dist

dist: build
	goreleaser release --skip-publish --snapshot --rm-dist
