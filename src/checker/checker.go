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
	var tried int
	for tried = 1; tried < maxRetries+1; tried++ {

		resp, err := client.Get(checkSite)
		if err != nil {
			continue
		}

		if resp.StatusCode == 200 {
			defer resp.Body.Close()
			alive = true
			break
		}
	}

	if alive {
		logger.LogProxy(proxy, tried)
		utils.Save(proxy, cfg)
		*checkedProxy = append(*checkedProxy, *proxy)
	}

	return nil
}
