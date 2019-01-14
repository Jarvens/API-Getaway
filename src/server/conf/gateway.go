// auth: kunlun
// date: 2019-01-14
// description: 网关配置
package conf

import (
	"conf"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type GateWayInfo struct {
	GateWayName      string   `json:"gateway_name"`        //网关名称
	GateWayAlias     string   `json:"gateway_alias"`       //网关别名
	Status           string   `json:"status"`              //网关状态  on  off
	ApiConfPath      string   `json:"api_conf_path"`       //接口配置路径
	ApiGroupConfPath string   `json:"api_group_conf_path"` //接口分组配置路径
	StrategyConfPath string   `json:"strategy_conf_path"`  //策略配置路径
	IpLimitType      string   `json:"ip_limit_type"`       //Ip限制类型
	IpWhiteList      []string `json:"ip_white_list"`       //Ip白名单
	UpdateTime       string   `json:"update_time"`         //更新时间
	CreateTime       string   `json:"create_time"`         //创建时间
}

// 网关解析
func ParseGateWayInfo(path string) map[string]*GateWayInfo {
	gatewayInfo := make(map[string]*GateWayInfo)
	dirPath, err := conf.GetDir(path)
	if err == nil {
		pthSep := string(os.PathSeparator)
		for _, p := range dirPath {
			gateway := &GateWayInfo{}
			c, err := ioutil.ReadFile(p + pthSep + "gateway.conf")
			if err != nil {
				continue
			}
			err = yaml.Unmarshal(c, &gateway)
			if err != nil {
				continue
			}
			_, ok := gatewayInfo[gateway.GateWayAlias]
			if ok {
				continue
			}
			gatewayInfo[gateway.GateWayAlias] = gateway
		}
	}
	return gatewayInfo
}
