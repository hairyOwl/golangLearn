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
	"time"
)

// 定义接口
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("www.baidu.com")
}
func post(poster Poster) {
	poster.Post("https://www.baidu.com",
		map[string]string{
			"name": "百度",
			"type": "搜索工具",
		})
}

// 接口的组合
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post("https://www.baidu.com",
		map[string]string{
			"contents": "a session",
		})
	return s.Get("https://www.baidu.com")
}

// 接口实现类的方法
func inspect(r Retriever) {
	fmt.Printf("%T,%v\n", r, r)
	switch v := r.(type) {
	case *urlGet.Retriever:
		fmt.Println("Contents", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent", v.UserAgent)

	}
}

func main() {
	var r Retriever
	r = &urlGet.Retriever{Contents: "testing 接口的实现 "}
	//fmt.Printf("%T,%v\n", r, r) //urlGet.Retriever,{testing 接口的实现 }
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
		Url:       "https://www.baidu.com",
	}
	//fmt.Printf("%T,%v\n", r, r) //*real.Retriever,&{https://www.baidu.com Mozilla/5.0 1m0s}
	inspect(r)

	//Type assertion  .()类型的名字 取得interface具体的类型
	realRetriever := r.(*real.Retriever) //严格
	fmt.Println(realRetriever.TimeOut)
	if urlGetRetriever, ok := r.(*urlGet.Retriever); ok { //不严格
		fmt.Println(urlGetRetriever)
	} else {
		fmt.Println("not a urlGetRetriever")
	}

	//fmt.Println(download(r))
	s := &urlGet.Retriever{Contents: "接口的组合 "}
	fmt.Println(session(s))
}
