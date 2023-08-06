/*
@Time : 2023/6/12 17:19
@Author : 小哈
@File : show
@note: import多包导入顺序
*/
package show2

import (
	"fmt"
	"project01/learn1"
)

func init() {
	fmt.Println("show init")
}

func Show() {
	//依赖learn 导入顺序变化 不论在main中的调用顺序，会先导入learn.go
	//在main中执行Show函数 相当于导入了两次learn.go 1show.go 2.main.go 最后learn.init只执行一次
	learn1.Learn()
	fmt.Println("b")
}
