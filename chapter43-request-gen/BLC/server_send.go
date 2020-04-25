package BLC

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

//请发送文件

//3 发送请求
func sendMessage(to string, msg []byte) {
	fmt.Println("向服务器发送请求....")
	//1.连接服务器
	conn, err := net.Dial(PROTOCOL, to)

	if nil != err {
		log.Println("connect to server [%s] failed %v\n", err)
	}
	//要发送的数据
	_, err = io.Copy(conn, bytes.NewReader(msg))
	if nil != err {
		log.Panicf("add the data to conn failed!%v\n", err)
	}

}

//从区块链版本验证
func sendVersion(toAddress string, bc *BlockChain) {
	//1. 获取当前节点的区块的高度
	height := bc.GetHeight()
	fmt.Printf("height:%v\n",height)
	//2. 组装生成version
	versionData := Version{Height: int(height), AddrFrom: nodeAddress}
	//3. 数据的序列化
	data := gobEncode(versionData)
	//4 将命令与版本组成成完整的请求
	request := []byte{}
	request = append(commandToBytes(CMD_VERSION), data...)
	//4. 发送请求
	sendMessage(toAddress, request)
}

//
func sendGetBlocks(dataFrom string ) {

}

//发送获取区块的请求
func sendGetData(toAddress string, hash []byte) {
	//1 生成数据
	data := gobEncode(GetData{AddrFrom:nodeAddress, ID:hash})
	// 2 组装请求
	request := append(commandToBytes(CMD_GETBLOKCS), data...)
	//3 发送请求
	sendMessage(toAddress, request)
}

//向其他的节点展示
func sendINV(toAddress string,hash [][]byte) {
	//1 生成数据
	data :=gobEncode(INV{AddrFrom:nodeAddress, Hashes:hash})
	// 2 组装请求
	request:=append(commandToBytes(CMD_INV),data...)
	//3 发送请求
	sendMessage(toAddress, request)

}

//发送区块信息 向其他节点
func sendBlock(toAddress string, block []byte) {
	//1 生成数据
	data :=gobEncode(BlockData{AddrFrom:nodeAddress,Block:block})
	// 2 组装请求
	request:=append(commandToBytes(CMD_BLOCK),data...)
	//3 发送请求
	sendMessage(toAddress, request)
}
