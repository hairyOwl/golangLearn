/**
*@Date: 星期日 2023/8/6 21:49
*@File:
*@Author: owl
*@Description: 运算符 主程序
**/

package main

import (
	"fmt"
	"project03/LearnOperator"
)

func main() {
	//算术运算符
	fmt.Println("算术运算符")
	LearnOperator.UseCalc()

	//关系运算符
	fmt.Println("关系运算符")
	LearnOperator.UseRelation()

}
