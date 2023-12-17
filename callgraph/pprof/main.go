package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	memprofile = flag.String("memprofile", "", "write memory profile to file")
)

func main() {
	flag.Parse()

	runtime.SetCPUProfileRate(500)

	// CPU 分析开始
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	// 程序的主逻辑
	A()

	// 内存分析结束
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // 获取最准确的内存分析数据，需要在写入内存分析文件前调用 GC
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}

}

func A() {
	fmt.Println("Function A called.")
	// 执行 CPU 密集型任务, 选一个足够大的数以确保能够被 pprof 采集到 CPU 执行数据
	fib(40)

	B()
	C()
}

func B() {
	fmt.Println("Function B called.")
	fib(40)
}

func C() {
	fmt.Println("Function C called.")
	fib(40)

	D()
}

func D() {
	fmt.Println("Function D called.")
	fib(40)
}

// fib 计算斐波那契数列的第 n 项（递归方式）。
// 这个任务是 CPU 密集型的，因为递归方法需要大量的函数调用和计算。
// 由于递归方法的效率非常低，尤其是在没有使用缓存或动态规划技术的情况下，
// 这个计算可能会耗时很长，很可能会超过 500 毫秒。如果您的机器计算速度非常快，
// 你可能需要增加 n 的值以确保达到预期的耗时。如果你的计算机计算速度很慢，
// 也可以考虑减少 n 的值，以确保测试程序不会执行太久。
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
