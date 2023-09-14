package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	// do client
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://localhost:20000", "http://localhost:20002", "http://localhost:20004"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	defer cli.Close()
	fmt.Println("connect to etcd success")
	defer cli.Close()
	go Watch(cli)
	Create(cli)
	Read(cli)
	Delete(cli)
	Update(cli)
	select {}
}

func Watch(cli *clientv3.Client) {
	rch := cli.Watch(context.Background(), "name")
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s, Key: %s, Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
	fmt.Println("out")
}

func Create(cli *clientv3.Client) {
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	_, err := cli.Put(ctx, "name", "test for put")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed: %v\n", err)
		return
	}
}

func Read(cli *clientv3.Client) {
	// get
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	resp, err := cli.Get(ctx, "name")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("Type: %s, Key:%s, Value:%s\n", "READ", ev.Key, ev.Value)
	}
}

func Update(cli *clientv3.Client) {
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	_, err := cli.Put(ctx, "name", "test for update")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
}
func Delete(cli *clientv3.Client) {
	// delete
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	_, err := cli.Delete(ctx, "name")
	cancel()
	if err != nil {
		fmt.Printf("delete from etcd failed, err:%v\n", err)
		return
	}
}
