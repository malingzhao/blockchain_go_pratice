package BLC

import (
	"flag"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"os"
)

//对blockchain的命令行进行管理
//client对象
type CLI struct {
}

//用法展示
func PrintUsage() {
	fmt.Println("Usage:")
	//创建区块链
	fmt.Printf("\tcreateblockchain  --address ADDRESS --创建区块链\n")
	//添加区块
	fmt.Printf("\taddblcok --data DATA  --添加区块\n")
	//打印完整的区块信息
	fmt.Println("\tprintchain  --打印完整区块信息\n")
}

//初始化区块链
func (cli *CLI) createBlockChain(address string) {
	CreateBlockChainWithGensisBlcok(address)
}

//添加区块
func (cli *CLI) addBlock(txs []*Transaction) {

	blockchain := BlockChainObject()
	blockchain.AddBlock(txs)
	//获取到blockchain的对象实力
}

//打印区块链完整信息
func (cli *CLI) printChain() {
	//判断数据库是否存在
	if !dbExists() {
		fmt.Println("数据库不存在")
		os.Exit(1)
	}
	blockchain := BlockChainObject()
	blockchain.PrintChain()
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
	createBLCWithGensisBlockCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)

	//数据参数处理
	//添加区块
	flagAddBlockArg := addBlockCmd.String("data", "send 100 btc to player", "添加区块数据")

	flagCreateBlockchainArg := createBLCWithGensisBlockCmd.String("address", "troytan", "指定系统奖励的旷工地址")
	//创建区块链的时候指定的旷工的地址

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
		if err := createBLCWithGensisBlockCmd.Parse(os.Args[2:]); nil != err {

			log.Panicf("parse printChainCmd failed!%v\n", err)

		}
	default:
		//没有传递任何命令或者传递的命令不在上级的命令列表之中
		PrintUsage()
		os.Exit(1)
	}
	//添加区块信息
	if addBlockCmd.Parsed() {
		if *flagAddBlockArg == "" {
			PrintUsage()
		}
		//调用
		cli.addBlock([]*Transaction{})

	}

	//输出区块链信息
	if printChainCmd.Parsed() {
		cli.printChain()
	}
	//创建区块连命令
	if createBLCWithGensisBlockCmd.Parsed() {
		if *flagCreateBlockchainArg == "" {
			PrintUsage()
			os.Exit(1)
		}

		cli.createBlockChain(*flagCreateBlockchainArg)
	}

}

//获取一个blockchain的对象
func BlockChainObject() *BlockChain {
	//获取DB
	db, err := bolt.Open(dbName, 0600, nil)
	//获取tip
	if nil != err {
		log.Printf("open the db [%s] failed! %v\n", dbName, err)
	}

	var tip []byte
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if nil != b {
			tip = b.Get([]byte("l"))
		}
		return nil
	})

	if nil != err {
		log.Panicf("get the blockchain object failed!%v\n", err)
	}
	return &BlockChain{DB: db, Tip: tip}
}
