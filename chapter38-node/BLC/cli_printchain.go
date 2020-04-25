package BLC

import (
	"fmt"
	"os"
)

//打印区块链完整信息
func (cli *CLI) printChain(nodeId string ) {
	//判断数据库是否存在
	if !dbExists(nodeId) {
		fmt.Println("数据库不存在")
		os.Exit(1)
	}
	blockchain := BlockChainObject(nodeId)
	blockchain.PrintChain()
	defer blockchain.DB.Close()
}

