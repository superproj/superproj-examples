package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	// 创建 Etcd 客户端
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2579"},
		// Endpoints: []string{"localhost:2379", "localhost:22379", "localhost:32379"}
		// 指定创建客户端时的连接超时时间，这里指定了 5s。
		// 一旦客户端创建成功，我们就不用再关心后续底层连接的状态了，客户端内部会自动重连。
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	// 关闭客户端，释放资源
	defer client.Close()

	// 创建 KV 客户端
	kv := clientv3.NewKV(client)

	// 设置 1 秒超时，访问 etcd 有超时控制
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	// 设置 key
	_, err = kv.Put(ctx, "/onex.io/miners/sdktest1", "Hello Etcd!")
	//操作完毕，取消 etcd
	cancel()
	if err != nil {
		panic(err)
	}

	// 这里再新增 2 个 key，用来做测试
	_, err = kv.Put(context.Background(), "/onex.io/miners/sdktest2", "Hello World!")
	if err != nil {
		panic(err)
	}
	_, err = kv.Put(context.Background(), "/onex.io/minersspam", "spam")
	if err != nil {
		panic(err)
	}

	// 取值，设置超时为 1 秒
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := kv.Get(ctx, "/onex.io/miners/sdktest1")
	cancel()
	if err != nil {
		panic(err)
	}

	// 打印 Get 返回结果
	for _, ev := range resp.Kvs {
		fmt.Printf("%s: %s\n", ev.Key, ev.Value)
	}

	// 获取所有以 `/onex.io/miners/` 为前缀的 key
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err = kv.Get(context.TODO(), "/onex.io/miners/", clientv3.WithPrefix())
	cancel()
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Kvs)

	// 删除 key，设置超时为 1 秒
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	_, err = kv.Delete(ctx, "/onex.io/miners")
	cancel()
	if err != nil {
		panic(err)
	}

	// Op 操作
	ops := []clientv3.Op{
		clientv3.OpPut("/onex.io/miners/puttest", "value1"),
		clientv3.OpGet("/onex.io/miners/puttest"),
		clientv3.OpPut("/onex.io/miners/puttest", "value2"),
	}
	for _, op := range ops {
		_, err := kv.Do(context.TODO(), op)
		if err != nil {
			panic(err)
		}
	}

	// Txn 操作
	_, err = kv.Txn(context.TODO()).If(
		clientv3.Compare(clientv3.Value("k1"), ">", "v1"),
		clientv3.Compare(clientv3.Version("k1"), "=", 2),
	).Then(
		clientv3.OpPut("k2", "v2"), clientv3.OpPut("k3", "v3"),
	).Else(
		clientv3.OpPut("k4", "v4"), clientv3.OpPut("k5", "v5"),
	).Commit()
	if err != nil {
		panic(err)
	}
}
