package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

//对blockchain的命令行进行管理
//client对象
type CLI struct {
	BC *BlockChain //blockchain对象
}

//用法展示
func PrintUsage() {
	fmt.Println("Usage:")
	//创建区块链
	fmt.Printf("createblockchain --创建区块链\n")
	//添加区块
	fmt.Printf("addblcok --data DATA  --添加区块\n")
	//打印完整的区块信息
	fmt.Println("printchain  --打印完整区块信息\n")
}

//初始化区块链
func (cli *CLI) createBlockChain() {
	CreateBlockChainWithGensisBlcok()
}

//添加区块
func (cli *CLI) addBlock(data string) {
	//获取到blockchain的对象实力
	cli.BC.AddBlock([]byte(data))
}

//打印区块链完整信息
func (cli *CLI) printChain() {
	cli.BC.PrintChain()
}

//参数数量的检测函数
func IsValidArgs() {
	if len(os.Args) < 2 {
		PrintUsage()
		//直接退出
		os.Exit(1)
	}

}

//运行命令行函数
func (cli *CLI) Run() {
	IsValidArgs()
	//新建相关命令
	//添加区块
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	//输出区块链完整信息
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	//输出区块链
	createBlockChainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)

	//数据参数处理
	flagAddBlockArg := addBlockCmd.String("data", "send 100 btc to player", "添加区块数据")

	//判断命令
	switch os.Args[1] {
	case "addblock":
		if err := addBlockCmd.Parse(os.Args[2:]); nil != err {
			log.Panicf("parse addBlockCmd failed!%v\n", err)
		}
	case "printchain":
		if err := printChainCmd.Parse(os.Args[2:]); nil != err {
			log.Panicf("parse printChainCmd: failed!%v\n", err)
		}
	case "createblockchain":
		if err := createBlockChainCmd.Parse(os.Args[2:]); nil != err {
			log.Panicf("parse printChainCmd failed!%v\n", err)
		}
	default:
		//没有传递任何命令或者传递的命令不在上级的命令列表之中
		PrintUsage()
		os.Exit(1)
	}
	//添加区块信息
	if addBlockCmd.Parsed(){
		if *flagAddBlockArg==""{
			PrintUsage()
			//调用
			cli.addBlock(*flagAddBlockArg)

		}
	}

	//输出区块链信息
	if printChainCmd.Parsed() {
		cli.printChain()
	}
	//创建区块连命令
	if createBlockChainCmd.Parsed(){
		cli.createBlockChain()
	}

}
