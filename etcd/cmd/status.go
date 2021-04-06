package main

import (
	"context"
	"fmt"
	"log"
)

//go run comm.go status.go
//etcdctl endpoint status -w json
//etcdctl --endpoints=192.168.195.61:2379,192.168.195.62:2379,192.168.195.63:2379 endpoint status
func main() {
	defer cli.Close()

	for _, ep := range endpoints {
		resp, err := cli.Status(context.Background(), ep)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("endpoint: %s / Leader: %v\n", ep, resp.Header.MemberId == resp.Leader)
	}
	fmt.Println("------------------done")
}
