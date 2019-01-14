// auth: kunlun
// date: 2019-01-14
// description: 配置文件解析
package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func ParseConfigInfo() GlobalConfig {
	var g GlobalConfig
	err := yaml.Unmarshal([]byte(Configure), &g)
	if err != nil {
		panic("全局网关配置错误!")
	}
	path, err := GetDir(g.GateConfigPath)
	if err != nil {
		panic("网关路径配置错误!")
	}

	gateWayList := GateWayList(path)
	g.GateWayList = gateWayList
	return g

}

// 查询网关列表
func GateWayList(path []string) []GateWayInfo {
	gateWayList := make([]GateWayInfo, 0)
	PthSep := string(os.PathSeparator)
	for _, p := range path {
		var gateWayInfo GateWayInfo
		c, err := ioutil.ReadFile(p + PthSep)
		if err != nil {
			panic("网关路径错误")
		}
		err = yaml.Unmarshal(c, &gateWayInfo)
		if err != nil {
			panic("网关配置读取错误")
		}

		if gateWayInfo.Status != "on" {
			continue
		}

		gateWayInfo.ApiList = GetApiList(gateWayInfo.ApiConfigPath)
		gateWayInfo.StrategyList = GetStrategyList(gateWayInfo.StrategyConfigPath)
		gateWayList = append(gateWayList, gateWayInfo)
	}

	return gateWayList
}

// 读取接口列表
func GetApiList(path string) []ApiInfo {

	var api []ApiInfo
	c, err := ioutil.ReadFile(path)
	if err != nil {
		panic("接口配置路径错误 :" + path)
	}

	err = yaml.Unmarshal(c, &api)
	if err != nil {
		panic("接口配置错误:" + path)
	}

	return api
}

// 策略集合
func GetStrategyList(path string) []StrategyInfo {
	var strategy []StrategyInfo

	c, err := ioutil.ReadFile(path)
	if err != nil {
		panic("策略配置路径错误:" + path)
	}
	err = yaml.Unmarshal(c, &strategy)
	if err != nil {
		panic("策略配置错误:" + path)
	}

	return strategy
}
