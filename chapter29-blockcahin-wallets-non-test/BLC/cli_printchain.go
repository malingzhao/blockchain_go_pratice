package BLC

import (
	"fmt"
	"os"
)

//打印区块链完整信息
func (cli *CLI) printChain() {
	//判断数据库是否存在
	if !dbExists() {
		fmt.Println("数据库不存在")
		os.Exit(1)
	}
	blockchain := BlockChainObject()
	blockchain.PrintChain()
	defer blockchain.DB.Close()
}

