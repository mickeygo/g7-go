package config

import (
	"fmt"

	"github.com/mickeygo/go-g7/pkg/errcode"
	"github.com/spf13/viper"
)

var (
	Env = "dev" // dev|prod
)

func init() {
	path := "config." + Env + ".json"
	viper.SetConfigName(path)     // 配置文件名称(含扩展名)
	viper.AddConfigPath("./conf") // 设置文件路径，可以设置多个

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err := fmt.Errorf("[%d] %s", errcode.Err_File_Config_Not_Found, err.Error())
			panic(err)
		} else {
			panic(err)
		}
	}
}
