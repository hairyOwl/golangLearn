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

// 寻找最长不含有重复字符串的子串,并返回子串长度
func longStr(str string) int {
	start := 0                         //子串的开始下标
	maxLen := 0                        //子串长度
	lastOccurred := make(map[byte]int) //字符最近重复的下标

	for i, char := range []byte(str) { //[]byte(str) 字符串转为字符数组
		fmt.Println("i=", i, str[:i])
		fmt.Println("start=", start, " maxLen=", maxLen)
		fmt.Println("lastOccurred=", lastOccurred)
		lastI, ok := lastOccurred[char]
		if ok && lastOccurred[char] >= start { //遍历的子串出现重复时
			start = lastI + 1 //开始指标后移到重复后
		}
		if maxLen < (i + 1 - start) { //如果当前不重复子串大于maxlen
			maxLen = i + 1 - start
		}
		lastOccurred[char] = i //遍历字符最近一次出现的下标更新
	}

	return maxLen
}

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

	/*
		寻找最长不含有重复字符串的子串,并返回子串长度
	*/
	fmt.Println(longStr("aaabcbbbb"))
}
