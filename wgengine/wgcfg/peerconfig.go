// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package wgcfg

import (
	"net/netip"

	"github.com/sagernet/wireguard-go/device"
)

// PeerConfig is the subset of a peer's WireGuard configuration a caller
// needs to set up routing and encryption for a single peer, independent
// of the full Config/Peer bookkeeping wgengine uses internally.
type PeerConfig struct {
	AllowedIPs   []netip.Prefix
	PresharedKey device.NoisePresharedKey
}
