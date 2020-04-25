package BLC

import (
	"fmt"
	"os"
)

//发起交易
func (cli *CLI) send(from, to, amount []string) {
	if !dbExists() {
		fmt.Println("数据库不存在")
		os.Exit(1)
	}
	blockchain := BlockChainObject()
	defer blockchain.DB.Close()
	blockchain.MineNewBlock(from, to, amount)
}