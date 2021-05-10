//配置文件定义及加载
package core

import (
	"encoding/json"
	"fmt"
	"os"
	"version/logger"
	"version/utils"
)

type Conf struct {
	Host     string
	Port     int
	PidName  string
	PidPath  string
	ConfPath string
}

var ConfObj *Conf

func init() {
	ConfObj = &Conf{
		Host:    "localhost",
		Port:    8001,
		PidName: "pid",
		PidPath: "/tmp/",
	}

	ConfObj.reload()

}

func (c *Conf) String() string {
	return fmt.Sprintf(`-----------------conf begin------------
Host:%s
Port=%d
PidName=%s
PidPath=%s
-------------------conf end-------------
`,
		c.Host,
		c.Port,
		c.PidName,
		c.PidPath,
	)
}

func (c *Conf) reload() {
	if !utils.FileExist(C_CONF_PATH) {
		//logger.Warnf("ConfPath:%s is not exist, going to use default conf\n%s\n", C_CONF_PATH, c)
		logger.Warnf("ConfPath:%s is not exist, going to use default conf\n", C_CONF_PATH)
		return
	}

	data, err := os.ReadFile(C_CONF_PATH)
	if err != nil {
		logger.Exitf("os.ReadFile|%s failed,err=%v\n", C_CONF_PATH, err)
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		logger.Exitf("json.Unmarshal|%s is not right json file,err=%v\n", C_CONF_PATH, err)
	}

	//logger.Info("the reload conf is \n", c)

}

func SetConfHost(host string) {
	ConfObj.Host = host
}

func SetConfPort(port int) {
	ConfObj.Port = port
}

func GetConfAddr() string {
	return fmt.Sprintf("%s:%d", ConfObj.Host, ConfObj.Port)
}

func GetConfPidFile() string {
	return fmt.Sprintf("%s/version.%s-%s-%d", ConfObj.PidPath, ConfObj.PidName, ConfObj.Host, ConfObj.Port)
}
