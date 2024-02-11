package main

import (
	"fmt"
	"time"
)

// 定义事件类型
type Event struct {
	Name string
}

// 定义事件处理函数类型
type EventHandler func(event Event)

// 事件队列
var eventQueue = make(chan Event)

// 注册事件处理函数
func RegisterEventHandler(handler EventHandler) {
	go func() {
		for {
			event := <-eventQueue
			handler(event)
		}
	}()
}

// 发布事件
func PublishEvent(event Event) {
	eventQueue <- event
}

func main() {
	// 注册事件处理函数
	RegisterEventHandler(func(event Event) {
		fmt.Printf("处理事件: %s\n", event.Name)
	})

	// 发布一个事件
	PublishEvent(Event{Name: "startup"})

	// 模拟其他工作...
	time.Sleep(1 * time.Second)

	// 发布另一个事件
	PublishEvent(Event{Name: "shutdown"})

	// 等待上一个事件处理完成
	time.Sleep(2 * time.Second)
}
