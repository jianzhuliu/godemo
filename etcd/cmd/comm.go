package main

import (
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var err error
var cli *clientv3.Client

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 10 * time.Second
	endpoints      = []string{"192.168.195.61:2379", "192.168.195.62:2379", "192.168.195.63:2379"}
)

func init() {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		log.Fatal(err)
	}

}
