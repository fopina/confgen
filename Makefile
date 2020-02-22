OUTPUT_FILE = dist/confgen

VERSION ?= DEV

all: clean build

test:
	@go test ./...

clean:
	@go clean
	@rm $(OUTPUT_FILE) -f

build:
	@mkdir -p dist
	go build -ldflags "-w -s -X main.version=${VERSION}" \
	         -o $(OUTPUT_FILE) \
	         main.go

release:
	@VERSION=$(VERSION) docker run --rm --privileged \
  				-v $(PWD):/go/src/pushit \
  				-v /var/run/docker.sock:/var/run/docker.sock \
  				-w /go/src/pushit \
				-e VERSION \
  				goreleaser/goreleaser --skip-publish --snapshot --rm-dist
