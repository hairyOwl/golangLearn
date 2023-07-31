/*
@Time : 2023/7/31 9:34
@Author : 小哈
@File : learn01
@note: 常量的声明
*/

package constLearn

import (
	"fmt"
	"reflect"
)

// 常量声明
const stuName string = "小王" //显式
const stuNum = 1121212      //隐式

// 组合声明
const (
	stuEmail    = "xxxx@xx.com"
	stuBirthday = "xxxx-xx-xx"
)

// 单行多个常量声明
const apple, banana string = "苹果", "香蕉" //显式
const a, b = 1, "this is 常量"            //隐式

// 内置表达式作为常量
const bLen = len(b)

// PrintConst 常量的声明
func PrintConst() {
	fmt.Println(stuName)
	fmt.Println(stuNum)
	fmt.Println(stuEmail, stuBirthday)
	fmt.Println(apple, banana)
	fmt.Println(a, reflect.TypeOf(a).Name(), b, reflect.TypeOf(b).Name())
	fmt.Println(bLen) //utf-8 中文字符占3个长度，英文字符占1个长度
}
