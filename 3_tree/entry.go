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

/*
	扩展已有类型
*/
// myTreeNode 扩展树遍历
type myTreeNode struct { //定义别名
	node *tree.Node
}

// 后序遍历 左右中
func (n *myTreeNode) postOrder() {
	if n == nil || n.node == nil {
		return
	}
	left := myTreeNode{n.node.Left}
	right := myTreeNode{n.node.Right}

	left.postOrder()
	right.postOrder()
	n.node.Print()

}

// myTreeNode 扩展树遍历
type myTreeNode1 struct { //使用内嵌
	*tree.Node //内嵌 省略node
}

// 后序遍历 左右中
func (n *myTreeNode1) postOrder1() {
	if n == nil || n.Node == nil {
		return
	}
	left := myTreeNode1{n.Left} //n.Node.Left 内嵌后可写为 n.Left
	right := myTreeNode1{n.Right}

	left.postOrder1()
	right.postOrder1()
	n.Print()

}
func (myNode *myTreeNode1) Traverse() {
	fmt.Println("Shadowed Methods")
}

func main() {
	//创建 树
	//fmt.Println(root) // {0 <nil> <nil>} Value Left Right

	//root := tree.Node{Value: 3}
	root := myTreeNode1{&tree.Node{Value: 3}} //内嵌声明后
	// 下方代码不用动 ，因为node 相当于平铺到myTreeNode1
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node) //指针和实例都可以用'.'
	root.Left.Right = tree.CreateNode(7)
	//fmt.Println(root) //{3 {0 nil nil} {5 nil nil} {0 nil nil}} x   {3 0xc000008090 0xc0000080a8}√

	fmt.Println("中序遍历")
	root.Node.Traverse() //0 7 3 0 5
	fmt.Println()
	root.Traverse() //Shadowed
	fmt.Println()

	//扩展已有类型和方法
	fmt.Println("后序遍历")
	//myRoot := myTreeNode{&root}
	//myRoot.postOrder()
	//fmt.Println("") // 7 0 0 5 3
	root.postOrder1()
	fmt.Println()

	/*nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes) //[{3 <nil> <nil>} {0 <nil> <nil>} {6 <nil> 0xc000110060}]
	*/
	//结构体方法调用
	/*root.Right.Left.Print()
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
	fmt.Println()*/

}
