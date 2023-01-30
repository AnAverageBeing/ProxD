package utils

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
	"GigaCat/ProxD/utils/logger"
	"bufio"
	"fmt"
	"os"
)

func GetUrlsList(fileLocation string, protocol proxy.ProxyProtocol) proxy.UrlsList {
	file, err := os.Open(fileLocation)
	if err != nil {
		logger.LogErr(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	urls := []string{}
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		logger.LogErr(err)
	}
	urlsList := proxy.UrlsList{
		URLs:     urls,
		Protocol: protocol,
	}
	return urlsList
}

func RemoveDuplicates(proxies *[]proxy.Proxy) []proxy.Proxy {
	keys := make(map[string]bool)
	list := []proxy.Proxy{}
	for _, proxy := range *proxies {
		key := fmt.Sprintf("%s:%d:%s", proxy.IP, proxy.Port, proxy.Protocol)
		if _, value := keys[key]; !value {
			keys[key] = true
			list = append(list, proxy)
		}
	}
	return list
}
