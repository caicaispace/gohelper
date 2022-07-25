package config_test

import (
	"fmt"
	"testing"

	"github.com/caicaispace/gohelper/config"
	"github.com/caicaispace/gohelper/setting"
)

func Test_ConfigIniLoadConfig(t *testing.T) {
	setting.Server.RootPath = "/home/xxx/dev/xxx/gateway/cmd/gateway"
	config := config.GetIniInstance().Config
	// s, _ := json.MarshalIndent(config, "", "\t")
	// fmt.Print(string(s))
	fmt.Println(config.String("es::es.user"))
	fmt.Println(config.String("es::es.password"))
	fmt.Println(config.String("es::es.ip"))
}
