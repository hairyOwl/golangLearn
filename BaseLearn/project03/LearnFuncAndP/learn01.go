/**
*@Date: 星期一 2023/8/14 21:38
*@File: learn05
*@Author: owl
*@Description: 函数与操作符
**/

package LearnFuncAndP

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 13 / 3 = 4....1
func Div(a, b int) (int, int) { //建议这种返回形式
	return a / b, a % b
}
func Div1(a, b int) (x, y int) {
	x = a / b
	y = a % b
	return x, y
}

// Eval +-*/
func Eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil //nil 空错误
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		x, _ := Div(a, b)
		return x, nil
	default:
		//panic("未知操作符" + op) //中断执行
		return 0, fmt.Errorf("未知操作符%s", op)
	}
}

func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// Apply eval的函数版本
func Apply(op func(int, int) int, a, b int) int {
	//获取函数名字
	p := reflect.ValueOf(op).Pointer()    //获取函数指针
	opName := runtime.FuncForPC(p).Name() //获取函数名字
	fmt.Println("响应的函数是", opName, "参数是", a, b)
	return op(a, b)
}

// Sum 可变参数列表
func Sum(numL ...int) int {
	s := 0
	for i := range numL {
		s += numL[i]
	}
	return s
}
