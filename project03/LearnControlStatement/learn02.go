/**
*@Date: 星期四 2023/8/10 22:10
*@File: learn02
*@Author: owl
*@Description: For循环  goto break continue
**/

package LearnControlStatement

import (
	"fmt"
)

// UseFor For循环
func UseFor() {
	//无限循环
	/*for {
		fmt.Println("in for")
		time.Sleep(1 * time.Second)
	}*/

	//有条件的循环
	for i := 1; i < 10; i += 2 {
		fmt.Println("i=", i)

	}

	//列表打印
	a := []string{"香蕉", "苹果", "梨"}
	for key, value := range a { //key可用_省略
		fmt.Println("key=", key, "value=", value)
	}
}

// UseKeyWord  goto break continue
func UseKeyWord() {
	/*
		1.goto
	*/
	fmt.Println("goto")
	goto one
	fmt.Println("中间代码块") //跳过不打印
one:
	fmt.Println("这里是 code one")
	/*
		cricle: //循环写法
			fmt.Println("无限循环")
			time.Sleep(1 * time.Second)
		goto cricle
	*/

	/*
		2.break 跳出循环体
	*/
	fmt.Println("break")
	for i := 1; i < 3; i++ {
		for j := 1; j < 5; j++ {
			if j == 3 {
				break //跳出本次循环
			}
			fmt.Println("i=", i, "j=", j)
		}
	}

	/*
		3.continue 跳出一次循环
	*/
	fmt.Println("continue")
	for i := 1; i < 3; i++ {
		for j := 1; j < 5; j++ {
			if j == 3 {
				continue //跳出本次循环
			}
			fmt.Println("i=", i, "j=", j)

		}
	}

}
