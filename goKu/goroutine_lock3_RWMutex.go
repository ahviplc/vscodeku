package main

import (
	"fmt"
	"sync"
	"time"
)

// golang 之sync &并发安全锁 - Dwyane.wang - 博客园
// https://www.cnblogs.com/flash55/p/12389832.html

// sync包
// 　　Sync包同步提供基本的同步原语，如互斥锁。 除了Once和WaitGroup类型之外，大多数类型都是供低级库例程使用的。 通过Channel和沟通可以更好地完成更高级别的同步。并且此包中的值在使用过后不要拷贝。
// 　　sync包中主要有：Locker, Cond, Map, Mutex, Once, Pool,、RWMutex, WaitGroup

// var x int64
// var wg sync.WaitGroup
// var lock sync.Mutex // 互斥锁：sync.Mutex

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex   // 互斥锁：sync.Mutex
	rwlock sync.RWMutex // 读写互斥锁：sync.RWMutex
)

func write() {
	// lock.Lock()   // 加互斥锁
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	fmt.Println("write...")
	rwlock.Unlock() // 解写锁
	// lock.Unlock()                     // 解互斥锁
	wg.Done()
}

func read() {
	// lock.Lock()                  // 加互斥锁
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	fmt.Println("read...")
	rwlock.RUnlock() // 解读锁
	// lock.Unlock()                // 解互斥锁
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
