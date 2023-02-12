package main

import (
	"GigaCat/ProxD/checker"
	"GigaCat/ProxD/proxy"
	"GigaCat/ProxD/scraper"
	"GigaCat/ProxD/utils"
	"GigaCat/ProxD/utils/logger"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/schollz/progressbar/v3"
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
var proxies = make([]proxy.Proxy, 0, 1000)
var checked = make([]proxy.Proxy, 0, 1000)
var urlsList = make([]proxy.UrlsList, 0, 3)

func main() {

	flag.Parse()
	utils.PrintLogo()
	cfg, err := utils.GetConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}

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

	logger.LogInfo("Scraping Proxies")

	scraper.Scrape(&urlsList, &proxies, time.Duration(float64(cfg.General.Timeout)*float64(time.Second)))

	utils.RemoveDuplicates(&proxies)

	logger.LogInfo(fmt.Sprintf("Scraped %d proxies", len(proxies)))
	logger.LogInfo("Checking Proxies")

	group := new(errgroup.Group)
	group.SetLimit(cfg.General.MaxConnections)

	bar := progressbar.Default(int64(len(proxies)))

	for _, p := range proxies {
		p := p
		group.Go(func() error {
			bar.Add(1)
			return checker.CheckProxy(&checked, &p, time.Duration(float64(cfg.General.Timeout)*float64(time.Second)), cfg.General.MaxRetries, &cfg)
		})
	}

	group.Wait()
}
