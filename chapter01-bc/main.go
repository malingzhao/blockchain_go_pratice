package main

import (
	"blockchain_go_pratice/chapter01-btc/BLC"
	"fmt"
)

//启动
func main() {
	block := BLC.NewBlock(1,nil,[]byte("the first block testing"))
fmt.Println("the first block: %v\n", block)
}