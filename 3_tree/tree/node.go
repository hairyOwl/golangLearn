/**
*@Date: 星期六 2023/8/26 22:27
*@File: node
*@Author: owl
*@Description: 结构体
**/

package tree

import "fmt"

// Node 结构体声明
type Node struct {
	Value       int
	Left, Right *Node
}

// Print 结构体方法，方法名前是接收者(node TreeNode)相当于this
func (node *Node) Print() { //(node TreeNode)值传递
	fmt.Print(node.Value, " ")
}
func (node *Node) SetValue(v int) { //值传递不能修改值 所以为了修改value,接收者用地址
	if node == nil {
		fmt.Println("nil指针设置value 忽略")
		return
	}
	node.Value = v
}

// CreateNode 自定义工厂函数限定指定构造格式
func CreateNode(v int) *Node {
	return &Node{Value: v} //返回局部变量的地址
}
