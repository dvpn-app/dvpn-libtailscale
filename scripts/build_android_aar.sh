#!/bin/bash

set -e

echo "Building Android AAR for dvpn-tailscale..."

export GOPATH=${GOPATH:-$(go env GOPATH)}
export PATH=$PATH:$GOPATH/bin

if ! command -v gomobile &> /dev/null; then
    echo "gomobile not found. Installing..."
    go install golang.org/x/mobile/cmd/gomobile@latest
fi

if [ ! -d "$GOPATH/pkg/gomobile" ]; then
    echo "Initializing gomobile..."
    gomobile init
fi

mkdir -p build/android

echo "Building AAR..."
gomobile bind -target=android \
    -androidapi 21 \
    -o build/android/libtailscale.aar \
    -v \
    ./libtailscale

echo "Build complete. AAR file is at: build/android/libtailscale.aar"