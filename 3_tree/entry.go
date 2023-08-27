/**
*@Date: 星期日 2023/8/27 22:13
*@File: entiy
*@Author: owl
*@Description: 包和封装（通过拆树的结构体和方法）
**/

package main

import (
	"3_tree/tree"
	"fmt"
)

func main() {
	//创建 树
	var root tree.Node
	fmt.Println(root) // {0 <nil> <nil>} Value Left Right

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node) //指针和实例都可以用'.'
	root.Left.Right = tree.CreateNode(7)
	fmt.Println(root) //{3 {0 nil nil} {5 nil nil} {0 nil nil}} x   {3 0xc000008090 0xc0000080a8}√

	fmt.Println("中序遍历")
	root.Traverse() //0 7 3 0 5
	fmt.Println()

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes) //[{3 <nil> <nil>} {0 <nil> <nil>} {6 <nil> 0xc000110060}]

	//结构体方法调用
	root.Right.Left.Print()
	root.Right.Left.SetValue(6)
	root.Right.Left.Print()
	fmt.Println()

	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(100)
	pRoot.Print()
	fmt.Println()

	root.Left.Left.SetValue(1)
	root.Left.Left = &tree.Node{Value: 8}
	root.Left.Left.SetValue(800)
	root.Left.Left.Print()
	fmt.Println()

}
