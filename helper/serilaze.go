package helper

import (
	"bytes"
	"encoding/gob"
)

//Serialization 序列化
func Serialization(o1 interface{}) []byte {
	buf := bytes.Buffer{}
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(o1)
	if err != nil {
		return nil
	}
	return buf.Bytes()
}

//Unserialization 反序列化
func Unserialization(b1 []byte, o1 interface{}) error {
	return nil
}
