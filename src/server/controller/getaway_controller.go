// auth: kunlun
// date: 2019-01-14
// description:  网关控制器
package controller

import (
	"common"
	"conf"
	"fmt"
	"net/http"
)

// 创建网关
func AddGateWay(response http.ResponseWriter, request *http.Request) {

	fmt.Printf("request: %v", request)
	res := common.Response{}
	response.Write(common.JsonByteConv(res.Success(nil)))
}

// 编辑网关信息
func EditGateWay(response http.ResponseWriter, request *http.Request) {
	res := common.Response{}
	response.Write(common.JsonByteConv(res.Success("编辑网关信息")))
}

// 删除网关信息
func DeleteGateWay(response http.ResponseWriter, request *http.Request) {
	res := common.Response{}
	response.Write(common.JsonByteConv(res.Success("删除网关信息")))
}

// 网关集合
func List(response http.ResponseWriter, request *http.Request) {

	var strategyList = make([]conf.StrategyInfo, 1)
	var strategyInfo = conf.StrategyInfo{}
	strategyInfo.ApiKey = "ApiKey: kadsjfiurhughwsfdjkanfkjgfh89374ytg4h3fnr"
	strategyInfo.Auth = "已授权"
	strategyInfo.BasicUserName = "admin"
	strategyInfo.BasicPassword = "admin"
	strategyInfo.StrategyName = "基础策略"
	strategyInfo.StrategyId = "10001"
	strategyList[0] = strategyInfo
	res := common.Response{}
	response.Write(common.JsonByteConv(res.Success(strategyList)))
}
