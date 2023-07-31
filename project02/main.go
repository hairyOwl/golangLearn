/*
@Time : 2023/6/26 17:03
@Author : 小哈
@File : main
@note: 数据结构的主入口
*/
package main

import (
	"project02/constLearn"
)

func main() {
	/*
		变量
	*/
	//数据类型
	//var i int = 1                                       //整形  1 uit8是1字节; 1 int32 是4个字节;int 会根据系统的位数动态变即win32是int32
	//var b bool = false                                  //布尔
	//var by rune = 1                                     //字节型
	//fmt.Println(unsafe.Sizeof(i), b, unsafe.Sizeof(by)) //获取内存大小32是8的4倍 1字节是8位

	//类型零值 和类型变量
	//typeLearn.LearnTypeUnSet()
	//typeLearn.LearnTypeOtherName()

	//变量声明
	//typeLearn.ForNPrint()

	//类型转换
	//typeLearn.ForNChange()

	// ForNSee 变量可见性规则
	//typeLearn.ForNSee()
	//fmt.Println(typeLearn.Car)
	//fmt.Println(typeLearn.str) //私有变量不能访问

	/*
		常量
	*/
	//常量的声明
	//constLearn.PrintConst()
	//特殊变量iota的使用
	constLearn.LearnIota()
}
