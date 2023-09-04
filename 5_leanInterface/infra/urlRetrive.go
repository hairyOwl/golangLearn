/**
*@Date: 星期一 2023/9/4 18:26
*@File: urlRetrive
*@Author: owl
*@Description: 接口
**/

package infra

import (
	"io"
	"net/http"
)

type Retrieve struct {
}

func (Retrieve) Get(url string) string {
	// 返回url 的response body
	//1. get response
	resp, err := http.Get(url)
	if err != nil {
		panic(err) //抛出异常
	}

	//3.关闭
	defer resp.Body.Close() //释放资源

	//2. 读取response
	bytes, _ := io.ReadAll(resp.Body) //([]bytes,error)
	return string(bytes)
}
