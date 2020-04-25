package BLC

import (
	"fmt"
	"os"
)

//发起交易
func (cli *CLI) send(from, to, amount []string, nodeId string) {

	if !dbExists(nodeId) {
		fmt.Println("数据库不存在")
		os.Exit(1)
	}
	blockchain := BlockChainObject(nodeId)
	defer blockchain.DB.Close()
	if len(from) != len(to) {
		fmt.Println("交易参数有误，请检查一致性.....")
		os.Exit(1)
	}
	blockchain.MineNewBlock(from, to, amount,nodeId)
}
