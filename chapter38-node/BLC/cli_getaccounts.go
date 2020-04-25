package BLC

import "fmt"

/**
获取地址列表
 */

func (cli *CLI) GetAccounts(nodeId string){
	wallets := NewWallets(nodeId)
	fmt.Println("\t账号列表")
	for key, _:= range wallets.Wallets{
		fmt.Printf("\t[%s]\n",key)
	}

	fmt.Printf("NODE_ID:%v\n", nodeId)
}
