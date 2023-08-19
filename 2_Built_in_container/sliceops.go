/**
*@Date: 星期六 2023/8/19 22:39
*@File: sliceops
*@Author: owl
*@Description: 切片的操作
**/

package main

import "fmt"

func printSlice(s []int) {
	fmt.Println(s, " len=", len(s), " cap=", cap(s))
}

func main() {

	/*
		切片的创建
	*/
	fmt.Println("切片的创建")
	var s []int    //切片的类型零值是nil
	fmt.Println(s) //[]
	for i := 0; i < 10; i++ {
		printSlice(s)
		s = append(s, i)
	}
	s1 := []int{2, 4, 6, 8} //是{2,4,6,8}数组的view
	printSlice(s1)
	s2 := make([]int, 16)     //知道声明切片的len，但不打算创建的时候赋值
	s3 := make([]int, 16, 32) //知道声明切片的len cap，但不打算创建的时候赋值
	printSlice(s2)
	printSlice(s3)

	/*
		切片的复制
	*/
	fmt.Println("切片的复制")
	copy(s2, s1) //dst目标 src被复制
	printSlice(s2)

	/*
		切片的删除
	*/
	fmt.Println("切片中的元素删除")
	s2 = append(s2[:3], s2[4:]...) //删除s2[3]
	printSlice(s2)

	fmt.Println("获取切片的头")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("获取切片的尾")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)

}
