package socks5

import (
	"golang.org/x/net/context"
	"net"
	"testing"
)

func TestDNSResolver(t *testing.T) {
	d := DNSResolver{}
	ctx := context.Background()

	_, addr, err := d.Resolve(ctx, "localhost")
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if !addr.(*net.IPAddr).IP.IsLoopback() {
		t.Fatalf("expected loopback")
	}
}
