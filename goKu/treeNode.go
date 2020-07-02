package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

// 使用 *treeNode 指针传递 可改变值 只有使用指针才可以改变结构内容 nil指针也可以调用方法
// treeNode 不改变值
// 设置值
func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node Ignored.")
		return
	}
	node.value = value
}

// traverse 遍历树节点
func (node *treeNode) traverse() {
	if node == nil {
		return
	}

	// go里面 nil也可以调用 只是一个普通的函数 不需要判断
	//if node.left != nil {
	//	node.left.traverse()
	//}

	node.left.traverse()
	node.print()
	node.right.traverse()
}

// 创建树节点
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

// 打印
func (node treeNode) print() {
	fmt.Print(node.value, " ")
}

//  print第二种写法 打印
func print2(node treeNode) {
	fmt.Print(node.value, " ")
}

func main() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)
	root.right.left.setValue(4)

	root.traverse() // 0 2 3 4 5

	fmt.Println()
	root.print()
	root.right.left.setValue(4)
	fmt.Println()
	root.right.left.print()
	fmt.Println()

	print2(root)

	root.setValue(100)

	pRoot := root
	pRoot.print()
	pRoot.setValue(200)
	pRoot.print()

	var pRoot2 *treeNode
	pRoot2.setValue(200) // Setting value to nil node Ignored.
	pRoot2 = &root
	pRoot2.setValue(300)
	pRoot2.print()

	// nodes := []treeNode{
	//	{value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//
	// fmt.Println(nodes)

	// 关于 值接收者 和 指针接收者
	// 1. 要改变内容必须使用指针接收者
	// 2. 结构过大也考虑使用指针接收者
	// 3. 一致性:如有指针接收者,最好都是指针接收者

	// 值接收则 是go语言特有
	// 值/指针接收者均可接收值/指针

	// 无论地址还是结构本身,一律使用 . 来访问成员

	// go语言仅支持封装 不支持继承和多态
	// go语言没有class 只有struct
}
