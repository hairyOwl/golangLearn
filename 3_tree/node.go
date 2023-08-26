/**
*@Date: 星期六 2023/8/26 22:27
*@File: node
*@Author: owl
*@Description: 结构体
**/

package main

import (
	"fmt"
)

// 结构体声明
type treeNode struct {
	value       int
	left, Right *treeNode
}

// 结构体方法，方法名前是接收者(node treeNode)相当于this
func (node *treeNode) print() { //(node treeNode)值传递
	fmt.Print(node.value, " ")
}
func (node *treeNode) setValue(v int) { //值传递不能修改值 所以为了修改value,接收者用地址
	if node == nil {
		fmt.Println("nil指针设置value 忽略")
		return
	}
	node.value = v
}

// 遍历树 中序遍历 左中右
func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.Right.traverse()

}

// 自定义工厂函数限定指定构造格式
func createNode(v int) *treeNode {
	return &treeNode{value: v} //返回局部变量的地址
}

func main() {
	//创建 树
	var root treeNode
	fmt.Println(root) // {0 <nil> <nil>} value left Right

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.Right = &treeNode{5, nil, nil}
	root.Right.left = new(treeNode) //指针和实例都可以用'.'
	root.left.Right = createNode(7)
	fmt.Println(root) //{3 {0 nil nil} {5 nil nil} {0 nil nil}} x   {3 0xc000008090 0xc0000080a8}√

	fmt.Println("中序遍历")
	root.traverse() //0 7 3 0 5
	fmt.Println()

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes) //[{3 <nil> <nil>} {0 <nil> <nil>} {6 <nil> 0xc000110060}]

	//结构体方法调用
	root.Right.left.print()
	root.Right.left.setValue(6)
	root.Right.left.print()
	fmt.Println()

	pRoot := &root
	pRoot.print()
	pRoot.setValue(100)
	pRoot.print()
	fmt.Println()

	root.left.left.setValue(1)
	root.left.left = &treeNode{value: 8}
	root.left.left.setValue(800)
	root.left.left.print()
	fmt.Println()

}
