package config

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/caicaispace/gohelper/netx"
	"github.com/caicaispace/gohelper/setting"
	"github.com/lxmgo/config"
)

type configIni struct {
	Config config.ConfigInterface
}

var (
	configService     *configIni
	onceConfigService sync.Once
)

func GetIniInstance() *configIni {
	onceConfigService.Do(func() {
		configService = &configIni{}
		configService.loadConfigFile()
	})
	return configService
}

func (cs *configIni) ConfigGetHost() string {
	localIp := setting.Server.Host
	if runtime.GOOS == "linux" {
		localIp = netx.LocalIP()
	}
	return localIp
}

func (cs *configIni) loadConfigFile() {
	newConfig, err := config.NewConfig(cs.getConfigFilePath() + "conf.ini")
	if err != nil {
		fmt.Println("config file read fail:" + err.Error())
		os.Exit(0)
	}
	cs.Config = newConfig
}

func (cs *configIni) getConfigFilePath() string {
	confDefaultPath := flag.String("config", setting.Server.RootPath+"/config/", "config file path")
	confPath := *confDefaultPath
	if runtime.GOOS == "windows" {
		if (*confDefaultPath)[len(*confDefaultPath)-1:] != "/" {
			confPath = confPath + "/"
		}
	}
	return confPath
}
