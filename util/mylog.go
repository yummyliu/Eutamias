package util

import "os"
import "github.com/op/go-logging"

func GetLogger(logFilePath, modeleName string) *Logger{
	var log = logging.MustGetLogger("util")
	var format = logging.MustStringFormatter(
		`%{id:08x}--%{time}--%{level:.10s}--%{shortfile} %{message}`,
	)

	logFile, err := os.OpenFile(logFilePath,os.O_WRONLY,0666)
	if err != nil {
		fmt.Println(err)
	}
	var backend = logging.NewLogBackend(logFile, "", 0)
//	var backend = logging.NewLogBackend(os.Stderr, "", 0)
	var backendFormatter = logging.NewBackendFormatter(backend, format)
	var backendLeveled = logging.AddModuleLevel(backendFormatter)
	backendLeveled.SetLevel(logging.CRITICAL, "util")

	logging.SetBackend(backendFormatter)

	log.Debugf("debug %s", "secret")
	log.Infof("%s", "info")
	log.Noticef("notice")
	log.Warningf("warning")
	log.Errorf("err")
	log.Criticalf("crit")
}
