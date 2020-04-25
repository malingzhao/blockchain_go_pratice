package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	msg := "Man"
	//
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Printf("encode result:%v\n", encoded)

	//编码
	b, err := base64.StdEncoding.DecodeString("TWFu")
	if nil != err {
		panic(err)
	}
	fmt.Printf("decide result:%s\n", b)
}
