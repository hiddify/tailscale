// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package netmon

import (
	"context"
	"syscall"

	N "github.com/sagernet/sing/common/network"
	"github.com/sagernet/tailscale/types/nettype"
)

// Hooks carries the platform-specific dial/listen customizations a caller
// wants netmon (and the dialer it wraps) to use. Only Dialer is consumed by
// this package's New; Control and ListenPacket are passed through for the
// caller's own use elsewhere in the dial path.
type Hooks struct {
	Dialer       N.Dialer
	Control      func(network, address string, conn syscall.RawConn) error
	ListenPacket func(ctx context.Context, network, address string) (nettype.PacketConn, error)
}
