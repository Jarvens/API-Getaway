// auth: kunlun
// date: 2019-01-14
// description:
package dao

import (
	"conf"
	conf2 "server/conf"
)

// 创建网关信息
func AddGateWay(gateWayName, gateWayAlias string) bool {

	gateWay := conf.ParseConfigInfo(conf2.GlobalConf.GateWayConfPath)

}
