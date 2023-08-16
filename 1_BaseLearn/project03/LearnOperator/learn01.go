/**
*@Date:  星期日 2023/8/6 21:58
*@File:
*@Author: owl
*@Description: 算术运算符
**/

package LearnOperator

import "fmt"

// UseCalc 算术运算符
func UseCalc() {
	var a = 2
	var b = 3
	fmt.Println("a", a, "b=", b)
	fmt.Println("a+b", a+b)
	fmt.Println("a-b", a-b)
	fmt.Println("a*b", a*b)
	fmt.Println("a/b", a/b)
	fmt.Println("a%b", a%b)
	a++
	fmt.Println("a++", a)
	a--
	fmt.Println("a--", a)

}
