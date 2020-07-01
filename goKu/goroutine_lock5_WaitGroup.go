package main

// waitgroup 用来等待一组goroutines的结束，在主Goroutine里声明，并且设置要等待的goroutine的个数，
// 每个goroutine执行完成之后调用 Done，最后在主Goroutines 里Wait即可。waitgroup含有三种方法
// func (wg *WaitGroup) Add(d int)  //计数器+d
// func (wg *WaitGroup) Done()  //计数器-1
// func (wg *WaitGroup) Wait()  //阻塞直到计数器变为0

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello() {
	defer wg.Done()
	fmt.Println("Hello Goroutine By LC!")
}
func main() {
	wg.Add(1)
	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	wg.Wait()
}

// 剩下:

// 竞态检测器(race detector)
// 　　在编写时即使最大的仔细还会出现并发上的错误，幸运的是，go语言在运行时和工具链装备了一个精致并易于使用的动态分析工具：竞态检测器。

// 简单的把 -race命令行参数加到go build, go run, go test命令里边即可使用该功能。它会让编译器为你的应用或测试构建一个修订后的版本。

// 　　竞态检测器会研究事件流，找到那些有问题的案例，即一个goroutine写入一个变量后，中间没有任何同步的操作，就有另一个goroutine写入了该变量。这种案例表明有对共享变量的并发访问，即数据动态。

// 　　竞态检测器报告所有实际运行了的数据竞态。它只能检测到那些在运行时发生的竞态，无法用来保证肯定不会发生京态。

// 　　有兴趣的可以仔细研究。

// 这些demo来自:
// golang 之sync &并发安全锁 - Dwyane.wang - 博客园
// https://www.cnblogs.com/flash55/p/12389832.html
