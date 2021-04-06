package main

import (
	"context"
	"fmt"
	"log"
)

//go run comm.go put.go
//etcd put foo bar
//etcd --endpoints=192.168.195.61:2379 put foo bar
func main() {

	defer cli.Close()

	_, err = cli.Put(context.TODO(), "foo", "bar")
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	presp, err := cli.Put(ctx, "sample_key", "sample_value")
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(presp)
}
