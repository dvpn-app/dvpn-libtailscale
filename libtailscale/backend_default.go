// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !android

package libtailscale

func setPlatformProtectFunc(protect func(fd int) error) {
	// No-op on non-Android platforms
}

func clearPlatformProtectFunc() {
	// No-op on non-Android platforms
}
