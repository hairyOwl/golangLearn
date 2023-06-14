/*
@Time : 2023/6/9 14:50
@Author : 小哈
@File : main
@note: 基础结构
*/
package main

import (
	console "fmt" //fmt自定义别名console
	"project01/learn1"
	_ "project01/show2" //只执行show2包的init函数
	. "time"            //引用时省略包名
)

func init() {
	console.Println("main init") //导包后执行
}

// main()函数
func main() {
	console.Println("main start")
	console.Print(DateTime)
	//LearnFunc()
	//show2.Show()
	learn1.Learn()
	console.Println("main end")
}
