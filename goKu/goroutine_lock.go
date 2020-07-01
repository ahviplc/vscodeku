package main

// go 携程共享 数据
// 加锁解锁操作
// 同步锁
import (
	"fmt"
	"sync"
)

func computed(data *int, lock *sync.Mutex) {
	lock.Lock() // 加锁霸占数据
	*data++
	lock.Unlock() // 解锁让出数据
}

func main() {
	var data int = 0
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ { // 启动10个携程
		go computed(&data, lock)
	}

	for {
		lock.Lock()
		c := data
		lock.Unlock()
		if c >= 5 {
			fmt.Println(data)
			break
		}
	}
}
