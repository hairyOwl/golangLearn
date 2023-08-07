/**
*@Date: 星期一 2023/8/7 22:12
*@File: learn03
*@Author: owl
*@Description: 逻辑运算符
**/

package LearnOperator

import "fmt"

// UseLogical 逻辑运算符
func UseLogical() {
	a := true
	b := true
	c := false

	fmt.Println("a=", a, "b=", b, "c=", c)
	fmt.Println("a&&b", a && b)
	fmt.Println("a&&c", a && c)
	fmt.Println("a||b", a || b)
	fmt.Println("a||c", a || c)
	fmt.Println("!(a&&b)", !(a && b))

}
