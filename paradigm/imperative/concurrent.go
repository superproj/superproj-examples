package main

import (
	"fmt"
	"sync"
)

func worker(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // 在协程结束时减少等待组中的计数
	for job := range jobs {
		// 模拟执行任务
		result := job * 2
		// 将结果发送到结果通道
		results <- result
	}
}

func main() {
	// 创建任务通道
	jobs := make(chan int, 10)
	// 创建结果通道
	results := make(chan int, 10)
	// 创建等待组
	var wg sync.WaitGroup

	// 启动多个工作协程
	for i := 1; i <= 5; i++ {
		wg.Add(1) // 增加等待组中的计数
		go worker(jobs, results, &wg)
	}

	// 发送任务到任务通道
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	// 关闭任务通道
	close(jobs)

	// 等待所有工作协程完成
	go func() {
		wg.Wait()
		close(results) // 关闭结果通道
	}()

	// 收集并打印结果
	for result := range results {
		fmt.Println("结果:", result)
	}
}
