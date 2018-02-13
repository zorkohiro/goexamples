package main

import (
	"log"
	"log/syslog"
)

func main() {
	loghdl, err := syslog.New(syslog.LOG_DAEMON|syslog.LOG_INFO, "logger")
	if err != nil {
		log.Fatal("could not setup up syslog")
	}
	loghdl.Info("got me")
}
