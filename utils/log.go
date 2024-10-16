package utils

import (
	"encoding/json"
)

// ToByte :nodoc:
func ToByte(i interface{}) []byte {
	bt, _ := json.Marshal(i)
	return bt
}

// Dump to json using json marshal
func Dump(i interface{}) string {
	return string(ToByte(i))
}
