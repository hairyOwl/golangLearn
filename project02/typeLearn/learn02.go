/*
@Time : 2023/7/26 10:29
@Author : 小哈
@File : learn02.go
@note: 变量声明
*/
package typeLearn

import (
	"fmt"
	"reflect"
)

// ForNPrint 变量声明
func ForNPrint() {
	// 单个变量声明和赋值 函数体内外都行
	var a int32
	var b int32 = 456
	a = 123 // 变量赋值
	b = 111

	// 分组声明和赋值 函数体内外都行
	var (
		c string
		d int32  = 1
		e string = "abc"
	)
	var i, j, k int32 = 0, 1, 10
	var x, y, z = 1, "test", 12.1

	//省略 var 声明变量 只能在函数体内
	n, m := 1, 2
	t, h := 1, "test2"

	//下划线 变量声明
	var p, _, q = 1, 2, 3

	println(a, b)
	println(c, d, e)
	println(i, j, k)
	println(x, y, z)
	println(reflect.TypeOf(x).Name(), reflect.TypeOf(y).Name(), reflect.TypeOf(z).Name())
	println(n, m)
	println(t, h)
	println(p, q)

}

// ForNChange 类型转换
func ForNChange() {
	var a int32 = 2
	var b float32 = 3.01
	c := float32(a) //类型转换
	d := int32(b)   //损失精度

	//不能转换 不兼容的类型不能转换
	//var flag bool = true
	//g := int32(flag)

	fmt.Println(a, b, c, d)
	fmt.Println(reflect.TypeOf(c).Name(), reflect.TypeOf(d).Name())

}

// ForNSee 变量可见性规则
var str = "小写变量"
var Car = "汽车"

func ForNSee() {

}
