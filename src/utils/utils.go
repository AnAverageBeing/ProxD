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
	"strconv"
	"strings"
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

func GetFromFile(protocol proxy.ProxyProtocol, fileLocation string) (proxies []proxy.Proxy) {
	file, err := os.Open(fileLocation)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// split the line by ":"
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}
		ip := parts[0]
		port, err := strconv.ParseUint(parts[1], 10, 16)
		if err != nil {
			continue
		}

		proxies = append(proxies, proxy.Proxy{
			IP:       ip,
			Port:     uint16(port),
			Protocol: protocol,
		})
	}
	return
}

func RemoveDuplicates(proxies *[]proxy.Proxy) {
	keys := make(map[string]bool)
	list := []proxy.Proxy{}
	for _, proxy := range *proxies {
		key := fmt.Sprintf("%s:%d:%s", proxy.IP, proxy.Port, proxy.Protocol)
		if _, value := keys[key]; !value {
			keys[key] = true
			list = append(list, proxy)
		}
	}
	*proxies = list
}
