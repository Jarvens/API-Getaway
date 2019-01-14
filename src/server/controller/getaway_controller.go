// auth: kunlun
// date: 2019-01-14
// description:  网关控制器
package controller

import (
	"common"
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
