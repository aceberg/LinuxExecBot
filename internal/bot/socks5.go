package bot

import (
	"context"
	"net"
	"net/http"

	"golang.org/x/net/proxy"

	"github.com/aceberg/LinuxExecBot/internal/check"
)

func addSocks5(socks5addr string) *http.Client {

	dialer, err := proxy.SOCKS5("tcp", socks5addr, nil, proxy.Direct)
	check.IfError(err)

	transport := &http.Transport{}

	transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		return dialer.Dial(network, addr)
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	return httpClient
}
