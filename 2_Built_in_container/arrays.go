/**
*@Date: 星期三 2023/8/16 22:33
*@File: arrays
*@Author: owl
*@Description: 数组
**/

package main

import "fmt"

func printArray(arr [5]int) {
	arr[0] = 100            //值类型 不会改变参数的实际值，函数体内生效
	for i, v := range arr { //下标、值    i  i,v  _,v
		fmt.Println(i, v)
	}
}

func printArrayP(arr *[5]int) {
	arr[0] = 100            //会改变参数的实际值
	for i, v := range arr { //下标、值    i  i,v  _,v
		fmt.Println(i, v)
	}
}

func main() {
	/*
		数组的定义格式
	*/
	var arr1 [5]int // [0,0,0,0,0]
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{2, 3, 4, 7, 8}
	//二维数组
	var grid [5][4]float64

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	/*
		数组遍历
	*/
	for i := 0; i < len(arr3); i++ { //不推荐
		fmt.Println(arr3[i])
	}

	printArray(arr1)
	//printArray(arr2) //[3]int  [5]int是不同的值类型
	fmt.Println(arr1)

	printArrayP(&arr3)
	fmt.Println(arr3)

}
