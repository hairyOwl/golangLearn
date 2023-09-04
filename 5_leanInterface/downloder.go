/**
*@Date: 星期六 2023/9/2 23:45
*@File: downloder.go
*@Author: owl
*@Description: 主入口
**/

package main

import (
	"5_leanInterface/testing"
	"fmt"
)

// 用接口动态配置接收者
type retriever interface {
	Get(string) string
}

// 获取接口实列
func getRetrieve() retriever {
	return testing.Retrieve{}
	//return infra.Retrieve{}
}

func main() {
	//var retrieve infra.Retrieve = infra.Retrieve{}  指定了类型 没有完全解耦
	retrieve := getRetrieve() //使用接口
	fmt.Println(retrieve.Get("https://www.baidu.com"))
}
