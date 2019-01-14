// auth: kunlun
// date: 2019-01-14
// description:
package conf

import "time"

type ApiInfo struct {
	ApiName    string `json:"api_name"`    //接口名称
	RequestUrl string `json:"request_url"` //请求地址
	MethodType string `json:"method_type"` //请求方式
	GroupName  string `json:"group_name"`  //分组名称
	GroupId    int    `group_id`           //分组id
	Status     bool   `json:"status"`      //接口状态
}

type ApiGroup struct {
	GroupName string    `json:"group_name"` //分组名称
	GroupId   int       `json:"group_id"`   //分组id
	ApiInfo   []ApiInfo `json:"api_info"`   //接口信息
}

type GateWayInfo struct {
	GateWayName string     `json:"gateway_name"`  //网关名称
	IpWhiteList []string   `json:"ip_white_list"` //白名单
	IpBlackList []string   `json:"ip_black_list"` //黑名单
	TimeOut     int        `json:"timeout"`       //超时时间
	ApiList     []ApiInfo  `json:"api_list"`      //api列表
	GroupList   []ApiGroup `json:"group_list"`    //接口分组
	UpdateTime  time.Time  `json:"update_time"`   //更新时间
	CreateTime  time.Time  `json:"create_time"`   //创建时间
	Status      bool       `json:"status"`        //网关状态
}
