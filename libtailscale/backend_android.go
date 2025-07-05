// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build android

package libtailscale

import (
	"tailscale.com/net/netns"
)

func init() {
	// Android-specific initialization
}

func setPlatformProtectFunc(protect func(fd int) error) {
	netns.SetAndroidProtectFunc(protect)
}

func clearPlatformProtectFunc() {
	netns.SetAndroidProtectFunc(nil)
}
