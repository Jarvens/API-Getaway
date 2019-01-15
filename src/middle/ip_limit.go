// auth: kunlun
// date: 2019-01-15
// description: IpLimit
package middle

import (
	"net/http"
	"strings"
	"yihuo"
)

func IpLimit(context *yihuo.Context, res http.Response, req *http.Request) (bool, string) {
	remoteAddr := req.RemoteAddr
	remoteIp := InterceptIP(remoteAddr, ":")
	if !globalIpLimit(context, remoteIp) {
		return false, "[全局] 非法ip请求"
	} else if !strategyIpLimit(context, remoteIp) {
		return false, "[策略] 非法ip 请求"
	}
	return true, ""
}

// Ip 拦截
func InterceptIP(str, substr string) string {
	result := strings.Index(str, substr)
	var res string
	if result > 7 {
		res = str[:result]
	}
	return res
}

func globalIpLimit(context *yihuo.Context, remoteIp string) bool {

	// 检查是否在黑名单集合中
	if context.GateWayInfo.IpLimitType == "black" {
		for _, ip := range context.GateWayInfo.IpBlackList {
			if ip == remoteIp {
				return false
			}
		}
		return true
		//检查是否在白名单结合中
	} else if context.GateWayInfo.IpLimitType == "white" {
		for _, ip := range context.GateWayInfo.IpWhiteList {
			if ip == remoteIp {
				return true
			}
		}
		return false
	}
	return true
}

func strategyIpLimit(context *yihuo.Context, remoteIp string) bool {

	//检查ip是否在策略黑名单中
	if context.StrategyInfo.IpLimitType == "black" {
		for _, ip := range context.StrategyInfo.IpBlackList {
			if ip == remoteIp {
				return false
			}
		}
		return true
		//检查ip是否在策略白名单中
	} else if context.StrategyInfo.IpLimitType == "white" {
		for _, ip := range context.StrategyInfo.IpWhiteList {
			if ip == remoteIp {
				return true
			}
		}
		return false
	}
	return true
}
