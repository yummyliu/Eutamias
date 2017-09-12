package main

import (
	"fmt"
	ini "github.com/go-ini/ini"
)

type Config struct {
	logPath		string
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
	lf, err := serve_sec.GetKey("logfilepath")
	if err != nil {
		return err
	}
	con.logPath = lf.String()
	cport, err := serve_sec.GetKey("ClientListenport")
	if err != nil {
		return err
	}
	con.cport, err = cport.Int()
	if err != nil {
		return err
	}
	nport, err := serve_sec.GetKey("NserverListenPort")
	if err != nil {
		return err
	}
	con.nport, err = nport.Int()
	if err != nil {
		return err
	}

	return nil
}
