package main

import (
	"context"
	"fmt"
)

//go run comm.go watch.go
//etcdctl watch foo
//etcdctl --endpoints=192.168.195.61:2379,192.168.195.62:2379,192.168.195.63:2379 watch foo

func main() {
	rwc := cli.Watch(context.Background(), "foo")
	for wresp := range rwc {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q: %q \n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}

}
