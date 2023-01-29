package main

import (
	"GigaCat/ProxD/proxy"
	"GigaCat/ProxD/scraper"
	"fmt"
)

/*
*
*   Author       ->  AnAverageBeing
*   License      ->  GNU
*   Discord      ->  GigaCat#1521
*   Description  ->  Blazingly fast proxy utility tool
*
 */

func main() {
	lists := []proxy.ProxyList{
		{
			URLs: []string{
				"https://api.proxyscrape.com/v2/?request=getproxies&protocol=socks4",
				"https://openproxy.space/list/socks4",
				"https://openproxylist.xyz/socks4.txt",
				"https://proxyscan.io/download?type=socks4",
				"https://proxysearcher.sourceforge.net/Proxy%20List.php?type=socks",
				"https://proxyspace.pro/socks4.txt",
				"https://raw.githubusercontent.com/B4RC0DE-TM/proxy-list/main/SOCKS4.txt",
				"https://raw.githubusercontent.com/BlackSnowDot/proxylist-update-every-minute/main/socks.txt",
				"https://raw.githubusercontent.com/ErcinDedeoglu/proxies/main/proxies/socks4.txt",
				"https://raw.githubusercontent.com/hanwayTech/free-proxy-list/main/socks4.txt",
				"https://raw.githubusercontent.com/HyperBeats/proxy-list/main/socks4.txt",
				"https://raw.githubusercontent.com/jetkai/proxy-list/main/online-proxies/txt/proxies-socks4.txt",
				"https://raw.githubusercontent.com/mmpx12/proxy-list/master/socks4.txt",
				"https://raw.githubusercontent.com/ObcbO/getproxy/master/socks4.txt",
				"https://raw.githubusercontent.com/officialputuid/KangProxy/KangProxy/socks4/socks4.txt",
				"https://raw.githubusercontent.com/proxylist-to/proxy-list/main/socks4.txt",
				"https://raw.githubusercontent.com/roosterkid/openproxylist/main/SOCKS4_RAW.txt",
				"https://raw.githubusercontent.com/saschazesiger/Free-Proxies/master/proxies/socks4.txt",
				"https://raw.githubusercontent.com/ShiftyTR/Proxy-List/master/socks4.txt",
				"https://raw.githubusercontent.com/TheSpeedX/PROXY-List/master/socks4.txt",
				"https://raw.githubusercontent.com/Zaeem20/FREE_PROXIES_LIST/master/socks4.txt",
				"https://raw.githubusercontent.com/zevtyardt/proxy-list/main/socks4.txt",
				"https://spys.me/socks.txt",
				"https://www.freeproxychecker.com/result/socks4_proxies.txt",
				"https://www.proxy-list.download/api/v1/get?type=socks4",
			},
			Protocol: proxy.HTTP,
		},
	}
	var proxies []proxy.Proxy
	scraper.Scrape(lists, &proxies)

	// for _, p := range proxies {
	// 	fmt.Println(p)
	// }

	fmt.Println(len(proxies))
}
