/**
*@Date: 星期日 2023/8/27 23:13
*@File: entry
*@Author: owl
*@Description: 通过定义别名 扩展已有类型
**/

package main

import (
	"3_tree/queue"
	"fmt"
)

func main() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(4)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

}
