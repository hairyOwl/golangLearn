/**
*@Date: 星期日 2023/8/6 21:49
*@File:
*@Author: owl
*@Description: 运算符 主程序
**/

package main

import (
	"fmt"
	"project03/LearnControlStatement"
	_ "project03/LearnOperator"
)

func main() {
	/*
	   运算符
	*/
	////算术运算符
	//fmt.Println("算术运算符")
	//LearnOperator.UseCalc()
	//
	////关系运算符
	//fmt.Println("关系运算符")
	//LearnOperator.UseRelation()
	//
	////逻辑运算符
	//fmt.Println("逻辑运算符")
	//LearnOperator.UseLogical()
	//
	////按位运算符
	//fmt.Println("按位运算符")
	//LearnOperator.UseBitwise()

	/*
		控制语句
	*/
	//条件语句
	fmt.Println("条件语句")
	LearnControlStatement.UseIf()

	//选择语句
	fmt.Println("选择语句")
	LearnControlStatement.UseSwitch()

}
