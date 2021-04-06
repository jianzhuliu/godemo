package main

import (
	"context"
	"fmt"
	"log"

	"go.etcd.io/etcd/clientv3"
)

//go run comm.go txn.go

func main() {

	defer cli.Close()

	k, v := "foo", "b"
	_, err := cli.Put(context.Background(), k, v)
	if err != nil {
		log.Fatal(err)
	}

	kvc := clientv3.NewKV(cli)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)

	_, err = kvc.Txn(ctx).
		If(clientv3.Compare(clientv3.Value(k), ">", "a")).
		Then(clientv3.OpPut(k, "A")).
		Else(clientv3.OpPut(k, "C")).
		Commit()
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}

	gresp, err := kvc.Get(context.Background(), k)
	if err != nil {
		log.Fatal(err)
	}

	for _, ev := range gresp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
