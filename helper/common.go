package helper

import (
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

//RandStringWithCharset 通过给定字符串 生成随机字符串
func RandStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

//RandString 随机生成一个字符串
func RandString(length int) string {
	return RandStringWithCharset(length, charset)
}

// CheckFileExists 检查文件或者目录是否存在
func CheckFileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

//GetCurrentExeDir 获取当前程序目录
func GetExeDir() (string, error) {
	//获取当前程序目录
	exeDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	ExeDir := strings.Replace(exeDir, "\\", "/", -1)
	return ExeDir, nil
}
