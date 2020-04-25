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

//发起交易
func (cli *CLI) send(){
	if !dbExists() {
		fmt.Println("数据库不存在")
		os.Exit(1)
	}
	blockchain :=BlockChainObject()
	defer blockchain.DB.Close()
	blockchain.MineNewBlock()
}

//用法展示
func PrintUsage() {
	fmt.Println("Usage:")
	//创建区块链
	fmt.Printf("\tcreateblockchain  -address ADDRESS --创建区块链\n")
	//添加区块
	fmt.Printf("\taddblcok --data DATA  --添加区块\n")
	//打印完整的区块信息
	fmt.Printf("\tprintchain  --打印完整区块信息\n")
	//通过命令行转账
	fmt.Printf("\tsend  -from FROM -to TO -amount AMOUNT -- 发起转账\n")

	fmt.Printf("\t转账参数说明\n")

	fmt.Printf("\t\t-from FROM --转账原地址\n")
	fmt.Printf("\t\t-to TO --转账目标地址\n")
	fmt.Printf("\t\t-amoount AMOUNT 转账金额\n")
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

	//创建区块的时候指定旷工的地址接收奖励
	createBLCWithGensisBlockCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)

	sendCmd:= flag.NewFlagSet("send",flag.ExitOnError)



	//数据参数处理
	//添加区块
	flagAddBlockArg := addBlockCmd.String("data", "send 100 btc to player", "添加区块数据")

	flagCreateBlockchainArg := createBLCWithGensisBlockCmd.String("address", "troytan", "指定系统奖励的旷工地址")
	//创建区块链的时候指定的旷工的地址

	//发起交易参数
	flagSendFromArg :=sendCmd.String("from","","转账源地址")
	flagSendToArg :=sendCmd.String("to","","目标地址")
	flagSendAmountArg :=sendCmd.String("amount","","转账金额")

	//判断命令
	switch os.Args[1] {
	case "send":
		if err :=sendCmd.Parse(os.Args[2:]);nil!=err{
			log.Panicf("parse sendCmd failed!%v\n", err)
		}

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
	//发起转账
	if sendCmd.Parsed(){
		if *flagSendFromArg==""{
			fmt.Println("源地址不能为空")
		}
		if *flagSendToArg==""{
			fmt.Println("目标地址不能为空")
		}
		if *flagSendAmountArg==""{
			fmt.Println("转账金额不能为空")
		}
		fmt.Printf("\tFROM:[%s]\n",JsonToSlice(*flagSendFromArg))
		fmt.Printf("\tTO:[%s]\n",JsonToSlice(*flagSendToArg))
		fmt.Printf("\tAMOUNT:[%s]\n",JsonToSlice(*flagSendAmountArg))
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
