package BLC

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
	"os"
)

//实现int64转成[]byte 的字节数组  任何时候都能使用 没有(block *Block)
func IntToHex(data int64) []byte {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, data)
	if err!=nil{
		log.Panic("int transact to []byte failed %v\n",err)
	}

	return buffer.Bytes()
}


//标准json转切片
func JsonToSlice(jsonString string)[]string{
	var strSlice []string

	//json
	if err :=json.Unmarshal([]byte(jsonString), &strSlice);err!=nil{
		log.Panicf("json to []string failed!%v\n",err)
	}
	return  strSlice
}


//参数数量的检测函数
func IsValidArgs() {
	if len(os.Args) < 2 {
		PrintUsage()
		//直接退出
		os.Exit(1)
	}

}


//反转字节数组
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

//string 转ripemd160
func StringToHash160(address string) []byte {
	pubKeyHash := Base58Decode([]byte(address))
	hash160 := pubKeyHash[:len(pubKeyHash)-addressCheckSumLen]
	return hash160
}
