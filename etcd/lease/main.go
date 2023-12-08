package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2579"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	defer client.Close()

	lease := clientv3.NewLease(client)
	kv := clientv3.NewKV(client)

	// 创建一个 TTL 为 10s 的租约
	resp, err := lease.Grant(context.TODO(), 10)
	if err != nil {
		panic(err)
	}

	// 存储一个 10s 过期的 key
	_, err = kv.Put(context.TODO(), "/onex.io/miners/vanish", "vanish in 10s", clientv3.WithLease(resp.ID))
	if err != nil {
		panic(err)
	}

	// 对租约进行续期
	_, err = lease.KeepAliveOnce(context.TODO(), resp.ID)
	if err != nil {
		panic(err)
	}

	// 永久续期
	ch, err := client.KeepAlive(context.TODO(), resp.ID)
	if err != nil {
		panic(err)
	}

	for {
		resp := <-ch
		fmt.Println("ttl:", resp.TTL)
	}
}
