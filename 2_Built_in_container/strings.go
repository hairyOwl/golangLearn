/**
*@Date: 星期三 2023/8/23 22:19
*@File: strings
*@Author: owl
*@Description:
**/

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "吃葡萄，eat" //utf-8
	fmt.Println(s, len(s))
	fmt.Printf("%s\n", []byte(s))
	for _, b := range []byte(s) {
		fmt.Printf("%x ", b) //[e5 90 83]【吃】 [e8 91 a1]【葡】 [e8 90 84]【萄】 [ef bc 8c]【，】 65 61 74
	}
	fmt.Println()
	for i, ch := range s { //unicode ch是rune
		fmt.Printf("【%d %x】", i, ch)
	}
	fmt.Println()
	fmt.Println("rune count", utf8.RuneCountInString(s)) //rune count 7
	byteL := []byte(s)
	for len(byteL) > 0 { //对s逐字符转换
		ch, size := utf8.DecodeRune(byteL)
		byteL = byteL[size:]            //byte数组切片 切掉转换后的编码 如【吃】就是切掉 [e5 90 83]
		fmt.Printf("【%d %c】", size, ch) //【3 吃】【3 葡】【3 萄】【3 ，】【1 e】【1 a】【1 t】
	}
	fmt.Println()

	for i, ch := range []rune(s) { //新开一个rune数组放s的内容 rune数组每个字符占4个字节
		fmt.Printf("【%d,%c】", i, ch) //【0,吃】【1,葡】【2,萄】【3,，】【4,e】【5,a】【6,t】
	}
	fmt.Println()

}
