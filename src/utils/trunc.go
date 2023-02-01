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
	"os"
)

func Trunc(cfg *Config) error {
	file, err := os.OpenFile(cfg.HTTP.SaveFile, os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	file, err = os.OpenFile(cfg.SOCKS4.SaveFile, os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	file, err = os.OpenFile(cfg.SOCKS5.SaveFile, os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
