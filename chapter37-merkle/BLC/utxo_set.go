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
		fmt.Printf("%v\n", bucket)
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
				if nil != err {
					log.Panicf("put the utxo into table failed!%v\n", err)
				}
			}
		}
		return nil
	})

}

//数据集的序列化
func (txOutputs *TxOutputs) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	if err := encoder.Encode(txOutputs); nil != err {
		log.Panicf("serialize the utxo failed")
	}

	return result.Bytes()
}

//输出集反序列化
func DeserializeTxOutputs(txOutputsBytes []byte) *TxOutputs {
	var txOutputs TxOutputs
	decoder := gob.NewDecoder(bytes.NewReader(txOutputsBytes))
	if err := decoder.Decode(&txOutputs); nil != err {
		log.Panicf("deserialize the struct utxo failed:%v\n", err)
	}
	return &txOutputs
}

//通过指定地址查询
func (utxoSet *UTXOSet) FindUTXOWithAddress(address string) []*UTXO {
	var utxos []*UTXO
	err := utxoSet.Blockchain.DB.View(func(tx *bolt.Tx) error {
		//1.获取table表
		//2.
		b := tx.Bucket([]byte(utxoTableName))
		if nil != b {
			//cursor 的使用
			c := b.Cursor()
			//t通过游标遍历boltdb数据库
			for k, v := c.First(); k != nil; k, v = c.Next() {
				txOutputs := DeserializeTxOutputs(v)
				for _, utxo := range txOutputs.TxOutputs {
					if utxo.UnLockScriptPubKeyWithAddress(address) {
						utxo_single := UTXO{Output: utxo}
						fmt.Printf("uxto_single:%v\n",utxo_single)
						utxos = append(utxos, &utxo_single)
					}
				}
			}
		}
		return nil
	})
	if nil != err {
		log.Panicf("find the utxo of [%s] failed!%v\n", address, err)
	}
	return utxos
}
//查询余额
func (utxoSet *UTXOSet) GetBalance(address string) int {
	UTXOS := utxoSet.FindUTXOWithAddress(address)
	var amount int
	for _, utxo := range UTXOS {
		fmt.Printf("utxi-txhash:%x\n", utxo.TxHash)
		fmt.Printf("utxi-txhash:%x\n", utxo.Index)
		fmt.Printf("utxi-txhash:%x\n", utxo.Output.Ripemd160Hash)
		fmt.Printf("utxi-txhash:%x\n", utxo.Output.Value)
		amount += utxo.Output.Value
		fmt.Printf("amount:%v\n", amount)
	}
	return amount
}