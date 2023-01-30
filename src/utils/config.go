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
	"gopkg.in/gcfg.v1"
)

type Config struct {
	General struct {
		Timeout       int
		SourceTimeout int
		MaxThreads    int
		CheckWebsite  string
	}
	HTTP struct {
		Enabled     bool
		SourcesFile string
		UrlListFile string
		SaveFile    string
	}
	SOCKS4 struct {
		Enabled     bool
		SourcesFile string
		UrlListFile string
		SaveFile    string
	}
	SOCKS5 struct {
		Enabled     bool
		SourcesFile string
		UrlListFile string
		SaveFile    string
	}
}

func GetConfig(location *string) (Config, error) {
	var cfg Config
	err := gcfg.ReadFileInto(&cfg, *location)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
