
.PHONY: all clean android-aar test build

all: build

build:
	go build ./...

test:
	go test ./...

android-aar:
	@echo "Building Android AAR..."
	@./scripts/build_android_aar.sh

clean:
	rm -rf build/
	go clean ./...

deps:
	go mod download
	go mod tidy

install-gomobile:
	go install golang.org/x/mobile/cmd/gomobile@latest
	gomobile init