// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !android

package libtailscale

import (
	"log"
)

// initLogging initializes logging for non-Android platforms
func initLogging(appCtx AppContext) {
	// On non-Android platforms, we just use standard logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
