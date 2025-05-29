package config

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
	"ymz465/go-horizon/helper"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	ExeDir        string
	LogDir        string
	AccessLog     string
	RunLog        string
	HTMLDir       string
	PIDFile       string
	TLSPublicKey  string
	TLSPrivateKey string
)

// Init 初始化
func Init() {
	var err error
	ExeDir, err = helper.GetExeDir()
	if err != nil {
		log.Println("获取当前程序运行目录错误: ", err)
		os.Exit(0)
	}
	viper.SetConfigName("config")
	// for my dev
	ymzDEV := os.Getenv("YMZ_DEV")
	if ymzDEV != "" {
		ExeDir = "."
	}
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
	TLSPublicKey = filepath.Join(ExeDir, "certs", "server.pem")
	TLSPrivateKey = filepath.Join(ExeDir, "certs", "server.key")
	//
	if !helper.CheckFileExists(LogDir) {
		// log dir not exists, create it
		err := os.MkdirAll(LogDir, os.ModePerm)
		if err != nil {
			log.Fatal("Can't create " + LogDir + " dir, please check file system right")
		}
	}
	if viper.GetString("runmode") != "debug" {
		//
		logFile, err := getLoggerFile(RunLog)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		log.SetOutput(logFile)
		// set gin access log
		gin.DisableConsoleColor()
		gin.SetMode(gin.ReleaseMode)
		// set log file
		accessLogFile, err := getAccessLogFile()
		if err != nil {
			log.Fatal("Create log file error: ", err)
		}
		gin.DefaultWriter = io.MultiWriter(accessLogFile)
	}
}

// getLoggerFile get log file
func getLoggerFile(logFilePath string) (*os.File, error) {
	if helper.CheckFileExists(logFilePath) {
		//exists file, judge does file need rename by file size
		fInfo, err := os.Stat(logFilePath)
		if err != nil {
			log.Fatal("Log file size check error: ", err)
		}
		if fInfo.Size() < 1024*1024*1000 {
			// less than 1G then append model open file
			return os.OpenFile(logFilePath, os.O_APPEND|os.O_WRONLY, 0600)
		}
		accessLogDir := filepath.Dir(logFilePath)
		// more than 1G then rename exists file
		newName := fmt.Sprintf("%s/run_%s.log", accessLogDir, time.Now().Format("2006-01-02"))
		os.Rename(logFilePath, newName)
	}
	return os.Create(logFilePath)
}

// getAccessLogFile create new access log file
func getAccessLogFile() (*os.File, error) {
	// Logging to a file.
	today := time.Now().Format("2006-01-02")
	logPath := filepath.Join(LogDir, "access.log")
	if helper.CheckFileExists(logPath) {
		//exists file, judge does file need rename by file size
		fInfo, err := os.Stat(logPath)
		if err != nil {
			log.Fatal("Log file size check error: ", err)
		}
		if fInfo.Size() < 1024*1024*1000 {
			// less than 1G then append model open file
			return os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY, 0600)
		}
		// more than 1G then rename exists file
		newName := fmt.Sprintf("%s/access_%s_%s.log", LogDir, today, helper.RandString(8))
		os.Rename(logPath, newName)
	}
	return os.Create(logPath)
}
