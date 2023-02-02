package main

import (
	"GigaCat/ProxD/checker"
	"GigaCat/ProxD/proxy"
	"GigaCat/ProxD/scraper"
	"GigaCat/ProxD/utils"
	"GigaCat/ProxD/utils/logger"
	"flag"
	"log"
	"runtime"
	"time"

	"golang.org/x/sync/errgroup"
)

/*
*
*   Author       ->  AnAverageBeing
*   License      ->  GNU
*   Discord      ->  GigaCat#1521
*   Description  ->  Blazingly fast proxy utility tool
*
 */

var (
	configFile = flag.String("config", "config.ini", "Location of the configuration file. Default is config.ini")
)
var proxies = make([]proxy.Proxy, 1000)

func main() {

	flag.Parse()
	utils.PrintLogo()
	cfg, err := utils.GetConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}

	urlsList := make([]proxy.UrlsList, 1)

	if cfg.HTTP.Enabled {
		proxies = append(proxies, utils.GetFromFile(proxy.HTTP, cfg.HTTP.SourcesFile)...)
		urlsList = append(urlsList, utils.GetUrlsList(cfg.HTTP.UrlListFile, proxy.HTTP))
	}
	if cfg.SOCKS4.Enabled {
		proxies = append(proxies, utils.GetFromFile(proxy.SOCKS4, cfg.SOCKS4.SourcesFile)...)
		urlsList = append(urlsList, utils.GetUrlsList(cfg.SOCKS4.UrlListFile, proxy.SOCKS4))

	}
	if cfg.SOCKS5.Enabled {
		proxies = append(proxies, utils.GetFromFile(proxy.SOCKS5, cfg.SOCKS5.SourcesFile)...)
		urlsList = append(urlsList, utils.GetUrlsList(cfg.SOCKS5.UrlListFile, proxy.SOCKS5))
	}

	err = utils.Trunc(&cfg)
	if err != nil {
		logger.LogErr(err)
	}

	runtime.GOMAXPROCS(cfg.General.MaxThreads)
	logger.LogInfo("Scraping Proxies")

	scraper.Scrape(&urlsList, &proxies, time.Duration(float64(cfg.General.Timeout)*float64(time.Second)))

	proxies = utils.RemoveDuplicates(&proxies)

	logger.LogInfo("Checking Proxies")

	checked := make([]proxy.Proxy, 1000)

	group := new(errgroup.Group)
	group.SetLimit(-1)

	for _, p := range proxies {
		p := p
		group.Go(func() error {
			return checker.CheckProxy(&checked, &p, time.Duration(float64(cfg.General.Timeout)*float64(time.Second)), cfg.General.MaxRetries, &cfg)
		})
	}

	group.Wait()
}
