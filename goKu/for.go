package main

import (
	"fmt"
	"strings"
)

func main() {

	var n [10]int /* n 是一个长度为 10 的数组 */
	var i, j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	intData := [...]int{1, 2, 3, 4}

	for key, value := range intData {
		fmt.Printf("key:%d  value:%d\n", key, value)
	}

	strData := []string{"JAVA", "PYTHON", "GO"}

	for i := 0; i < len(strData); i++ {
		fmt.Printf("key:%d  value:%v\n", i, strData[i])
	}

	fmt.Println(strings.Title("i am lc"))
}
