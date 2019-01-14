// auth: kunlun
// date: 2019-01-14
// description:
package conf

import "io/ioutil"

type GlobalConfig struct {
	Host            string `json:"host"`              //主机地址
	Port            string `json:"port"`              //主机端口
	GateWayConfPath string `json:"gateway_conf_path"` //网关配置路径
	LoginName       string `json:"login_name"`        //登录名
	LoginPassword   string `json:"login_password"`    //登录密码
}

var GlobalConf GlobalConfig

// 写入配置文件
func WriteConfgToFile(path string, data []byte) bool {
	err := ioutil.WriteFile(path, data, 0666)
	if err != nil {
		panic(err)
	}
	return true
}
