package main

import (
	"context"
	"fmt"

	"go.etcd.io/etcd/clientv3"
)

//go run comm.go watch_range.go
//etcdctl watch foo foo4
//etcdctl --endpoints=192.168.195.61:2379,192.168.195.62:2379,192.168.195.63:2379 watch foo --prefix

/*
for i in {1..5};do etcdctl --endpoints=192.168.195.61:2379 put foo$i value$i; done

for i in {1..5};do etcdctl --endpoints=192.168.195.61:2379 del foo$i value$i; done

*/

func main() {
	rwc := cli.Watch(context.Background(), "foo", clientv3.WithRange("foo4"))
	for wresp := range rwc {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q: %q \n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}

}
