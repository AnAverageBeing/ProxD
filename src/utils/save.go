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
	"fmt"
	"log"
	"os"
	"strconv"
)

func Save(p *proxy.Proxy, cfg *Config) error {

	var file *os.File
	var err error

	if p.Protocol == proxy.HTTP {
		file, err = os.OpenFile(cfg.HTTP.SaveFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	}
	if p.Protocol == proxy.SOCKS4 {
		file, err = os.OpenFile(cfg.SOCKS4.SaveFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	}
	if p.Protocol == proxy.SOCKS5 {
		file, err = os.OpenFile(cfg.SOCKS5.SaveFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)

	}

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s\r\n", p.IP+":"+strconv.Itoa(int(p.Port))))

	if err != nil {
		log.Fatal(err)
	}
	return nil
}
