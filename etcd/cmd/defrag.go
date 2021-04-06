package main

import (
	"context"
	"fmt"
	"log"
)

//go run comm.go defrag.go
//etcdctl endpoint defrag
//etcdctl --endpoints=192.168.195.61:2379,192.168.195.62:2379,192.168.195.63:2379 defrag
func main() {
	defer cli.Close()

	for _, ep := range endpoints {
		_, err := cli.Defragment(context.Background(), ep)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("------------------done")
}
