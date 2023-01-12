package helper

import "strings"

func StrInSlice(list []string, a string) bool {
	if a == "" {
		return false
	}
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//ParseSquareBrackets 过滤方括号
func ParseSquareBrackets(s1 string) string {
	s1 = strings.Replace(s1, "[", "", -1)
	s1 = strings.Replace(s1, "]", "", -1)
	return s1
}
