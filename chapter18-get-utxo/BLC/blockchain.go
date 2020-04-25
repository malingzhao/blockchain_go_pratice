package BLC

import (
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
	"strconv"
)

//区块链的管理文件
//数据库名称
//表名称
const dbName = "block.db"
const blockTableName = "blocks"

//区块链的基本结构
type BlockChain struct {
	//Blocks []*Block //区块的切片
	DB  *bolt.DB //数据库对象
	Tip []byte   //保存最新区块的哈希值
}

//初始化区块链
func CreateBlockChainWithGensisBlcok(address string) *BlockChain {

	//文件已存在，说明创世区块已存在
	if dbExists() {
		fmt.Println("创世区块已存在")
		os.Exit(1)
	}
	//保存最新区块的哈希值
	var blockHash []byte
	//1. 创建或者打开一个数据库

	//3. 把创世区块传入数据库中
	// w  r x
	//4 2 1 读 写
	db, err := bolt.Open(dbName, 0600, nil)
	//fmt.Printf("dbName:%s",dbName)
	if err != nil {
		log.Panicf("create db [%s] failed %v\n", dbName, err)
	}
	//2. 创建桶
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b == nil {
			//代表没找到
			b, err := tx.CreateBucket([]byte(blockTableName))

			if err != nil {
				log.Panicf("create bucket [%s] failed %v\n", b, err)
			}
			//生成一个coinbase交易
			txCoinbase := NewCoinbaseTransaction(address)
			//生成创世区块
			gensisBlock := CreateGensisBlock([]*Transaction{txCoinbase})

			//存储
			//1. key value 分别以什么数据代表- hash
			//2. 如何把block结构存储到数据库中
			//3. josn
			err = b.Put(gensisBlock.Hash, gensisBlock.Serialize())

			if err != nil {
				log.Panicf("Insert the gensis block fauled %v\n", err)
			}
			blockHash = gensisBlock.Hash

			err = b.Put([]byte("l"), gensisBlock.Hash)
			if err != nil {
				log.Panicf("save the lastest hash of gensisblock failed %v\n", err)
			}

		}
		//存储最新区块的还行
		//l ---- >> lastest
		return nil
	})
	//defer db.Close()
	return &BlockChain{DB: db, Tip: blockHash}
}

//添加区块到区块链中
func (bc *BlockChain) AddBlock(txs []*Transaction) {

	//更新区块数据(insert)
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		//1. 获取数据库的桶
		b := tx.Bucket([]byte(blockTableName))
		if nil != b {
			fmt.Printf("lastest hash:%v\n", b.Get([]byte("l")))
			//2.获取 最后插入的区块
			blockByte := b.Get(bc.Tip)
			//区块数据的反序列化
			lastest_block := DeSerializeBlock(blockByte)
			//3. 新建区块
			newBlock := NewBlock(lastest_block.Height+1, lastest_block.Hash, txs)
			//4. 存入数据里
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if nil != err {
				log.Panicf("insert the new block to db failed %v", err)
			}
			//更新最新区块的哈希(数据库）
			err = b.Put([]byte("l"), newBlock.Hash)
			if nil != err {
				log.Panicf("update the lastest block to db failed %v", err)
			}
			//更新区块链对象中的最新区块哈希
			bc.Tip = newBlock.Hash
		}
		return nil
	})

	if err != nil {
		log.Panicf("blockchain add block error")
	}
}

//初始化区块
//生成创世区块
func CreateGensisBlock(txs []*Transaction) *Block {
	return NewBlock(1, nil, txs)
}

//遍历数据库 输出所有的区块信息
func (bc *BlockChain) PrintChain() {

	var currentBlock *Block
	//读取数据库
	fmt.Println("打印区块链完整信息.......")
	//获取迭代对象
	bcit := bc.Iterator()

	//循环读取
	//什么时候退出
	for {
		fmt.Println("-----------------------------------")

		currentBlock = bcit.Next()
		//输出区块详情
		fmt.Printf("\tHash:%x\n", currentBlock.Hash)
		fmt.Printf("\tPrevBlockHash:%x\n", currentBlock.PrevBlockHash)
		fmt.Printf("\tTimeStamp:%v\n", currentBlock.TimeStamp)
		fmt.Printf("\tHeight:%d\n", currentBlock.Height)
		fmt.Printf("\tNonce:%d\n", currentBlock.Nonce)

		fmt.Printf("\tTxs:%v\n", currentBlock.TXs)

		for _, tx := range currentBlock.TXs {
			fmt.Printf("\t\t tx-hash:%x\n", tx.TxHash)
			fmt.Printf("\t\t输入....\n")
			for _, vin := range tx.Vins {
				fmt.Printf("\t\tvin-TxHash:%x\n", vin.TxHash)
				fmt.Printf("\t\tvin-vout:%v\n", vin.Vout)
				fmt.Printf("\t\tvin-scriptSig:%s\n", vin.ScriptSig)
			}
			fmt.Printf("\t\t输出...\n")

			for _, vout := range tx.Vouts {
				fmt.Printf("\t\tvout-value：%d\n", vout.Value)
				fmt.Printf("\t\tvout-scriptPubKey:%v\n", vout.ScriptPubKey)
			}
		}
		//退出条件
		//转换为big.Int
		var hashInt big.Int
		hashInt.SetBytes(currentBlock.PrevBlockHash)
		//比较
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			//遍历到创世区块
			break
		}

	}
}

//判断数据库文件是否存在
func dbExists() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		//数据库文件不存在
		return false
	}
	return true

}

//实现挖矿的功能
//通过接收交易生成区块

func (blockchain *BlockChain) MineNewBlock(from, to, amount []string) {
	//搁置交易生成步骤
	var txs []*Transaction
	var block *Block
	value, _ := strconv.Atoi(amount[0])
	//生成新的交易
	tx := NewSimpleTransaction(from[0], to[0], value)
	//追加新的交易到交易列表中去
	txs = append(txs, tx)

	//从数据库中获取最新的一个区块
	blockchain.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte([]byte(blockTableName)))
		if nil != b {
			//获取最新的区块的哈希值
			hash := b.Get([]byte("l"))
			//获取最新区块
			blockBytes := b.Get(hash)
			//反序列化
			block = DeSerializeBlock(blockBytes)
		}
		return nil
	})
	//通过数据库中最新的区块生成新的区块
	block = NewBlock(block.Height+1, block.Hash, txs)
	//持久化新生成的区块到数据库中
	blockchain.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if nil != b {
			err := b.Put(block.Hash, block.Serialize())
			if err != nil {
				log.Panicf("update the new block to db failed %v\n", err)
			}
			//更新最新区块的哈希值
			err = b.Put([]byte("l"), block.Hash)
			if err != nil {
				log.Panicf("update the lastest  block to db failed %v\n", err)
			}
			blockchain.Tip = block.Hash
		}
		return nil
	})

}

//获取指定地址所有的已经花费的输出 获取索引段额列表  string 存取的是用户名  int存放索引
func (blockchain *BlockChain) SpentOutputs(address string) map[string][]int {
	//已花费输出的缓存
	spentTXOutputs := make(map[string][]int)
	//获取迭代器对象
	bcit := blockchain.Iterator()
	for {
		block := bcit.Next()
		for _, tx := range block.TXs {
			//排除coinbase交易
			if !tx.IsCoinBaseTransaction(){

				for _, in := range tx.Vins {
					if in.CheckPubKeyWithAddress(address) {
						key := hex.EncodeToString(in.TxHash)
						//添加到已花费输出的花村中 如果在vin中出现了vout的话那么说明交易已经被引用了
						spentTXOutputs[key] = append(spentTXOutputs[key], in.Vout)
					}
				}
			}

		}
		//退出循环条件
		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)
		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break
		}
	}
	return nil
}

//查找指定地址的UTXO
/*
遍历查找区块链数据库中的每一个区块中的每一个交易
查找每一个交易中的每一个输出
哦安短每个输出是否满足以下条件：
 1. 属于传入的地址
 2. 是否被花费
  1. 首先遍历一次区块链数据库 将所有已花费的UTXO存入一个缓存
  2. 再次遍历区块链数据库，检查每一个UTXO是否包含在前面的已经花费输出额缓存中
*/
func (blockchain *BlockChain) UnUTXOS(address string) []*TxOutput {

	//1. 遍历数据库 查找所有的与address相关的交易
	//2. 获取迭代器
	bcit := blockchain.Iterator()
	//当前地址未花费的输出列表
	var unUTXOS []*TxOutput

	spentTXOuputs := blockchain.SpentOutputs(address)

	//获取指定地址所有的已经花费的输出
	for {
		//迭代 不断的获取下一个区块
		block := bcit.Next()
		//遍历区块中的每一笔交易
		for _, tx := range block.TXs {
		//跳转
		work:
			for index, vout := range tx.Vouts {
				//index 当前输出在当前交易的索引位置
				if vout.CheckPubKeyWithAddress(address) {
                     var isSpentOutput bool //默认false

					//当前的vout属于传入地址
					if len(spentTXOuputs) != 0 {
						for txHash, indexArray := range spentTXOuputs {

							for _, i := range indexArray {
								//txHash : 当前输出所引用的交易哈希
								//indexArray 哈希关联的vout的索引列表
								if txHash == hex.EncodeToString(tx.TxHash) && index == i {
									//txHash == hex.EncodeToString(tx.TxHash)
									//说明当前交易tx至少已经有输出被其他交易的输入引用
									//index==i 说明当期那 的输出被其他交易引用
									//跳转到最外层的循环 判断下一个vout
									isSpentOutput=true
									continue work
								} else {
									unUTXOS = append(unUTXOS, vout)
								}
							}

						}
						if isSpentOutput==false {
							unUTXOS = append(unUTXOS,vout)
						}

					} else {
						//将当前所有输出都添加到未花费的输出中

						unUTXOS = append(unUTXOS, vout)
					}
				}
				//vout 当前输出
			}

		}
		//退出循环条件
		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)
		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break
		}
	}

	return unUTXOS
}

//查询余额
func (blockchain *BlockChain) getBlance(address string) int {
	var amount int //余额
	utxos := blockchain.UnUTXOS(address)

	for _,utxo :=range utxos {
		amount +=utxo.Value
	}
	return  amount
}

