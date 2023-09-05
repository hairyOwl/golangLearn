/**
*@Date: 星期二 2023/9/5 20:38
*@File: realRetriever
*@Author: owl
*@Description: 实现接口 打印网站body
**/

package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	Url       string
	UserAgent string
	TimeOut   time.Duration //时间段
}

func (r Retriever) Get(url string) string {
	resp, err := http.Get(r.Url)
	if err != nil {
		panic(err) //抛出异常
	}

	result, err := httputil.DumpResponse(resp, true)
	defer resp.Body.Close()
	if err != nil {
		panic(err) //抛出异常
	}
	return string(result)
}
