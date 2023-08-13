# Constants for cross compilation and packaging.
icon = ./gui/assets/Icon.png
name = Macaw

# Default path to the go binary directory.
GOBIN ?= ~/go/bin/

binaries:
	docker buildx build -f deploy/docker/Dockerfile --output bin --target binaries .

build:
	go build -o bin/macaw ./cmd/macaw/main.go

run:
	go run ./cmd/macaw/main.go
	
bundle:
	${GOBIN}fyne bundle -package assets -name AppIcon ${icon} > ./gui/assets/bundled.go

check:
	gofmt -s -w ./..

	${GOBIN}gosec ./...

	${GOBIN}staticcheck -f stylish ./...

tests:
	go test ./...
