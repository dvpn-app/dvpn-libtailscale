// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !android

package libtailscale

import (
	"errors"
	"github.com/tailscale/wireguard-go/tun"
)

func createTUNFromFD(fd int) (tun.Device, error) {
	return nil, errors.New("CreateUnmonitoredTUNFromFD is only available on Android")
}
