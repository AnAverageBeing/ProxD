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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type host struct {
	Origin string `json:"origin"`
}

func CheckProxy(checkedProxy *[]proxy.Proxy, proxy *proxy.Proxy, timeout time.Duration, maxRetries int, cfg *utils.Config) error {
	var tried int
	alive := false
	transport, err := Configure(proxy)
	if err != nil {
		return nil
	}
	client := http.Client{
		Transport: transport,
		Timeout:   timeout,
	}

	var req *http.Request
	var resp *http.Response
	var success bool
	var body []byte

	for i := 0; maxRetries > i; i++ {

		if tried > maxRetries || alive {
			break
		}
		req, err = http.NewRequest("GET", "https://httpbin.org/ip", nil)
		if err != nil {
			tried++
			continue
		}

		resp, success, err = doRequest(client, req)
		if err != nil || !success {
			tried++
			continue
		}

		alive = true
		break
	}

	if tried > maxRetries || !alive {
		return nil
	}

	var jsonData host
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return err
	}
	proxyUrl, err := transport.Proxy(req)
	if err != nil {
		return err
	}
	if fmt.Sprintf("%s:%s", jsonData.Origin, proxyUrl.Port()) == proxyUrl.Host {
		logger.LogProxy(proxy)
		utils.Save(proxy, cfg)
		*checkedProxy = append(*checkedProxy, *proxy)
	}
	return nil
}

func doRequest(client http.Client, req *http.Request) (*http.Response, bool, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, false, err
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return resp, true, nil
	}
	resp.Body.Close()
	return nil, false, nil
}
