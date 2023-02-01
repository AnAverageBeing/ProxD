package checker

import (
	"GigaCat/ProxD/proxy"
	"GigaCat/ProxD/utils/logger"
	"io"
	"net/http"
	"time"
)

func CheckProxy(checkedProxy *[]proxy.Proxy, proxy *proxy.Proxy, timeout time.Duration, maxRetries int) error {
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

	for i := 0; maxRetries > i; i++ {

		if tried > maxRetries || alive {
			break
		}
		req, err := http.NewRequest("GET", "https://httpbin.org/ip", reader)
		if err != nil {
			tried++
			continue
		}

		res, err := client.Do(req)

		if err != nil {
			tried++
			continue
		}

		defer res.Body.Close()
		alive = true
		break
	}

	if alive {
		*checkedProxy = append(*checkedProxy, *proxy)
		logger.LogProxy(proxy)
		return nil
	}

	return nil
}
