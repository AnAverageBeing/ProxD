package scraper

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
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly"
)

var re = regexp.MustCompile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}:\d{1,5}\b`)
var wg sync.WaitGroup
var c = colly.NewCollector()

func Scrape(list *[]proxy.UrlsList, proxies *[]proxy.Proxy, timeout time.Duration) {
	c.SetRequestTimeout(timeout)
	for _, proxyList := range *list {
		if len(proxyList.URLs) > 0 {
			wg.Add(1)
			go scrapeURL(proxyList.URLs, proxyList.Protocol, proxies)
		}
	}
	wg.Wait()
}

func scrapeURL(urls []string, protocol proxy.ProxyProtocol, proxies *[]proxy.Proxy) {
	defer wg.Done()
	c.OnResponse(func(r *colly.Response) {
		proxiesRaw := re.FindAllString(string(r.Body), -1)

		for _, raw := range proxiesRaw {
			ipAndPort := strings.Split(raw, ":")
			ip := ipAndPort[0]
			port, err := strconv.Atoi(ipAndPort[1])

			if err != nil {
				continue
			}

			*proxies = append(*proxies, proxy.Proxy{
				IP:       ip,
				Port:     uint16(port),
				Protocol: protocol,
			})
		}
	})

	var urlsWg sync.WaitGroup
	for _, url := range urls {
		urlsWg.Add(1)
		go c.Visit(url)
	}
	urlsWg.Wait()
}
