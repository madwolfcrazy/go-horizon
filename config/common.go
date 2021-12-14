package config

import (
	"log"
	"os"
	"path/filepath"
	"ymz465/go-horizon/helper"

	"github.com/spf13/viper"
)

var (
	ExeDir    string
	LogDir    string
	AccessLog string
	RunLog    string
	HTMLDir   string
	PIDFile   string
)

//Init 初始化
func Init() {
	ExeDir, err := helper.GetExeDir()
	if err != nil {
		log.Println("获取当前程序运行目录错误: ", err)
		os.Exit(0)
	}
	viper.SetConfigName("config")
	viper.AddConfigPath(ExeDir)
	if err := viper.ReadInConfig(); err != nil {
		log.Println("加载config.toml文件错误，请检查")
		os.Exit(0)
	}
	LogDir = filepath.Join(ExeDir, "logs")
	AccessLog = filepath.Join(LogDir, "access.log")
	RunLog = filepath.Join(LogDir, "run.log")
	PIDFile = filepath.Join(LogDir, "run.pid")
	HTMLDir = filepath.Join(ExeDir, "html")
	//
}
