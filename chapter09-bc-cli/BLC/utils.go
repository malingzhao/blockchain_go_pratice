package BLC

import (
	"bytes"
	"encoding/binary"
	"log"
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

