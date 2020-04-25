package main

import "blockchain_go_pratice/chapter14-arg-json/BLC"

//启动
func main() {
	//
	//bc := BLC.CreateBlockChainWithGensisBlcok()
	//bc.AddBlock([]byte("a send 100 eth to b"))
	//bc.AddBlock([]byte("b send 100 eth to c"))
	//bc.PrintChain()

	cli := BLC.CLI{}
	cli.Run()
}
