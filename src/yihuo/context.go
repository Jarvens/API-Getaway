// auth: kunlun
// date: 2019-01-15
// description: 上下文
package yihuo

import "conf"

type Context struct {
	GateWayInfo  GateWay
	StrategyInfo Strategy
	ApiInfo      Api
	Rate         map[string]Rate
	VisitCount   *Count
}

type GateWay struct {
	GateWayAlias  string   `json:"gateway_alias"`
	GateWayStatus string   `json:"gateway_status"`
	IpLimitType   string   `json:"ip_limit_type"`
	IpWhiteList   []string `json:"ip_white_list"`
	IpBlackList   []string `json:"ip_black_list"`
}

type Strategy struct {
	StrategyId    string               `json:"strategy_id"`
	Auth          string               `json:"auth"`
	BasicUserName string               `json:"basic_user_name"`
	BasicPassword string               `json:"basic_password"`
	ApiKey        string               `json:"api_key"`
	IpLimitType   string               `json:"ip_limit_type"`
	IpWhiteList   []string             `json:"ip_white_list"`
	IpBlackList   []string             `json:"ip_black_list"`
	RateLimitList []conf.RateLimitInfo `json:"rate_limit_list"`
}

type Api struct {
	RequestURL     string               `json:"request_url"`
	BackendPath    string               `json:"backend_path" yaml:"backend_path"`
	ProxyURL       string               `json:"proxy_url"`
	ProxyMethod    string               `json:"proxy_method"`
	IsRaw          bool                 `json:"is_raw"`
	ProxyParams    []conf.Param         `json:"proxy_params"`
	ConstantParams []conf.ConstantParam `json:"constant_params"`
	Follow         bool                 `json:"follow"`
}
