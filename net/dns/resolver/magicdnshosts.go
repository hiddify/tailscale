// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package resolver

import (
	"net/netip"

	"github.com/sagernet/tailscale/util/dnsname"
)

// MagicDNSHosts is a lookup table of MagicDNS hostnames (self and peers,
// plus any configured ExtraRecords) to the IP addresses they resolve to,
// as computed by LocalBackend for the current netmap. It supports only
// exact-name lookups: this tailscale version does not generate wildcard
// or subdomain-eligible MagicDNS records, so SubdomainHost always reports
// no match.
type MagicDNSHosts map[dnsname.FQDN][]netip.Addr

// LookupHost returns the addresses, if any, that fqdn resolves to via
// MagicDNS.
func (m MagicDNSHosts) LookupHost(fqdn dnsname.FQDN) ([]netip.Addr, bool) {
	if m == nil {
		return nil, false
	}
	addrs, ok := m[fqdn]
	return addrs, ok
}

// SubdomainHost reports whether fqdn is itself a MagicDNS host whose
// subdomains should also resolve to it. No such hosts exist in this
// tailscale version, so this always returns false.
func (m MagicDNSHosts) SubdomainHost(fqdn dnsname.FQDN) bool {
	return false
}
