package socks5

import (
	"net"

	"golang.org/x/net/context"
)

// NameResolver is used to implement custom name resolution
type NameResolver interface {
	Resolve(ctx context.Context, name string) (context.Context, net.Addr, error)
}

// DNSResolver uses the system DNS to resolve host names
type DNSResolver struct{}

func (d DNSResolver) Resolve(ctx context.Context, name string) (context.Context, net.Addr, error) {

	return d.ResolveIP(ctx, name)
}

func (d DNSResolver) ResolveIP(ctx context.Context, name string) (context.Context, *net.IPAddr, error) {
	addr, err := ResolveIP("ip", name)
	return ctx, addr, err
}

func ResolveIP(network, name string) (*net.IPAddr, error) {
	addr, err := net.ResolveIPAddr(network, name)
	if err != nil {
		ip := net.ParseIP(name)
		if ip != nil {
			return &net.IPAddr{
				IP: ip,
			}, nil
		}
		return &net.IPAddr{
			IP: net.ParseIP("127.0.0.1"),
		}, nil
	}
	return addr, err
}
