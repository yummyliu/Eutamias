package main

import (
	"fmt"
	ini "github.com/go-ini/ini"
	"strings"
)

type Config struct {
	LogFilePath string
	Nips        []string
	Nports      []string
	cport       int
}

type ConfigError struct {
	info string
}

func (c *ConfigError) Error() string{
	return fmt.Sprintf("%s", c.info)
}

func (con *Config) Read(cfgPath string) error {
	cfg, err := ini.InsensitiveLoad(cfgPath)
	if err != nil {
		return nil
	}

	serve_sec, err := cfg.GetSection("servers")
	if err != nil {
		return nil
	}

	_nports, err := serve_sec.GetKey("nports")
	if err != nil {
		return nil
	}
	nports := strings.Split(_nports.String(), ",")

	_nips, err := serve_sec.GetKey("nips")
	if err != nil {
		return nil
	}
	nips := strings.Split(_nips.String(), ",")
	if len(nports) != len(nips) {
		return &ConfigError{
			info : "nports != nips",
		}
	}

	lf, err := serve_sec.GetKey("logfilepath")
	if err != nil {
		return nil
	}
	cport, err := serve_sec.GetKey("clientport")
	if err != nil {
		return nil
	}
	fmt.Println("%s",lf.String())
	con.LogFilePath = lf.String()
	con.Nips = nips
	con.Nports = nports
	con.cport, err = cport.Int()
	return err
}
