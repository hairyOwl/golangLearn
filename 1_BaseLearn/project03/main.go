/**
*@Date: 星期日 2023/8/6 21:49
*@File:
*@Author: owl
*@Description: 运算符 主程序
**/

package main

import (
	"fmt"
	"math"
	"project03/LearnFuncAndP"
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
	/*//条件语句
	fmt.Println("条件语句")
	LearnControlStatement.UseIf()

	//选择语句
	fmt.Println("选择语句")
	LearnControlStatement.UseSwitch()

	//For循环
	fmt.Println("For循环")
	LearnControlStatement.UseFor()

	//goto break continue
	LearnControlStatement.UseKeyWord()*/

	/*
		函数
	*/
	//func operator
	//fmt.Println(LearnOperator.Eval(1, 2, "a"))
	if result, err := LearnFuncAndP.Eval(3, 4, "*"); err != nil {
		//出错
		fmt.Println(err)
	} else {
		//正确
		fmt.Println(result)
	}
	fmt.Println(LearnFuncAndP.Div(12, 5))
	//函数式编程
	fmt.Println(LearnFuncAndP.Apply(LearnFuncAndP.Pow, 3, 2))
	fmt.Println(LearnFuncAndP.Apply(
		func(i int, i2 int) int { //函数参数直接写匿名函数
			return int(math.Pow(float64(i), float64(i2)))
		}, 3, 3))
	fmt.Println(LearnFuncAndP.Sum(1, 3, 5))

	/*
		指针
	*/
	a, b := 3, 4
	LearnFuncAndP.Swap(&a, &b)
	fmt.Println("a=", a, "b=", b)
	a, b = LearnFuncAndP.Swap2(a, b)
	fmt.Println("a=", a, "b=", b)

}
