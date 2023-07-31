/*
@Time : 2023/7/31 11:01
@Author : 小哈
@File : cLearn02
@note: 特殊常量 iota的使用
*/
package constLearn

import "fmt"

// iota在const关键字出现时将被重置为0
const i_a = iota
const i_b = iota

// const中每新增一行常量声明将使iota计数一次
const (
	i_c = iota
	i_d = iota
	i_e = iota
)

// 1.跳值使用法
const (
	i0 = iota
	i1 = iota
	_  //回收 实现跳值
	i3 = iota
)

// 2.插值使用法
const (
	a1 = iota
	a2 = 3.14
	a3 = iota
)

// 3.表达式隐式使用法 iota的省略
const ( //0 1 2
	b1 = iota * 2
	b2 = iota
	b3 = iota
)
const ( //0 3 6 9 4
	c1 = iota * 2
	c2 = iota * 3
	c3 //继承最近一行非空表达式
	c4
	c5 = iota
)

// 单行使用法
const ( // 0 3 1 4 2
	d1, d2 = iota, iota + 3
	d3, d4
	d5 = iota
)

// LearnIota iota的使用
func LearnIota() {
	fmt.Println(i_a, i_b, i_c, i_d, i_e) // 0 0
	fmt.Println(i0, i1, i3)
	fmt.Println(a1, a2, a3)
	fmt.Println(b1, b2, b3)
	fmt.Println(c1, c2, c3, c4, c5)
	fmt.Println(d1, d2, d3, d4, d5)

}
