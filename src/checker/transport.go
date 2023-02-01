package checker

/*
*
*   Author       ->  AnAverageBeing
*   License      ->  GNU
*   Discord      ->  GigaCat#1521
*   Description  ->  Blazingly fast proxy utility tool
*
 */

import (
	"GigaCat/ProxD/proxy"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func Configure(proxy *proxy.Proxy) (*http.Transport, error) {
	proxyUrl, err := url.Parse(fmt.Sprintf("%s://%s", proxy.Protocol, proxy.IP+":"+strconv.Itoa(int(proxy.Port))))
	if err != nil {
		return nil, err
	}
	transport := &http.Transport{
		Proxy:             http.ProxyURL(proxyUrl),
		ForceAttemptHTTP2: true,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	return transport, nil
}
