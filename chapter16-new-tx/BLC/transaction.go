package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

/*
交易管理文件
*/
//定义一个基本结构
type Transaction struct {
	//交易哈希(表示)
	TxHash []byte
	//输入列表
	Vins []*TxInput
	//输出列表
	Vouts []*TxOutput
}
//实现coinbase交易
func NewCoinbaseTransaction(address string ) *Transaction {
	var txCoinbase *Transaction
	//输入
	//coinbase 特点
	//txHash:nil  交易哈希
	//vout: 引用的上一笔交易的索引
	//ScriptSig：用户名
	txInput := &TxInput{[]byte{}, -1, "syetm reward"}
	//输出
	//value
	//address
	txOutput := &TxOutput{10, address}
	//输入输出组装交易
	txCoinbase = &Transaction{nil, []*TxInput{txInput}, []*TxOutput{txOutput}}
	txCoinbase.HashTransaction()
	return txCoinbase
}

//生成交易哈希(交易的序列化)
func (tx *Transaction) HashTransaction(){
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	if err :=encoder.Encode(tx);err!=nil{
		log.Panicf("tx Hash encoded failed %v\n",err)
	}
	//生成哈希值
	hash:=sha256.Sum256(result.Bytes())
	tx.TxHash = hash[:]
}
//生成普通转账交易
func NewSimpleTransaction(from string, to string , amount int) *Transaction{
   //输入列表
	var txInputs []*TxInput
	//输出列表
	var txOutputs []*TxOutput


	//输入
	txInput :=&TxInput{}
	txInputs = append(txInputs,txInput)
	//输出
	txOutput :=&TxOutput{amount, to}

	txOutputs = append(txOutputs, txOutput)
	//输出(找零)
	if amount < 10{
		//找零
		txOutput = &TxOutput{10-amount, from}
		txOutputs = append(txOutputs,txOutput)
	}
	tx :=Transaction{[]byte("f449245013e0d88c713617511436ef2f546ff3" +
		"c600a44976c0f0a651f6d722f6"),txInputs,txOutputs}

	tx.HashTransaction()
	return &tx

}

