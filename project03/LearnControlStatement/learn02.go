/**
*@Date: 星期四 2023/8/10 22:10
*@File: learn02
*@Author: owl
*@Description: For循环  goto break continue
**/

package LearnControlStatement

import "fmt"

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

}
