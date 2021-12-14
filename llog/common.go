package llog

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
	"ymz465/go-horizon/config"
	"ymz465/go-horizon/helper"
)

var logger *log.Logger

func Init() {
	if !helper.CheckFileExists(config.LogDir) {
		// 日志目录不存在新建
		err := os.MkdirAll(config.LogDir, os.ModePerm)
		if err != nil {
			log.Fatal("Can't create " + config.LogDir + " dir, please check file system right")
		}
	}
}

//GetLoggerFile get log file
func GetLoggerFile(logFilePath string) (*os.File, error) {
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
		newName := fmt.Sprintf("%s/d15_%s.log", accessLogDir, time.Now().Format("2006-01-02"))
		os.Rename(logFilePath, newName)
	}
	return os.Create(logFilePath)
}

//GetAccessLogFile create new access log file
func GetAccessLogFile() (*os.File, error) {
	// Logging to a file.
	today := time.Now().Format("2006-01-02")
	logPath := filepath.Join(config.LogDir, "access.log")
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
		newName := fmt.Sprintf("%s/access_%s_%s.log", config.LogDir, today, helper.RandString(8))
		os.Rename(logPath, newName)
	}
	return os.Create(logPath)
}

//Print 记录日志记录
func Print(v ...interface{}) {
	logger.Println(v...)
}

//Println 记录日志记录
func Println(v ...interface{}) {
	logger.Println(v...)
}

//Printf 记录日志记录
func Printf(format string, v ...interface{}) {
	logger.Printf(format, v...)
}
