/**
*@Date: 星期日 2023/8/27 22:26
*@File: traversal
*@Author: owl
*@Description: 包和封装（通过拆树的结构体和方法）
**/

package tree

// Traverse 遍历树 中序遍历 左中右
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()

}
