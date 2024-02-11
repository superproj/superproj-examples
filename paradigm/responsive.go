package main

import (
	"fmt"
	"time"
)

// 定义事件发射器结构体
type EventEmitter struct {
	channels map[string][]func(interface{})
}

// 创建新的事件发射器
func NewEmitter() *EventEmitter {
	return &EventEmitter{
		channels: make(map[string][]func(interface{})),
	}
}

// 注册事件处理函数
func (e *EventEmitter) On(eventName string, handler func(interface{})) {
	if _, exists := e.channels[eventName]; !exists {
		e.channels[eventName] = make([]func(interface{}), 0)
	}
	e.channels[eventName] = append(e.channels[eventName], handler)
}

// 触发事件
func (e *EventEmitter) Emit(eventName string, data interface{}) {
	if handlers, exists := e.channels[eventName]; exists {
		for _, handler := range handlers {
			go handler(data)
		}
	}
}

func main() {
	emitter := NewEmitter()

	// 注册事件处理函数
	emitter.On("click", func(data interface{}) {
		fmt.Println("Clicked:", data)
	})

	// 触发事件
	emitter.Emit("click", "Button 1")
	emitter.Emit("click", "Button 2")
	emitter.Emit("click", "Button 3")

	// 使用无限循环来创建一个阻塞状态
	for {
		time.Sleep(1 * time.Second)
	}
}
