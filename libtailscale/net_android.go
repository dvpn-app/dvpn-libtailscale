// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build android

package libtailscale

import (
	"github.com/tailscale/wireguard-go/tun"
)

func createTUNFromFD(fd int) (tun.Device, error) {
	device, _, err := tun.CreateUnmonitoredTUNFromFD(fd)
	return device, err
}
