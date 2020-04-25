package main

import (
	"blockchain_go_pratice/chapter05-boltdb/BLC"
)

//启动
func main() {

	bc := BLC.CreateBlockChainWithGensisBlcok()
	bc.AddBlock([]byte("a send 100 eth to b"))
	bc.AddBlock([]byte("b send 100 eth to c"))

	//fmt.Println("blockchain:%v\n", bc.Blocks[0])
	////上链 height int64, prevBlcokHash []byte, data []byte
	//bc.AddBlock(bc.Blocks[len(bc.Blocks) -1].Height +1,bc.Blocks[len(bc.Blocks)-1].Hash,
	//	[]byte("Alice send 10 btc to Bob") )
	////上链 height int64, prevBlcokHash []byte, data []byte
	//bc.AddBlock(bc.Blocks[len(bc.Blocks) -1].Height +1,bc.Blocks[len(bc.Blocks)-1].Hash,
	//	[]byte("Bod send 5 t troytan") )
	//
	//for _,block := range bc.Blocks{
	//	fmt.Printf("prevBlcokHash:%x , block:%x\n",block.PrevBlockHash, block.Hash)
	//}
}
