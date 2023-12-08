package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	// 创建客户端
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2579"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	defer client.Close()

	// 创建 KV 客户端
	kv := clientv3.NewKV(client)

	// 模拟 etcd 中 KV 的变化
	go func() {
		for {
			kv.Put(context.TODO(), "/onex.io/miners/sdktest", "hello sdk")
			kv.Delete(context.TODO(), "/onex.io/miners/sdktest")
			time.Sleep(2 * time.Second)
		}
	}()

	// 先 GET 到当前的值，并监听后续变化
	resp, err := kv.Get(context.TODO(), "/onex.io/miners/sdktest")
	if err != nil {
		panic(err)
	}

	// 现在 key 是存在的
	if len(resp.Kvs) != 0 {
		fmt.Println("Current value:", string(resp.Kvs[0].Value))
	}

	// 当前 etcd 集群事务ID, 单调递增的
	watchStartRevision := resp.Header.Revision + 1

	// 创建一个 watcher
	watcher := clientv3.NewWatcher(client)

	// 启动监听
	fmt.Println("Start to watch from revision:", watchStartRevision)

	ctx, cancel := context.WithCancel(context.TODO())
	time.AfterFunc(5*time.Second, func() {
		cancel()
	})

	respChan := watcher.Watch(ctx, "/onex.io/miners/sdktest", clientv3.WithRev(watchStartRevision))

	// 处理 kv 变化事件
	for ch := range respChan {
		for _, event := range ch.Events {
			switch event.Type {
			case mvccpb.PUT:
				fmt.Printf("[PUT] %q: %q\n", event.Kv.Key, event.Kv.Value)
			case mvccpb.DELETE:
				fmt.Printf("[DELETE] %q: %q\n", event.Kv.Key, event.Kv.Value)
			default:
				fmt.Printf("[%s] %q: %q\n", event.Type, event.Kv.Key, event.Kv.Value)
			}
		}
	}
}
