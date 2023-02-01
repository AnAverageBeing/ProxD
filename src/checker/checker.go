package checker

import (
	"GigaCat/ProxD/proxy"
	"GigaCat/ProxD/utils"
	"encoding/json"
	"fmt"
	"io"
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
	var reader io.ReadCloser
	var Req *http.Request
	var Resp *http.Response

	for i := 0; maxRetries > i; i++ {

		if tried > maxRetries || alive {
			break
		}
		req, err := http.NewRequest("GET", "https://httpbin.org/ip", reader)
		if err != nil {
			tried++
			continue
		}

		resp, err := client.Do(req)

		if err != nil {
			tried++
			continue
		}

		defer resp.Body.Close()

		alive = true
		Req = req
		Resp = resp

		break
	}

	if tried > maxRetries || !alive {
		return nil
	}

	var jsonData host
	err = json.NewDecoder(Resp.Body).Decode(&jsonData)
	if err != nil {
		return nil
	}
	proxyUrl, err := transport.Proxy(Req)
	if err != nil {
		return nil
	}
	if fmt.Sprintf("%s:%s", jsonData.Origin, proxyUrl.Port()) == proxyUrl.Host {
		utils.Save(proxy, cfg)
		*checkedProxy = append(*checkedProxy, *proxy)
	}
	return nil
}
