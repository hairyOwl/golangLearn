/**
*@Date: 星期一 2023/8/7 22:21
*@File: learn04
*@Author: owl
*@Description:位移运算符
**/

package LearnOperator

import "fmt"

// UseBitwise 按位运算符
func UseBitwise() {
	a := 1 //  01
	b := 2 //  10

	fmt.Println("a=", a, "b=", b)
	fmt.Println("a&b", a&b)
	fmt.Println("a|b", a|b)
	fmt.Println("a^b", a^b)
	fmt.Println("a<<1", a<<1) // 01 --> 10
	fmt.Println("b>>1", b>>1) // 10 --> 01

}
