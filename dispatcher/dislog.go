package main

import (
	"github.com/op/go-logging"
	"os"
)
func init_log(logFilePath string) error {
	log = logging.MustGetLogger("main")
	var format = logging.MustStringFormatter(
		`%{id:08x}--%{time}--%{level:.10s}--%{shortfile} %{message}`,
	)

	if logFilePath != "" {
		logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return err
		}
		var backend = logging.NewLogBackend(logFile, "D:", 0)
		var backendFormatter = logging.NewBackendFormatter(backend, format)
		var backendLeveled = logging.AddModuleLevel(backendFormatter)
		backendLeveled.SetLevel(logging.DEBUG, "main")
		logging.SetBackend(backendFormatter)
	} else {
		var backend = logging.NewLogBackend(os.Stderr, "D:", 0)
		var backendFormatter = logging.NewBackendFormatter(backend, format)
		var backendLeveled = logging.AddModuleLevel(backendFormatter)
		backendLeveled.SetLevel(logging.DEBUG, "main")
		logging.SetBackend(backendFormatter)
	}
	return nil
}
