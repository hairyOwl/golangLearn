/**
*@Date: 星期一 2023/8/21 22:53
*@File: Maps
*@Author: owl
*@Description: Map结构
**/

package main

import (
	"fmt"
)

func main() {
	/*
		map的声明
	*/
	fmt.Println("map的声明")
	m := map[string]string{ //map[key的类型]value的类型
		"num":    "001",
		"name":   "张三",
		"course": "sport",
		"score":  "A",
	}

	m2 := make(map[string]int)                           //map[]   m2 == empty map
	var m3 map[int]int                                   //map[] m3 == nil
	fmt.Println("m=", m, len(m), " m2=", m2, " m3=", m3) //map[course:sport name:张三 num:001 score:A]

	/*
		map的遍历
	*/
	fmt.Println("map的遍历")
	for k, v := range m {
		fmt.Println(k, v)
	}

	/*
		获取值
	*/

	fmt.Println("获取值")
	mName, ok := m["name"] //ok bool 判断是否存在这个key
	fmt.Println(mName, ok)

	//ame, ok := m["ame"]  //不存在的key
	if ame, ok := m["ame"]; ok {
		fmt.Println(ame) //zero value
	} else {
		fmt.Println("这个key不存在")
	}

	/*
		删除元素
	*/
	fmt.Println("删除元素")
	name, ok := m["name"]
	fmt.Println("before delete", name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println("after delete", name, ok)
}
