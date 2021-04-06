package main

import (
	"context"
	"fmt"
	"log"
)

//go run comm.go get.go
//etcdctl get foo
//etcdctl --endpoints=192.168.195.61:2379,192.168.195.62:2379,192.168.195.63:2379 get foo

func main() {

	defer cli.Close()

	_, err = cli.Put(context.TODO(), "foo", "bar")
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := cli.Get(ctx, "foo")
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
