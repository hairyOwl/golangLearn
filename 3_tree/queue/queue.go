/**
*@Date: 星期日 2023/8/27 23:00
*@File: queue
*@Author: owl
*@Description: 通过定义别名 扩展已有类型
**/

package queue

type Queue []int

func (q *Queue) Push(v int) { //改变原值所以用指针
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
