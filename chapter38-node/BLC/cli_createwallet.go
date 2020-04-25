package BLC

import "fmt"

//创建钱包集合
func (cli *CLI) CreatwWallets( nodeId string  ){

	fmt.Println("nodeId",nodeId)
	wallets :=NewWallets(nodeId ) //创建一个集合对象
	wallets.CreateWallet(nodeId)
	fmt.Printf("wallets:%v\n",wallets)
}
