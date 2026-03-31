package bot

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/proxy"

	"github.com/aceberg/LinuxExecBot/internal/check"
)

func addSocks5(socks5addr string) *http.Client {
	var httpClient *http.Client
	var err error

	for {
		httpClient = createClient(socks5addr)
		_, err = httpClient.Get("https://api.telegram.org")
		if check.IfError(err) {
			log.Println("INFO: Waiting for SOCKS5 proxy...")
		} else {
			log.Println("INFO: Connected to SOCKS5 proxy:", socks5addr)
			break
		}

		time.Sleep(time.Duration(5) * time.Second)
	}

	return httpClient
}

func createClient(socks5addr string) *http.Client {

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
