OUTPUT_FILE = dist/confgen

VERSION ?= DEV

all: clean build

test-lite:
	@go test ./... -coverprofile lite.out

test-sprig:
	@go test -tags sprig ./... -coverprofile sprig.out

test: test-lite test-sprig

cover: test
cover:
	@cp lite.out combo.out
	@cat sprig.out | grep -v ^"mode: set" >> combo.out
	@go tool cover -func=combo.out

clean:
	@go clean
	@rm -f $(OUTPUT_FILE) $(OUTPUT_FILE)-sprig

build-lite:
	@mkdir -p dist
	go build -ldflags "-w -s -X main.version=${VERSION}" \
	         -o $(OUTPUT_FILE) \
	         main.go

build-sprig:
	@mkdir -p dist
	go build -ldflags "-w -s -X main.version=${VERSION}" \
			 -tags sprig \
	         -o $(OUTPUT_FILE)-sprig \
	         main.go

build: build-lite build-sprig

release:
	@VERSION=$(VERSION) docker run --rm --privileged \
  				-v $(PWD):/go/src/app \
  				-v /var/run/docker.sock:/var/run/docker.sock \
  				-w /go/src/app \
				-e VERSION \
				-e GORELEASER_CURRENT_TAG \
  				goreleaser/goreleaser --skip-publish --snapshot --rm-dist
