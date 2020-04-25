package BLC

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

//网络服务文件管理
//3000作为引导节点 主节点的地址
var knownNodes = []string{"localhost:3000"}

// 节点地址
var nodeAddress string

//启动服务
func startServer(nodeID string) {
	fmt.Printf("启动节点[%s]...\n", nodeID)
	//节点地址赋值
	nodeAddress := fmt.Sprintf("localhost:%s", nodeID)
	//1. 监听节点
	listen, err := net.Listen(PROTOCOL, nodeAddress)
	if err != nil {
		log.Panicf("listen address of %s  failed!:%v\n", nodeAddress, err)
	}
	defer listen.Close()
	//获取blockchain的对象
	bc:=BlockChainObject(nodeAddress)
	//两个节点 主节点负责保存数据 钱包节点负责发送请求 同步数据
	if nodeAddress != knownNodes[0] {
		//不是节点 发送请求  同步数据
		//sendMessage(knownNodes[0], nodeAddress)
		sendVersion(knownNodes[0], bc)

	}

	for {
		//生成连接 接受请求
		conn, err := listen.Accept()
		if nil != err {
			log.Panicf("accept connect failed%v!\n", err)
		}
		request, err := ioutil.ReadAll(conn)
		if nil != err {
			log.Panicf("Receive Message failed!%v\n", err)
		}
		//3. 处理请求
		fmt.Printf("Receive a Message:%v\n", request)
		//单独启动一个goroutinue进行请求处理
		go handleConnection(conn,bc)
	}
}

//master worker
//请求处理函数
func handleConnection(conn net.Conn, bc *BlockChain) {
	request, err := ioutil.ReadAll(conn)
	if nil != err {
		log.Panicf("Received a Request failed!%v\n", err)
	}
	cmd := bytesToCommand(request[:12])
	fmt.Printf("Receive a Command :%s\n", cmd)
	switch cmd {
	case CMD_VERSION:
		handlerVersion(request,bc)
	case CMD_GETDATA:
		handleGetData(request,bc)
	case CMD_GETBLOKCS:
		handleGetBlocks(request,bc)
	case CMD_INV:
		handleInv(request,bc)
	case CMD_BLOCK:
		handleBlock(request,bc)
	default:
		fmt.Println("Unknown command")
	}

}

//gob 编码
func gobEncode(data interface{}) []byte {
	var result bytes.Buffer
	gob.NewEncoder(&result)
	return result.Bytes()
}

//命令转换为请求([]byte)
func commandToBytes(command string) []byte {
	var bytes [COMMAND_LENGTH]byte
	for i, c := range command {
		bytes[i] = byte(c)
	}
	return bytes[:]
}

//请求处理函数
