package main

import (
	"blockchain_go_pratice/chapter29-blockcahin-wallets-non-test/BLC"
	"fmt"
)

//启动
func main() {


		wallets :=BLC.NewWallets() //生成集合对象
		wallets.CreateWallet()
		fmt.Printf("wallets:%v\n",wallets.Wallets)


}
