package helper

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const charset2 = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" +
	"~!@#$%^&*()_+-=[]{}|;:,./<>?"

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

//RandStringHard 随机生成一个字符串
func RandStringHard(length int) string {
	return RandStringWithCharset(length, charset2)
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

//ReadPIDFile 获取pid
func ReadPIDFile(fileName string) (int, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return 0, err
	}
	pid, _ := strconv.Atoi(string(data))
	return pid, nil
}

// SHA1 计算SHA1
func SHA1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	l := fmt.Sprintf("%x", h.Sum(nil))
	return l
}

//Decimal 保留3位小数
func Decimal(value float64, num int) float64 {
	f1 := "%" + fmt.Sprintf(".%df", num)
	value, _ = strconv.ParseFloat(fmt.Sprintf(f1, value), 64)
	return value
}

func IntToBytes(n int64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(n))
	return b
}

//Trim 去除字符串前后空格 tab
func Trim(s1 string) string {
	s1 = strings.Trim(s1, " ")
	s1 = strings.Trim(s1, "\t")
	s1 = strings.Trim(s1, "\n")
	return s1
}

//UUID4 生成UUID4
func UUID4() string {
	u := uuid.New()
	return u.String()
}
