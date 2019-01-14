// auth: kunlun
// date: 2019-01-14
// description:
package dao

import "server/conf"

// 创建网关信息
func AddGateWay(gateWayName, gateWayAlias string) bool {

	gateWay := conf.ParseGateWayInfo(conf.GlobalConf.GateWayConfPath)
	_, ok := gateWay[gateWayAlias]
	if ok {
		return false
	} else {

	}
	return true
}
