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
	"GigaCat/ProxD/utils"
	"GigaCat/ProxD/utils/logger"
	"net/http"
	"time"
)

func CheckProxy(checkedProxy *[]proxy.Proxy, proxy *proxy.Proxy, timeout time.Duration, maxRetries int, cfg *utils.Config) error {

	alive := false
	checkSite := cfg.General.CheckWebsite

	transport, err := Configure(proxy, timeout)

	if err != nil {
		return nil
	}

	client := http.Client{
		Transport: transport,
		Timeout:   timeout,
	}

	for tried := 0; maxRetries > tried; tried++ {

		_, err := client.Get(checkSite)

		if err != nil {
			continue
		}

		alive = true
		break
	}

	if !alive {
		return nil
	}

	logger.LogProxy(proxy)
	utils.Save(proxy, cfg)
	*checkedProxy = append(*checkedProxy, *proxy)

	return nil
}
