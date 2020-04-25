package BLC

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

//utxo 持久化管理
// 用于存放utxo的bucket

const utxoTableName = "utxoTable"

//utxo Set 结构 保存指定区块链总所有的utxo
type UTXOSet struct {
	Blockchain *BlockChain
}

//更新

//查找

//重置

//重置
func (utxoSet *UTXOSet) ResetUTXOSet() {
	//在第一次创建的时候就更新UTXO table
	utxoSet.Blockchain.DB.Update(func(tx *bolt.Tx) error {
		//查找utxo table

		fmt.Printf("%s\n", utxoTableName)
		b := tx.Bucket([]byte(utxoTableName))
		//如果查找不到
		if nil != b {
			//删除这个桶
			err := tx.DeleteBucket([]byte(utxoTableName))
			if nil != err {
				log.Panicf("delete the utxo table failed!%v\n", err)
			}
		}
		//创建
		bucket, err := tx.CreateBucket([]byte(utxoTableName))
		fmt.Printf("%v\n",  bucket)
		if nil != err {
			log.Panicf("create bucket failed!%v\n", err)
		}
		if nil != bucket {
			//查找当前所有UTXO
			txOutputMap := utxoSet.Blockchain.FindUTXOMap()
			fmt.Printf("%v\n", txOutputMap)
			for keyHash, outputs := range txOutputMap {
				//将所有的utxo存入
				txHash, _ := hex.DecodeString(keyHash)
				fmt.Printf("TxHash:%v\n", txHash)
				//存入 utxo table
				err = bucket.Put(txHash, outputs.Serialize())
				if nil!=err {
					log.Panicf("put the utxo into table failed!%v\n", err)
				}
			}
		}
		return  nil
	})

}



func  (txOutputs  *TxOutputs) Serialize() []byte{
	var result bytes.Buffer
	encoder:= gob.NewEncoder(&result)
		if err := encoder.Encode(txOutputs); nil!=err {
			log.Panicf("serialize the utxo failed")
		}

	return result.Bytes()
}