/**
*@Date: 星期三 2023/8/30 23:18
*@File: zaptest
*@Author: owl
*@Description: go.mod 依赖管理
**/

package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction() //返回类型 (*Logger, error)
	logger.Warn("这是一条警告")

}
