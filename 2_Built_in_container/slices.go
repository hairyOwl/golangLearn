/**
*@Date: 星期四 2023/8/17 21:38
*@File: slices
*@Author: owl
*@Description:切片
**/

package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	/*
		切片
	*/
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6]", arr[2:6]) //[2 3 4 5]
	fmt.Println("arr[:6]", arr[:6])   //[0 1 2 3 4 5]
	fmt.Println("arr[2:]", arr[2:])   //[2 3 4 5 6 7]
	fmt.Println("arr[:]", arr[:])     //[0 1 2 3 4 5 6 7]

	/*
		切片是一个视图
	*/
	s1 := arr[2:]
	s2 := arr[:]
	//直接修改 arr的值
	updateSlice(s1)
	fmt.Println("s1", s1) //[100 3 4 5 6 7]
	updateSlice(s2)
	fmt.Println("s2", s2)   //[100 1 100 3 4 5 6 7]
	fmt.Println("arr", arr) //[100 1 100 3 4 5 6 7]

	/*
		Reslice
	*/
	fmt.Println("Reslice")
	fmt.Println(s2) //[100 1 100 3 4 5 6 7]
	s2 = s2[:5]
	fmt.Println(s2) //[100 1 100 3 4]
	s2 = s2[2:]
	fmt.Println(s2) //[100 3 4]

	/*
		Slice的扩展 Extending slice
	*/
	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	s1 = arr[2:6]
	s2 = s1[3:5] //[s1[3],s1[4]]
	//fmt.Println(s1[4]) //ERROR : index out of range
	fmt.Println(arr)
	fmt.Println(s1)                                                   //[2 3 4 5]
	fmt.Println(s2)                                                   //[5] x   [5 6] √
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1)) // 4 6
	fmt.Printf("s2=%v len(s2)=%d cap(s2)=%d\n", s2, len(s2), cap(s2)) // 2 3
	fmt.Println(s1[3:7])                                              //slice bounds out of range

}
