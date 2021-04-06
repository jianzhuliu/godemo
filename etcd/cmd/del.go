package main

import (
	"context"
	"fmt"
	"log"

	"go.etcd.io/etcd/clientv3"
)

//go run comm.go del.go
//etcdctl del foo --prefix
//for i in {1..5};do etcdctl --endpoints=192.168.195.61:2379 put foo$i value$i; done

func main() {

	defer cli.Close()

	ctx := context.Background()
	gresp, err := cli.Get(ctx, "foo", clientv3.WithPrefix())
	if err != nil {
		log.Fatal(err)
	}

	dresp, err := cli.Delete(context.TODO(), "foo", clientv3.WithPrefix())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("has deleted all keys, %t \n", int64(len(gresp.Kvs)) == dresp.Deleted)
}
