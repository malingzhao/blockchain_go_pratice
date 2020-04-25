package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
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
func CreateBlockChainWithGensisBlcok() *BlockChain {

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
			//生成创世区块
			gensisBlock := CreateGensisBlock([]byte("init blockchain"))



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
func (bc *BlockChain) AddBlock(data []byte) {

	//更新区块数据(insert)
	err:=bc.DB.Update(func(tx *bolt.Tx) error {
		//1. 获取数据库的桶
		b := tx.Bucket([]byte(blockTableName))
		if nil != b {
			fmt.Printf("lastest hash:%v\n",b.Get([]byte("l")))
			//2.获取 最后插入的区块
			blockByte := b.Get(bc.Tip)
			//区块数据的反序列化
			lastest_block := DeSerializeBlock(blockByte)
			//3. 新建区块
			newBlock := NewBlock(lastest_block.Height+1, lastest_block.Hash, data)
			//4. 存入数据里
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if nil != err {
				log.Panicf("insert the new block to db failed %v", err)
			}
			//更新最新区块的哈希(数据库）
			err=b.Put([]byte("l"), newBlock.Hash)
			if nil != err {
				log.Panicf("update the lastest block to db failed %v", err)
			}
			//更新区块链对象中的最新区块哈希
			bc.Tip = newBlock.Hash
		}
		return nil
	})

	if err!=nil{
		log.Panicf("blockchain add block error")
	}
}

//初始化区块
//生成创世区块
func CreateGensisBlock(data []byte) *Block {
	return NewBlock(1, nil, data)
}
