package BLC

import "fmt"

func (cli *CLI) getBalance(from string) {
	//查找改地址的UTXO
	//获取区块链对象
	blockchain := BlockChainObject()
	defer blockchain.DB.Close()
	amount := blockchain.getBlance(from)
	fmt.Printf("\t地址[%s]的余额：[%d]\n", from, amount)
}

