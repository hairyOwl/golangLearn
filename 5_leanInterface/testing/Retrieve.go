/**
*@Date: 星期一 2023/9/4 19:37
*@File: Retrieve
*@Author: owl
*@Description: 接口的概念
**/

package testing

type Retrieve struct {
}

func (Retrieve) Get(url string) string {
	return "测试 Retrieve.Get"
}
