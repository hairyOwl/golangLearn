/**
*@Date: 星期二 2023/8/8 22:08
*@File: leanrn01
*@Author: owl
*@Description: 条件语句、选择语句
**/

package LearnControlStatement

import "fmt"

// UseIf 条件语句
func UseIf() {
	a := 1
	// if
	if a > 0 {
		fmt.Println("a>0")
	}

	//if ... else
	if a < 0 {
		fmt.Println("a<0")
	} else {
		fmt.Println("a>=0")
	}

	//if嵌套
	b := 3
	if b > 0 {
		fmt.Println("b>0")
		if b > 5 {
			fmt.Println("b>5")
		} else {
			fmt.Println("5>b>0")
		}
	}
}

// UseSelect 选择语句
func UseSelect() {

}
