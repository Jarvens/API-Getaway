// auth: kunlun
// date: 2019-01-14
// description:
package common

import (
	"encoding/json"
	"log"
)

func StringConv(data interface{}) string {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Panic(err)
	}
	return string(jsonBytes)
}

func JsonByteConv(data interface{}) []byte {
	jsonByte, err := json.Marshal(data)
	if err != nil {
		log.Panic(err)
	}
	return jsonByte
}
