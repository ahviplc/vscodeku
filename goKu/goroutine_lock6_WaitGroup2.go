package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
)

// 获取ip
func GetIp() {
	ip, err := net.ResolveIPAddr("ip", "www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ip.IP) // 180.101.49.11
}

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.apache.org/",
		"http://www.baidu.com/",
	}

	GetIp()

	for _, url := range urls {
		// Increment the WaitGroup counter.
		// 递增 WaitGroup 计数器。
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		// 启动一个Go程来取回URL。
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			// 取回URL
			response, err := http.Get(url)
			fmt.Println("err flag:", err)
			if err != nil { // 如果不等于nil 则打印报错信息
				fmt.Println("error:{}", err)
			}
			fmt.Println(response.StatusCode) // 200
			fmt.Println(response.Status)     // 200 OK
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	// 等待所有的HTTP取回操作完成。
	wg.Wait()
}
