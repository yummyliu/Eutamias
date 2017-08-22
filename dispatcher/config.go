package main

import (
	"log"
	ini "github.com/go-ini/ini"
)

type Config struct {
	Nips   string
	Nports string
}

func read_cfg(cfg_path string) *Config {
	//ini_:= make(map[string]string)
	cfg,err := ini.InsensitiveLoad(cfg_path)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	serve_sec,err := cfg.GetSection("servers")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	nports,err := serve_sec.GetKey("nports")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	nips,err := serve_sec.GetKey("nips")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &Config{
		Nips : nips.String(),
		Nports : nports.String(),
	}
}
