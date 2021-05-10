package main

import (
	"flag"
	"fmt"
	"version/core"
	"version/logger"
)

var (
	host string
	port int
)

func init() {
	flag.StringVar(&host, "host", "127.0.0.1", "set the listen host")
	flag.IntVar(&port, "port", 8002, "set the listen port")

	flag.Usage = usage
}

func usage() {
	fmt.Println("Usage version [-host] [-port]")
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	core.SetConfHost(host)
	core.SetConfPort(port)
	logger.Infof("conf is\n%s", core.ConfObj)
	core.Run()
}
