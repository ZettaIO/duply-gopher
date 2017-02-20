BUILDTAGS=
GOOS?=linux
REPO_INFO=$(shell git config --get remote.origin.url)

ifndef COMMIT
  COMMIT := git-$(shell git rev-parse --short HEAD)
endif

build: clean
	CGO_ENABLED=0 GOOS=${GOOS} go build -a -installsuffix cgo \
		-ldflags "-s -w -X ${PKG}/pkg/version.RELEASE=${RELEASE} -X ${PKG}/pkg/version.COMMIT=${COMMIT} -X ${PKG}/pkg/version.REPO=${REPO_INFO}" \
		-o duply-go cmd/duply-gopher/main.go

clean:
	rm -f duply-go
