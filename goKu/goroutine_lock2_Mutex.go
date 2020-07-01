package main

import (
	"fmt"
	"sync"
)

// golang 之sync &并发安全锁 - Dwyane.wang - 博客园
// https://www.cnblogs.com/flash55/p/12389832.html

// sync包
// 　　Sync包同步提供基本的同步原语，如互斥锁。 除了Once和WaitGroup类型之外，大多数类型都是供低级库例程使用的。 通过Channel和沟通可以更好地完成更高级别的同步。并且此包中的值在使用过后不要拷贝。
// 　　sync包中主要有：Locker, Cond, Map, Mutex, Once, Pool,、RWMutex, WaitGroup

var x int64
var wg sync.WaitGroup
var lock sync.Mutex // 互斥锁：sync.Mutex

// 未加锁
func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg.Done()
}

// 加锁
func addWithLock() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 解锁
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	// // 未加锁调用 每一次都会小于10000 例如 7459 7292 8042 and so on.
	// go add()
	// go add()

	// 加锁调用 每次都会输出为10000
	go addWithLock()
	go addWithLock()

	wg.Wait()
	fmt.Println(x)
}
