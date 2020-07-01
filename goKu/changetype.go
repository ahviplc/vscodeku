package main

import "fmt"

func main() {
	// 将整型转化为浮点型，并计算结果，将结果赋值给浮点型变量
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)
}
