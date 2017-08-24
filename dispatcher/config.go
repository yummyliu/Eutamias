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
	nport		int
}

type ConfigError struct {
	info string
}

func (c *ConfigError) Error() string {
	return fmt.Sprintf("%s", c.info)
}

func (con *Config) Read(cfgPath string) error {
	cfg, err := ini.InsensitiveLoad(cfgPath)
	if err != nil {
		return err
	}

	serve_sec, err := cfg.GetSection("servers")
	if err != nil {
		return err
	}

	_nports, err := serve_sec.GetKey("nports")
	if err != nil {
		return err
	}
	nports := strings.Split(_nports.String(), ",")

	_nips, err := serve_sec.GetKey("nips")
	if err != nil {
		return err
	}
	nips := strings.Split(_nips.String(), ",")
	if len(nports) != len(nips) {
		return &ConfigError{
			info: "nports != nips",
		}
	}

	lf, err := serve_sec.GetKey("logfilepath")
	if err != nil {
		return err
	}
	cport, err := serve_sec.GetKey("clientport")
	if err != nil {
		return err
	}
	ccport, err := cport.Int()
	if err != nil {
		return err
	}
	nport, err := serve_sec.GetKey("NserverListenPort")
	if err != nil {
		return err
	}
	nnport, err := nport.Int()
	if err != nil {
		return err
	}

	con.LogFilePath = lf.String()
	con.Nips = nips
	con.Nports = nports
	con.cport = ccport
	con.nport = nnport
	return nil
}
