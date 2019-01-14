// auth: kunlun
// date: 2019-01-14
// description:
package conf

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var (
	Configure string
)

type ApiInfo struct {
	ApiName    string `json:"api_name"`    //接口名称
	RequestUrl string `json:"request_url"` //请求地址
	MethodType string `json:"method_type"` //请求方式
	GroupName  string `json:"group_name"`  //分组名称
	GroupId    int    `group_id`           //分组id
	Status     bool   `json:"status"`      //接口状态
}

type ApiGroup struct {
	GroupName string `json:"group_name"` //分组名称
	GroupId   int    `json:"group_id"`   //分组id
	Apis      Api    `json:"api_info"`   //接口信息
}

type GateWayInfo struct {
	GateWayName        string     `json:"gateway_name"`         //网关名称
	IpWhiteList        []string   `json:"ip_white_list"`        //白名单
	IpBlackList        []string   `json:"ip_black_list"`        //黑名单
	TimeOut            int        `json:"timeout"`              //超时时间
	ApiList            Api        `json:"api_list"`             //api列表
	GroupList          []ApiGroup `json:"group_list"`           //接口分组
	UpdateTime         time.Time  `json:"update_time"`          //更新时间
	CreateTime         time.Time  `json:"create_time"`          //创建时间
	Status             string     `json:"status"`               //网关状态
	ApiConfigPath      string     `json:"api_config_path"`      //接口配置路径
	StrategyConfigPath string     `json:"strategy_config_path"` //策略配置地址
	StrategyList       Strategy   `json:"strategy_list"`        //策略列表
}

type Strategy struct {
	Strategy []StrategyInfo `json:"strategy"`
}

type Api struct {
	ApiInfo []ApiInfo `json:"apiinfo" yaml:"apiinfo,omitempty"`
}

type GlobalConfig struct {
	Host           string        `json:"host"`                //地址
	Port           string        `json:"port"`                //端口
	GateConfigPath string        `json:"gateway_config_path"` //配置路径
	GateWayList    []GateWayInfo `json:"gateway_list"`        //网关集合
}

type StrategyInfo struct {
	StrategyName  string          `json:"strategy_name"`   //策略名称
	StrategyId    string          `json:"strategy_id"`     //策略id
	Auth          string          `json:"auth"`            //授权
	BasicUserName string          `json:"basic_user_name"` //默认用户名
	BasicPassword string          `json:"basic_password"`  //默认密码
	ApiKey        string          `json:"api_key"`         //api key
	IpLimitType   string          `json:"ip_limit_type"`   //限制类型
	IpWhiteList   []string        `json:"ip_white_list"`   //白名单集合
	IpBlackList   []string        `json:"ip_black_list"`   //黑名单集合
	RateLimitList []RateLimitInfo `json:"rate_limit_list"` //限流集合
	UpdateTime    string          `json:"update_time"`     //更新时间
	CreateTime    string          `json:"create_time"`     //创建时间
}

type RateLimitInfo struct {
	Allow     bool   `json:"allow"`      //是否受限
	Period    string `json:"period"`     //尝试次数
	Limit     int    `json:"limit"`      //限制次数
	Priority  int    `json:"priority"`   //
	StartTime int    `json:"start_time"` //开始时间
	EndTime   int    `json:"end_time"`   //结束时间
}

func init() {
	//初始化为 空字符
	Configure = ""
	fmt.Printf("初始化全局配置路径: %v", Configure)
}

func GetDir(path string) (files []string, err error) {
	files = make([]string, 0)
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("读取配置文件错误：%v", err)
		return nil, err
	}

	separator := string(os.PathSeparator)
	for _, val := range dir {
		if val.IsDir() {
			file, err1 := ioutil.ReadDir(path + separator + val.Name())
			if err1 != nil {
				fmt.Printf("读取文件错误：%v", err1)
			}
			for _, v1 := range file {
				files = append(files, path+separator+val.Name()+separator+v1.Name())
				fmt.Printf("文件名称为：%v", v1.Name())
			}
		}
	}
	return files, nil
}

func ReadConfigure(filePath string) (err error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	Configure = string(content)
	fmt.Printf("读取全局配置文件为: %s\n", Configure)
	return
}
