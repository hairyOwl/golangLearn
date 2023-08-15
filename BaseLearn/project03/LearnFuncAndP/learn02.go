/**
*@Date: 星期二 2023/8/15 22:09
*@File: learn02
*@Author: owl
*@Description: 指针
**/

package LearnFuncAndP

func Swap(a, b *int) {
	*b, *a = *a, *b //b指向的内容=a指向的内容 ， a指向的内容=b指向的内容
}

func Swap2(a, b int) (int, int) { //推荐这种
	return b, a
}
