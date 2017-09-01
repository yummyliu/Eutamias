package main

import (
	"fmt"
	ini "github.com/go-ini/ini"
)

type Config struct {
	logPath		string
	port        uint64
	ip			string
	dIp			string
	dPort		uint64
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
	listen_port, err := serve_sec.GetKey("Listenport")
	if err != nil {
		return err
	}
	con.port, err = listen_port.Uint64()
	if err != nil {
		return err
	}
	listen_ip, err := serve_sec.GetKey("ListenIp")
	if err != nil {
		return err
	}
	con.ip = listen_ip.String()

	dip, err := serve_sec.GetKey("dispatcherIp")
	if err != nil {
		return err
	}
	con.dIp= dip.String()
	if err != nil {
		return err
	}
	dport, err := serve_sec.GetKey("dispatcherPort")
	if err != nil {
		return err
	}
	con.dPort, err = dport.Uint64()
	if err != nil {
		return err
	}

	return nil
}
