# Constants for cross compilation and packaging.
appID = github.com.bogdanbojan.macaw
app-windows = macaw-windows
app-darwin = macaw-darwin
app-linux = macaw-linux
icon = ./Icon.png
name = Macaw

# Default path to the go binary directory.
GOBIN ?= ~/go/bin/

bundle:
	# Bundle the correct logo. 
	${GOBIN}fyne bundle -package assets -name AppIcon ${icon} > internal/assets/bundled.go

check:
	# Check the whole codebase for misspellings.
	${GOBIN}misspell -w .

	# Run full formating on the code.
	gofmt -s -w .

	# Check the whole program for security issues.
	${GOBIN}gosec ./...

	# Run staticcheck on the codebase.
	${GOBIN}staticcheck -f stylish ./...

darwin:
	${GOBIN}fyne-cross darwin -arch amd64 -app-id ${appID} -icon ${icon} -output ${name} -name ${app-darwin}

linux:
	${GOBIN}fyne-cross linux -arch amd64 -app-id ${appID} -icon ${icon} -name ${app-linux}

windows:
	${GOBIN}fyne-cross windows -arch amd64 -app-id ${appID} -icon ${icon} -name ${app-windows} 

cross-compile: darwin linux windows

build:
	go build -o bin/
run:
	go run .
tests:
	go test ./...

docker-binaries:
	docker buildx build --output bin --target binaries .
