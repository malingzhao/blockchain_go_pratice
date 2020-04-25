package BLC

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

//请求处理文件

func handlerVersion(request []byte, bc *BlockChain) {
	var buff bytes.Buffer
	var data Version
	fmt.Println("the request of version handle ....")
	//1. 解析请求
	dataBytes := request[:12]
	buff.Write(dataBytes)
	//2. 生成结构
	decoder := gob.NewDecoder(&buff)
	//3. 生层version结构

	if err := decoder.Decode(&data); nil != err {
		log.Panicf("Decode the version struct failed!%v\n", err)
	}
	//获取请求方法的区块高度
	versionHeight := data.Height
	//获取自身区块高度
	height := bc.GetHeight()
	//如果当前节点的区块高度 大于versionHeight
	//将当前节点的版本信息发送给请求节点
	if height > int64(versionHeight) {
		sendVersion(data.AddrFrom, bc)
		//如果当前节点的区块高度小于versionHeight 发起同步数据的请求
	} else if height < int64(versionHeight) {
		sendGetBlocks(data.AddrFrom)
	}
}

//GetBlock
func handleGetBlocks(request []byte, bc *BlockChain) {
	fmt.Println("the request of get blockcs handle...")
	var buff bytes.Buffer
	var data GetBlocks
	fmt.Println("the request of get blocks handle ....")
	//1. 解析请求
	dataBytes := request[:12]
	//2. 生成getblock结构
	buff.Write(dataBytes)
	decoder := gob.NewDecoder(&buff)
	if err := decoder.Decode(&data); nil != err {
		log.Panicf("Decode the get blocks struct failed!%v\n", err)
	}
	//3 获取区块链的所有区块哈希
	hashes := bc.GetBlockHases()
	sendINV(data.AddrFrom, hashes)
}

//INV
func handleInv(request []byte, bc *BlockChain) {
	var buff bytes.Buffer
	var data INV
	fmt.Println("the request of get inv  handle ....")
	//1. 解析请求
	dataBytes := request[:12]
	//2. 生成getData结构
	buff.Write(dataBytes)
	decoder := gob.NewDecoder(&buff)
	if err := decoder.Decode(&data); nil != err {
		log.Panicf("Decode the inv  struct failed!%v\n", err)
	}
	for _, hash := range data.Hashes {
		sendGetData(data.AddrFrom, hash)
	}
}

//Get Data
func handleGetData(request []byte, bc *BlockChain) {
	fmt.Println("the request of get blockcs handle...")
	var buff bytes.Buffer
	var data GetData
	fmt.Println("the request of get blocks handle ....")
	//1. 解析请求
	dataBytes := request[:12]
	//2. 生成getblock结构
	buff.Write(dataBytes)
	decoder := gob.NewDecoder(&buff)
	if err := decoder.Decode(&data); nil != err {
		log.Panicf("Decode the get blocks struct failed!%v\n", err)
	}
	//通过获取区块海西
	blockBytes := bc.GetBlock(data.ID)
	sendBlock(data.AddrFrom, blockBytes)
}

//Block
//接收到新的区块的时候进行处理
func handleBlock(request []byte, bc *BlockChain) {
	fmt.Println("the request of get  handle block handle...")
	var buff bytes.Buffer
	var data BlockData
	fmt.Println("the request of get blocks handle ....")
	//1. 解析请求
	dataBytes := request[:12]
	//2. 生成getblock结构
	buff.Write(dataBytes)
	decoder := gob.NewDecoder(&buff)
	if err := decoder.Decode(&data); nil != err {
		log.Panicf("Decode the get blockdata struct failed!%v\n", err)
	}
	//将接收到的区块添加到区块链中
	blockBytes := data.Block
	block := DeSerializeBlock(blockBytes)
	bc.AddBlock(block)

	//更新utxo
	utxoSet :=UTXOSet{bc}
	utxoSet.update()
}

//获取区块链的所有区块哈希
func (bc *BlockChain) GetBlockHases() [][]byte {
	var blockHashes [][]byte

	bcit := bc.Iterator()
	for {
		block := bcit.Next()
		blockHashes = append(blockHashes, block.Hash)
		if isBreakLoop(block.PrevBlockHash) {
			break
		}
	}
	return blockHashes
}
