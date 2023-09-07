/**
*@Date: 星期二 2023/9/5 20:29
*@File: urlRetriever
*@Author: owl
*@Description: 接口的定义和实现
**/

package urlGet

import (
	"fmt"
)

type Retriever struct {
	Contents string
}

// 系统常用接口 stringer
func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever: {Contents=%s}", r.Contents)
}

func (r *Retriever) Read(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *Retriever) Write(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}
func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

// Get 实现了使用者定义的方法就是实现了接口
func (r *Retriever) Get(s string) string {
	return r.Contents
}
