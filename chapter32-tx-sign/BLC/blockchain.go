package BLC

import (
	"bytes"
	"crypto/ecdsa"
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
		fmt.Println("创世区块已经存在")
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

	//当前区块
	var currentBlock *Block
	//读取数据库
	fmt.Println("打印区块链完整信息.......")
	//获取迭代对象
	bcit := bc.Iterator()

	//循环读取
	//什么时候退出
	for {
		fmt.Println("-----------------------------------")

		//遍历下一个区块
		currentBlock = bcit.Next()
		//输出区块详情
		fmt.Printf("\tHash:%x\n", currentBlock.Hash)
		fmt.Printf("\tPrevBlockHash:%x\n", currentBlock.PrevBlockHash)
		fmt.Printf("\tTimeStamp:%v\n", currentBlock.TimeStamp)
		fmt.Printf("\tHeight:%d\n", currentBlock.Height)
		fmt.Printf("\tNonce:%d\n", currentBlock.Nonce)

		fmt.Printf("\tTxs:%v\n", currentBlock.TXs)

		//输出交易的详细信息
		for _, tx := range currentBlock.TXs {
			fmt.Printf("\t\t tx-hash:%x\n", tx.TxHash)
			fmt.Printf("\t\t输入....\n")
			for _, vin := range tx.Vins {
				fmt.Printf("\t\tvin-TxHash:%x\n", vin.TxHash)
				fmt.Printf("\t\tvin-vout:%v\n", vin.Vout)
				fmt.Printf("\t\tvin-scriptSig:%x\n", vin.PublicKey)
				fmt.Printf("\t\tvin-scriptSig:%x\n", vin.Signature)
			}
			fmt.Printf("\t\t输出...\n")

			for _, vout := range tx.Vouts {
				fmt.Printf("\t\tvout-value：%d\n", vout.Value)
				fmt.Printf("\t\tvout-scriptPubKey:%s\n", hex.EncodeToString(vout.Ripemd160Hash))
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
//从from 转移金额 amount 到 to
func (blockchain *BlockChain) MineNewBlock(from, to, amount []string) {
	//搁置交易生成步骤
	var txs []*Transaction
	var block *Block
	//遍历交易的参与者

	for index, address := range from {
		value, _ := strconv.Atoi(amount[index])
		//生成新的交易
		tx := NewSimpleTransaction(address, to[index], value, blockchain, txs)
		//追加新的交易到交易列表中去
		txs = append(txs, tx)
	}

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

//获取指定地址所有的已经花费的输出 获取索引段列表
func (blockchain *BlockChain) SpentOutputs(address string) map[string][]int {
	//已花费输出的缓存
	spentTXOutputs := make(map[string][]int)
	//获取迭代器对象
	bcit := blockchain.Iterator()
	for {
		//遍历每一个区块
		block := bcit.Next()
		//遍历每一个区块中的交易
		for _, tx := range block.TXs {
			//排除coinbase交易coinbase交易没有vins
			if !tx.IsCoinBaseTransaction() {
				//遍历每一个输入列表
				for _, in := range tx.Vins {

					//判断用户
					if in.UnLoakRipemd160Hash(StringToHash160(address)) {
						//保存交易哈希
						key := hex.EncodeToString(in.TxHash)
						//添加索引到已经花费列表中  找到了
						spentTXOutputs[key] = append(spentTXOutputs[key], in.Vout)
					}
					//if in.CheckPubKeyWithAddress(address) {
					//	//保存交易哈希
					//	key := hex.EncodeToString(in.TxHash)
					//	//添加索引到已经花费列表中  找到了
					//	spentTXOutputs[key] = append(spentTXOutputs[key], in.Vout)
					//}
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
	return spentTXOutputs
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
func (blockchain *BlockChain) UnUTXOS(address string, txs []*Transaction) []*UTXO {

	//1. 遍历数据库 查找所有的与address相关的交易
	//2. 获取迭代器
	bcit := blockchain.Iterator()
	//当前地址未花费的输出列表
	var unUTXOS []*UTXO

	spentTXOuputs := blockchain.SpentOutputs(address)

	//获取指定地址所有的已经花费的输出
	for _, tx := range txs {
		//判断coinbase
		if !tx.IsCoinBaseTransaction() {
			for _, in := range tx.Vins {

				//判断用户
				if in.UnLoakRipemd160Hash(StringToHash160(address)) {
					//添加已花费输出的map中
					key := hex.EncodeToString(in.TxHash)
					spentTXOuputs[key] = append(spentTXOuputs[key], in.Vout)

				}
				////判断用户
				//if in.CheckPubKeyWithAddress(address) {
				//	//添加已花费输出的map中
				//	key := hex.EncodeToString(in.TxHash)
				//	spentTXOuputs[key] = append(spentTXOuputs[key], in.Vout)
				//}
			}
		}

	}
	//遍历缓存中的UTXO
	for _, tx := range txs {
	WorkCacheTX:
		//添加一个缓存跳转
		for index, vout := range tx.Vouts {
			if vout.UnLockScriptPubKeyWithAddress(address) {
				//if vout.CheckPubKeyWithAddress(address) {
				if len(spentTXOuputs) != 0 {
					//判断交易是否被其他交易引用
					var isUtxoTx bool
					for txHash, indexArray := range spentTXOuputs {
						txHashStr := hex.EncodeToString(tx.TxHash)
						if txHash == txHashStr {

							//当前遍历到的交易已经有输出被其他交易的输入所引用
							isUtxoTx = true
							//添加状态变量 判断指定的output 是否被引用
							var isSpentUTXO bool
							for _, voutIndex := range indexArray {
								if index == voutIndex {
									//该输出被引用了
									isSpentUTXO = true
									//直接退出当前的vout的判断逻辑进行下一个判断
									continue WorkCacheTX
								}
							}
							if isSpentUTXO == false {
								utxo := &UTXO{tx.TxHash, index, vout}
								unUTXOS = append(unUTXOS, utxo)
							}
						}
					}
					if isUtxoTx == false {
						//说明当前交易中所有与address相关的outputs都是UTXO
						utxo := &UTXO{tx.TxHash, index, vout}
						unUTXOS = append(unUTXOS, utxo)
					}
				} else {
					utxo := &UTXO{tx.TxHash, index, vout}
					unUTXOS = append(unUTXOS, utxo)
				}
			}
		}

	}

	//数据库迭代

	//缓存迭代
	//查找缓存中的已花费的输出
	//优先遍历缓存中的UTXO如果余额不够直接返回， 如果不足，再遍历db文件中的UTXO
	for {
		//迭代 不断的获取下一个区块
		block := bcit.Next()
		//遍历区块中的每一笔交易
		for _, tx := range block.TXs {
			//跳转
		work:
			for index, vout := range tx.Vouts {
				//index 当前输出在当前交易的索引位置
				if vout.UnLockScriptPubKeyWithAddress(address) {
					//if vout.CheckPubKeyWithAddress(address) {
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
									isSpentOutput = true
									continue work
								}
							}

						}

						/*

							type UTXO struct {
								// UTXO  对应的交易哈希
								TxHash []byte
								//UTXO在其所属交易的输出列表中的索引
								Index int
								//Output的本身
								Output *TxOutput
							}

						*/
						if isSpentOutput == false {
							utxo := &UTXO{tx.TxHash, index, vout}
							unUTXOS = append(unUTXOS, utxo)
						}

					} else {
						//将当前所有输出都添加到未花费的输出中
						utxo := &UTXO{tx.TxHash, index, vout}
						unUTXOS = append(unUTXOS, utxo)
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
	utxos := blockchain.UnUTXOS(address, []*Transaction{})

	for _, utxo := range utxos {
		amount += utxo.Output.Value
		fmt.Printf("utxo.Output.Value：%d", utxo.Output.Value)
		fmt.Printf("amount :%d\n", amount)
	}
	return amount
}

//查找指定地址的可用的UTXO 超过amount就中断查找
//更新当前数据库指定地址的UTXO数据量
//txs 缓存中的交易列表 用于多笔交易的处理
func (blockchain *BlockChain) FindSpendableUTXo(from string, amount int, txs []*Transaction) (int, map[string][]int) {

	//可用的UTXO
	spendableUTXO := make(map[string][]int)
	var value int
	utxos := blockchain.UnUTXOS(from, txs)

	//遍历utxo
	for _, utxo := range utxos {
		//计算交易哈希
		hash := hex.EncodeToString(utxo.TxHash)
		spendableUTXO[hash] = append(spendableUTXO[hash], utxo.Index)
		value += utxo.Output.Value
		if value > amount {
			break
		}
	}

	//所有比那里完成小于amount
	//资金不足
	if value < amount {
		fmt.Printf("地址[%s]余额不足,当前余额转账金额[%d]\n", from, value, amount)
	}
	return value, spendableUTXO
}

//通过指定的交易哈希查找交易
func (blockchain *BlockChain) FindTransaction(ID []byte) Transaction {
	bcit := blockchain.Iterator()
	for {
		block := bcit.Next()
		for _, tx := range block.TXs {
			if bytes.Compare(ID, tx.TxHash) == 0 {
				//找到该交易
				return *tx
			}
		}
		//退出
		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
	}

	fmt.Printf("没有找到交易[%x]\n", ID)
	return Transaction{}

}

//交易签名
//prevTxs 代表当前交易输入引用的vout的所属交易(查找发送者)
func (blockchain *BlockChain) SignTransaction(tx *Transaction, privKey ecdsa.PrivateKey,prevTxs map[string]Transaction) {
	//coninbasea交易不需要签名
	if tx.IsCoinBaseTransaction() {
		return
	}

	// 处理交易的input，朝招tx中input所引用的vout的所属交易（查找发送者）
	//对我们所花费的每一笔UTXO进行签名的最为安全的做法
	//存储引用的交易
	prevTxs = make(map[string]Transaction)
	for _, vin := range tx.Vins {
		//查找当前交易输入所引用的交易
		tx:=blockchain.FindTransaction(vin.TxHash)
		prevTxs[hex.EncodeToString(tx.TxHash)] = tx
	}
	//签名
	tx.Sign(privKey, prevTxs)
}
