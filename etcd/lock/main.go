package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
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

	// 创建两个单独的会话用来演示锁竞争
	s1, err := concurrency.NewSession(client)
	if err != nil {
		panic(err)
	}
	defer s1.Close()
	m1 := concurrency.NewMutex(s1, "/onex.io/lock")

	s2, err := concurrency.NewSession(client)
	if err != nil {
		panic(err)
	}
	defer s2.Close()
	m2 := concurrency.NewMutex(s2, "/onex.io/lock")

	// 会话 s1 获取锁
	if err := m1.Lock(context.TODO()); err != nil {
		panic(err)
	}
	fmt.Println("acquired lock for s1")

	m2Locked := make(chan struct{})
	go func() {
		defer close(m2Locked)
		// 等待直到会话 s1 释放了 /onex.io/lock 的锁
		if err := m2.Lock(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err := m1.Unlock(context.TODO()); err != nil {
		panic(err)
	}
	fmt.Println("released lock for s1")

	<-m2Locked
	fmt.Println("acquired lock for s2")
}
