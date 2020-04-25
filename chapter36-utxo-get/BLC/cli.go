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

//用法展示
func PrintUsage() {
	fmt.Println("Usage:")

	fmt.Printf("\tcreatewallet  -- 创建钱包\n")
	fmt.Printf("\taccounts  -- 获取账单列表\n")
	//创建区块链
	fmt.Printf("\tcreateblockchain  -address ADDRESS --创建区块链\n")
	//添加区块
	fmt.Printf("\taddblcok -data DATA  --添加区块\n")
	//打印完整的区块信息
	fmt.Printf("\tprintchain  --打印完整区块信息\n")
	//通过命令行转账
	fmt.Printf("\tsend  -from FROM -to TO -amount AMOUNT -- 发起转账\n")

	fmt.Printf("\t转账参数说明\n")

	fmt.Printf("\t\t-from FROM --转账原地址\n")
	fmt.Printf("\t\t-to TO --转账目标地址\n")
	fmt.Printf("\t\t-amoount AMOUNT 转账金额\n")

	//查询余额
	fmt.Printf("\tgetbalance -address FROM --查询指定地址的余额\n")
	fmt.Println("\t查询余额参数说明")
	fmt.Printf("\t\t--address --查询的余额的地址")

	fmt.Printf("\tutxo -test METHOD --测试UTXO Table功能中指定的方法\n")
	fmt.Printf("\t\tMETHOD --方法名\n")
	fmt.Printf("\t\t\treset --重置UTXOTable\n")
	fmt.Printf("\t\t\tbalance --查找所有UTXO\n")

	//

}

//添加区块
func (cli *CLI) addBlock(txs []*Transaction) {

	blockchain := BlockChainObject()
	blockchain.AddBlock(txs)
	defer blockchain.DB.Close()
	//获取到blockchain的对象实力
}

//运行命令行函数
func (cli *CLI) Run() {
	IsValidArgs()
	//新建相关命令
	//创建钱包
	createWalletsCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	//获取账单列表
	getAccountsCmd := flag.NewFlagSet("accounts", flag.ExitOnError)
	//添加区块
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	//输出区块链完整信息
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	//创建区块的时候指定旷工的地址接收奖励
	createBLCWithGensisBlockCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)

	//发起交易
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)

	//查询余额的命令
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)

	//utxo测试命令
	UTXOTestCmd := flag.NewFlagSet("utxo", flag.ExitOnError)

	//数据参数处理
	//添加区块
	flagAddBlockArg := addBlockCmd.String("data", "send 100 btc to player", "添加区块数据")

	//创建区块链的时候指定的旷工的地址
	flagCreateBlockchainArg := createBLCWithGensisBlockCmd.String("address", "troytan", "指定系统奖励的旷工地址")

	//发起交易的命令行参数
	flagSendFromArg := sendCmd.String("from", "", "转账源地址")
	flagSendToArg := sendCmd.String("to", "", "目标地址")
	flagSendAmountArg := sendCmd.String("amount", "", "转账金额")
	//查询余额的命令行参数
	flagGetBalanceArg := getBalanceCmd.String("address", "", "要查询的地址")

	//utxo测试命令行参数
	flagUTXOTestArg := UTXOTestCmd.String("method", "", "UTXO Table 相关操作")
	//判断命令
	switch os.Args[1] {
	case "utxo":
		err := UTXOTestCmd.Parse(os.Args[2:])
		if nil != err {
			log.Panicf("parse cmd of operate utxo  failed!%v\n", err)
		}

	case "accounts":
		err := getAccountsCmd.Parse(os.Args[2:])
		if nil != err {
			log.Panicf("parse cmd of get accounts  failed!%v\n", err)
		}
	case "createwallet":
		err := createWalletsCmd.Parse(os.Args[2:])
		if nil != err {
			log.Panicf("parse cmd of create wallet failed!%v\n", err)
		}
	case "getbalance":
		if err := getBalanceCmd.Parse(os.Args[2:]); nil != err {
			log.Panicf("parse getbalance failed!%v\n", err)
		}

	case "send":
		if err := sendCmd.Parse(os.Args[2:]); nil != err {
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
	//utxo table 操作
	if UTXOTestCmd.Parsed() {
		switch *flagUTXOTestArg{
		case "balance":
			cli.TestFindUTXOMap()
		case "reset":
			cli.TestResetUTXO()
		}
	}

	//获取账单
	if getAccountsCmd.Parsed() {
		cli.GetAccounts()
	}
	//发起转账
	if sendCmd.Parsed() {
		if *flagSendFromArg == "" {
			fmt.Println("源地址不能为空")
		}
		if *flagSendToArg == "" {
			fmt.Println("目标地址不能为空")
		}
		if *flagSendAmountArg == "" {
			fmt.Println("转账金额不能为空")
		}
		fmt.Printf("\tFROM:[%s]\n", JsonToSlice(*flagSendFromArg))
		fmt.Printf("\tTO:[%s]\n", JsonToSlice(*flagSendToArg))
		fmt.Printf("\tAMOUNT:[%s]\n", JsonToSlice(*flagSendAmountArg))
		cli.send(JsonToSlice(*flagSendFromArg), JsonToSlice(*flagSendToArg), JsonToSlice(*flagSendAmountArg))
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

		cli.CreateBlockchain(*flagCreateBlockchainArg)
	}
	//查询余额
	if getBalanceCmd.Parsed() {
		if *flagGetBalanceArg == "" {
			fmt.Println("请输入查询地址......")
			os.Exit(1)
		}
		cli.GetBalance(*flagGetBalanceArg)
	}

}
