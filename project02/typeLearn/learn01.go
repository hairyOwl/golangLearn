/*
@Time : 2023/7/18 17:11
@Author : 小哈
@File : learn01
@note: 类型零值 和类型变量
*/
package typeLearn

import (
	"reflect"
	"unsafe"
)

// 类型零值
func LearnTypeUnSet() {
	//数值
	var i int32
	var j float32
	var d complex64 //复数类型
	//布尔
	var b bool
	//字符串
	var s string
	println("类型零值")
	println(i, ";", j, ";", b, ";", d, ";", s)

}

// 类型别名
type 整形32 int32 //别名可以是中文 在函数体外声明
func LearnTypeOtherName() {
	var i 整形32
	var j int32
	var d 整形32

	println(i, ";", j)
	println("\n类型别名")
	println(reflect.TypeOf(i).Name(), ";", unsafe.Sizeof(i))
	println(reflect.TypeOf(j).Name(), ";", unsafe.Sizeof(j))
	println(reflect.TypeOf(i))
	//println(i + j) //加了别名 类型不匹配
	println(i + d)

}
