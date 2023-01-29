package proxy

/*
*
*   Author       ->  AnAverageBeing
*   License      ->  GNU
*   Discord      ->  GigaCat#1521
*   Description  ->  Blazingly fast proxy utility tool
*
 */

type ProxyProtocol string

const (
	HTTP   ProxyProtocol = "HTTP"
	SOCKS4 ProxyProtocol = "SOCKS4"
	SOCKS5 ProxyProtocol = "SOCKS5"
)

type Proxy struct {
	IP       string
	Port     uint16
	Protocol ProxyProtocol
}

type ProxyList struct {
	URLs     []string
	Protocol ProxyProtocol
}
