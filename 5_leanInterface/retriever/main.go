/**
*@Date: 星期二 2023/9/5 20:21
*@File: main
*@Author: owl
*@Description: 接口的定义和实现
**/

package main

import (
	"5_leanInterface/retriever/real"
	"5_leanInterface/retriever/urlGet"
	"fmt"
)

// 定义接口
type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("www.baidu.com")
}

func main() {
	var r Retriever
	r = urlGet.Retriever{Contents: "testing 接口的实现 "}
	r = real.Retriever{Url: "https://www.baidu.com"}
	fmt.Println(download(r))
}
