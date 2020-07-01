package main

import (
	"fmt"
	"sync"
)

// 单例（借助Once) 使用sync.Once 包可以实现线程安全的单例模式cccc

type singleton struct{}

var instance *singleton // *是指针运算符 , 可以表示一个变量是指针类型 , 也可以表示一个指针变量所指向的存储单元 , 也就是这个地址所存储的值 .
var once sync.Once

// 获取单例
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {
	single := GetInstance()
	fmt.Printf("single对象的地址: %x\n", &single) // 十六进制表示  & 是取地址符号 , 即取得某个变量的地址 , 如 ; &a
	fmt.Printf("asingle对象指针: %p\n", single)  // 指针
	fmt.Println("-------------------------------------------")
	single2 := GetInstance()
	fmt.Printf("single2对象的地址: %x\n", &single2)
	fmt.Printf("asingle2对象指针: %p\n", single2) // 指针
	fmt.Println("-------------------------------------------")
	single3 := GetInstance()
	fmt.Printf("single3对象的地址: %x\n", &single3)
	fmt.Printf("asingle3对象指针: %p\n", single3) // 指针
	fmt.Println("-------------------------------------------")
}
