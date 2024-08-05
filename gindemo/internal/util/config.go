package util

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func InitConfig() {
	var env string
	if len(os.Args) < 2 {
		env = "dev"
	} else {
		env = os.Args[1]
	}
	fmt.Println("env", env)
	var configFile string
	switch env {
	case "dev":
		configFile = "./config/dev.json"
	case "prod":
		configFile = "./config/prod.json"
	}
	v := viper.New()
	v.SetConfigFile(configFile)
	v.SetConfigType("json")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}

	mac_count := v.GetInt("mac_count")
	fmt.Println("mac_count", mac_count)

	baseurl := v.GetString("baseurl")
	fmt.Println("baseurl", baseurl)

	Config = v
}
