package proxy

/*
*
*   Author       ->  AnAverageBeing
*   License      ->  GNU
*   Discord      ->  GigaCat#1521
*   Description  ->  Blazingly fast proxy utility tool
*
 */

type ProxyProtocol int

const (
	HTTP ProxyProtocol = iota
	HTTPS
	SOCKS4
	SOCKS5
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
