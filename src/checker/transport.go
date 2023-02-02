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
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"h12.io/socks"
)

func Configure(p *proxy.Proxy, timeout time.Duration) (*http.Transport, error) {
	proxyUrl, err := url.Parse(fmt.Sprintf("%s://%s", p.Protocol, p.IP+":"+strconv.Itoa(int(p.Port))))
	if err != nil {
		return nil, err
	}

	if p.Protocol == proxy.SOCKS4 {
		dialSocksProxy := socks.Dial(proxyUrl.String() + "?timeout=" + fmt.Sprint(timeout.String()))
		tr := &http.Transport{Dial: dialSocksProxy}
		return tr, nil
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}

	return transport, nil
}
