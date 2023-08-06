/*
@Time : 2023/6/9 14:50
@Author : 小哈
@File : structL
@note: 基础结构
*/
//程序所属包
package main

//导入依赖包
import "fmt"

// 常量定义
const NAME string = "golang"

// 全局变量的声明与赋值
var a string = "go语言"

// 一般类型声明 相当于类型别名
type aInt int

var num aInt = 1

// 结构的声明
type Learn struct {
}

// 接口声明
type Ilearn interface {
}

// 函数定义
func LearnFunc() {

	fmt.Println("learnStruct", num)
}
