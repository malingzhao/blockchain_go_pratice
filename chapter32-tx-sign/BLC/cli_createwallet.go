package BLC

import "fmt"

//创建钱包集合
func (cli *CLI) CreatwWallets(){
	wallets :=NewWallets() //创建一个集合对象
	wallets.CreateWallet()
	fmt.Printf("wallets:%v\n",wallets)
}
