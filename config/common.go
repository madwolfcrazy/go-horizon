package config

import (
	"os"
        "log"

	"github.com/spf13/viper"
)

//Init 初始化
func Init() {
    viper.SetConfigName("config")
        viper.AddConfigPath(".")
        if err := viper.ReadInConfig(); err != nil {
                log.Println("加载config.toml文件错误，请检查")
                os.Exit(0)
        }
}

