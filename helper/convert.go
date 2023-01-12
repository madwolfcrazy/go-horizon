package helper

import "strconv"

//StringToUint
func StringToUint(s string) uint {
	a, _ := strconv.ParseUint(s, 10, 64)
	return uint(a)
}

//StringToInt
func StringToInt(s string) int {
	a, _ := strconv.ParseInt(s, 10, 64)
	return int(a)
}

//IFToInt64 转为int64类型
func IFToInt64(value interface{}) (val int64) {
	switch value := value.(type) { // shadow
	case int:
		val = int64(value)
	case int8:
		val = int64(value)
	case int16:
		val = int64(value)
	case int32:
		val = int64(value)
	case int64:
		val = value
	case uint:
		val = int64(value)
	case uint8:
		val = int64(value)
	case uint16:
		val = int64(value)
	case uint32:
		val = int64(value)
	case uint64:
		return int64(value)
	case string:
		// for testing and other apps - numbers may appear as strings
		var err error
		if val, err = strconv.ParseInt(value, 10, 64); err != nil {
			return int64(val)
		}
	default:
		return int64(val)
	}
	return int64(val)
}
