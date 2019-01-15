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
	ApiName        string          `json:"api_name" yaml:"api_name"`
	GroupID        int             `json:"group_id" yaml:"group_id"`
	RequestURL     string          `json:"request_url" yaml:"request_url"`
	RequestMethod  []string        `json:"request_method" yaml:"request_method"`
	BackendID      int             `json:"backend_id" yaml:"backend_id"`
	ProxyURL       string          `json:"proxy_url" yaml:"proxy_url"`
	ProxyMethod    string          `json:"proxy_method" yaml:"proxy_method"`
	IsRaw          bool            `json:"is_raw" yaml:"is_raw"`
	ProxyParams    []Param         `json:"proxy_params" yaml:"proxy_params"`
	ConstantParams []ConstantParam `json:"constant_params" yaml:"constant_params"`
	Follow         bool            `json:"follow" yaml:"follow"`
}

type ApiGroup struct {
	Group []ApiGroup `json:"group"`
}

type ApiGroupInfo struct {
	GroupID   int    `json:"group_id" yaml:"group_id"`
	GroupName string `json:"group_name" yaml:"group_name"`
}

type GateWayInfo struct {
	GateWayName        string    `json:"gateway_name"`  //网关名称
	GateWayAlias       string    `json:"gateway_alias"` //网关别名
	IPLimitType        string    `json:"ip_limit_type" `
	IpWhiteList        []string  `json:"ip_white_list"`        //白名单
	IpBlackList        []string  `json:"ip_black_list"`        //黑名单
	TimeOut            int       `json:"timeout"`              //超时时间
	ApiList            Api       `json:"api_list"`             //api列表
	GroupList          ApiGroup  `json:"group_list"`           //接口分组
	UpdateTime         time.Time `json:"update_time"`          //更新时间
	CreateTime         time.Time `json:"create_time"`          //创建时间
	GateWayStatus      string    `json:"gateway_status"`       //网关状态
	ApiConfigPath      string    `json:"api_config_path"`      //接口配置路径
	StrategyConfigPath string    `json:"strategy_config_path"` //策略配置地址
	StrategyList       Strategy  `json:"strategy_list"`        //策略列表
	BackendList        Backend
}

type Backend struct {
	Backend []BackendInfo `json:"backend" yaml:"backend"`
}

type BackendInfo struct {
	BackendID   int    `json:"backend_id" yaml:"backend_id"`
	BackendName string `json:"backend_name" yaml:"backend_name"`
	BackendPath string `json:"backend_path" yaml:"backend_path"`
}

type Strategy struct {
	Strategy []StrategyInfo `json:"strategy"`
}

type Api struct {
	Apis []ApiInfo `json:"apis"`
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

type Param struct {
	Key              string `json:"key" yaml:"key"`
	KeyPosition      string `json:"key_position" yaml:"key_position"`
	NotEmpty         bool   `json:"not_empty" yaml:"not_empty"`
	ProxyKey         string `json:"proxy_key" yaml:"proxy_key"`
	ProxyKeyPosition string `json:"proxy_key_position" yaml:"proxy_key_position"`
}

type ConstantParam struct {
	Position string `json:"position" yaml:"position"`
	Key      string `json:"key" yaml:"key"`
	Value    string `json:"value" yaml:"value"`
}

func init() {
	//初始化为 空字符
	Configure = ""
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
	//fmt.Printf("Global config info: %s\n", Configure)
	return
}
