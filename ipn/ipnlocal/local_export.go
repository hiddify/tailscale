package ipnlocal

import (
	"net/netip"
	"sync/atomic"

	"github.com/sagernet/tailscale/net/dns/resolver"
	"github.com/sagernet/tailscale/tailcfg"
	"github.com/sagernet/tailscale/types/netmap"
	"github.com/sagernet/tailscale/version"
	"github.com/sagernet/tailscale/wgengine"
	"github.com/sagernet/tailscale/wgengine/filter"
)

func (b *LocalBackend) ExportFilter() *atomic.Pointer[filter.Filter] {
	return &b.currentNode().filterAtomic
}

func (b *LocalBackend) ExportEngine() wgengine.Engine {
	return b.e
}

func (b *LocalBackend) NetMapNoPeers() *netmap.NetworkMap {
	return b.NetMap()
}

func (b *LocalBackend) PeerForIP(ip netip.Addr) (wgengine.PeerForIP, bool) {
	return b.e.PeerForIP(ip)
}

// ExportMagicDNSHosts returns the current MagicDNS hostname-to-address
// table (self, peers, and any configured ExtraRecords), matching what
// authReconfig programs into the DNS manager.
func (b *LocalBackend) ExportMagicDNSHosts() resolver.MagicDNSHosts {
	b.mu.Lock()
	prefs := b.pm.CurrentPrefs()
	keyExpired := b.keyExpired
	cn := b.currentNode()
	b.mu.Unlock()
	dcfg := cn.dnsConfigForNetmap(prefs, keyExpired, version.OS())
	if dcfg == nil {
		return nil
	}
	return resolver.MagicDNSHosts(dcfg.Hosts)
}

// ExportSelf returns the current node's own NodeView from the netmap.
func (b *LocalBackend) ExportSelf() tailcfg.NodeView {
	return b.currentNode().Self()
}

// ExportPeers returns the current peers from the netmap.
func (b *LocalBackend) ExportPeers() []tailcfg.NodeView {
	return b.currentNode().Peers()
}

func (b *LocalBackend) SetExternalSSHHostKeys(keys []string) {
	b.mu.Lock()
	b.externalSSHHostKeys = keys
	if b.hostinfo != nil {
		b.hostinfo.SSH_HostKeys = keys
	}
	b.mu.Unlock()
	b.doSetHostinfoFilterServices()
}
