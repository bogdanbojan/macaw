# Constants for cross compilation and packaging.
icon = ./gui/assets/Icon.png
name = Macaw

# Default path to the go binary directory.
GOBIN ?= ~/go/bin/

build:
	go build -o bin/macaw ./cmd/main.go

run:
	go run ./cmd/main.go
	
bundle:
	${GOBIN}fyne bundle -package assets -name AppIcon ${icon} > internal/assets/bundled.go

check:
	gofmt -s -w ./..

	${GOBIN}gosec ./...

	${GOBIN}staticcheck -f stylish ./...

tests:
	go test ./...

docker-binaries:
	docker buildx build -f zarf/docker/Dockerfile --output bin --target binaries .
