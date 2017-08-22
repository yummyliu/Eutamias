package main

import (
	"flag"
	"fmt"
	"math/rand"
	log "github.com/golang/glog"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Info("")
		fmt.Println("usage: dispatcher config")
		return
	}

	config := read_cfg(flag.Args()[0])
	log.Infof("%s:%s",config.Nips,config.Nports)
}
