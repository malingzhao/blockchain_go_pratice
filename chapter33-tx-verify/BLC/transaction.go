package BLC

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
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
func NewCoinbaseTransaction(address string) *Transaction {
	var txCoinbase *Transaction
	//输入
	//coinbase 特点
	//txHash:nil  交易哈希
	//vout: 引用的上一笔交易的索引
	//ScriptSig：用户名
	txInput := &TxInput{[]byte{}, -1, nil, nil}
	//输出
	//value
	//address
	//txOutput := &TxOutput{10, StringToHash160(address)}
	txOutput := NewTxOutput(10, address)
	//输入输出组装交易
	txCoinbase = &Transaction{nil, []*TxInput{txInput}, []*TxOutput{txOutput}}
	txCoinbase.HashTransaction()
	return txCoinbase
}

//生成交易哈希(交易的序列化)
func (tx *Transaction) HashTransaction() {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	if err := encoder.Encode(tx); err != nil {
		log.Panicf("tx Hash encoded failed %v\n", err)
	}
	//生成哈希值
	hash := sha256.Sum256(result.Bytes())
	tx.TxHash = hash[:]
}

//生成普通转账交易
func NewSimpleTransaction(from string, to string, amount int, bc *BlockChain, txs []*Transaction) *Transaction {
	//输入列表
	var txInputs []*TxInput
	//输出列表
	var txOutputs []*TxOutput

	//调用utxo可花费函数
	money, spendableUTXODic := bc.FindSpendableUTXo(from, amount, txs)
	fmt.Printf("money:%v\n", money)

	//获取钱包集合对象
	wallets := NewWallets()
	//查找对应的钱包结构
	wallet := wallets.Wallets[from]

	for txHash, indexArray := range spendableUTXODic {
		txHashBytes, err := hex.DecodeString(txHash)
		if nil != err {
			log.Panicf("decide string to byte[] failed:%v\n", err)
		}
		//遍历索引列表
		for _, index := range indexArray {
			txInput := &TxInput{txHashBytes, index, nil, wallet.PublicKey}
			txInputs = append(txInputs, txInput)
		}
	}

	////输入
	//txInput := &TxInput{[]byte("f449245013e0d88c713617511436ef2f546ff3" +
	//	"c600a44976c0f0a651f6d722f6"), 0, from}
	//txInputs = append(txInputs, txInput)
	//输出
	//txOutput := &TxOutput{amount, to}
	txOutput := NewTxOutput(amount, to)

	txOutputs = append(txOutputs, txOutput)
	//输出(找零)
	if amount < money {
		//找零
		//txOutput = &TxOutput{money - amount, from}
		txOutput = NewTxOutput(money-amount, from)
		txOutputs = append(txOutputs, txOutput)
	} else {
		log.Panicf("余额不足")
	}
	tx := Transaction{nil, txInputs, txOutputs}

	tx.HashTransaction() //生成一笔新的交易
	// 对交易进行签名
	bc.SignTransaction(&tx, wallet.PrivateKey, nil)
	return &tx

}

//判断指定的交易是否是一个Coinbase的交易
func (tx *Transaction) IsCoinBaseTransaction() bool {

	return tx.Vins[0].Vout == -1 && len(tx.Vins[0].TxHash) == 0
}

//新建output对象
func NewTxOutput(value int, address string) *TxOutput {
	txOutput := &TxOutput{}
	hash160 := StringToHash160(address)
	txOutput.Value = value
	txOutput.Ripemd160Hash = hash160
	return txOutput
}

func (tx *Transaction) Sign(privKey ecdsa.PrivateKey, prevTxs map[string]Transaction) {
	//处理输入保证交易的正确性

	//检查tx中的所有的每一个输入所引用的交易哈希是否包含在prevTxs
	//如果没包含在里面说明交易被修改了
	for _, vin := range tx.Vins {
		if prevTxs[hex.EncodeToString(vin.TxHash)].TxHash == nil {
			log.Panicf("ERROR:Prev transaction is not corrected !\n")
		}
	}

	//提取需要签名的属性
	txCopy := tx.TrimmedCopy()
	//处理交易副本的而输入
	for vin_id, vin := range txCopy.Vins {
		//获取关联交易
		prevTx := prevTxs[hex.EncodeToString(vin.TxHash)]
		//找到发送者（当前输入引用的哈希输出的哈希)
		txCopy.Vins[vin_id].PublicKey = prevTx.Vouts[vin.Vout].Ripemd160Hash
		//生成交易副本的哈希
		txCopy.TxHash = txCopy.Hash()
		//调用核心签名函
		r, s, err := ecdsa.Sign(rand.Reader, &privKey, txCopy.TxHash)
		if nil != err {
			log.Panicf("sign to transaction [%x] failed!%v\n", err)
		}
		//组成交易签名
		signature := append(r.Bytes(), s.Bytes()...)
		tx.Vins[vin_id].Signature = signature
	}

}

//交易拷贝 生成一个专门用于交易签名的副本
func (tx *Transaction) TrimmedCopy() Transaction {
	//重新组装一个生成一个新的交易
	var inputs []*TxInput
	var outputs []*TxOutput
	//组装input
	for _, vin := range tx.Vins {
		inputs = append(inputs, &TxInput{vin.TxHash, vin.Vout, nil, nil})
	}

	//组装output
	for _, vout := range tx.Vouts {
		outputs = append(outputs, &TxOutput{vout.Value, vout.Ripemd160Hash})
	}
	txCopy := Transaction{tx.TxHash, inputs, outputs}

	return txCopy
}

//设置用于签名交易的哈希
func (tx *Transaction) Hash() []byte {
	txCopy := tx
	tx.TxHash = []byte{}
	hash := sha256.Sum256(txCopy.Serialize())
	return hash[:]

}

//交易序列化
func (tx *Transaction) Serialize() []byte {
	var buffer bytes.Buffer
	//gob
	//新建编码对象
	encoder := gob.NewEncoder(&buffer)
	//编码 （序列化）
	if err := encoder.Encode(tx); nil != err {
		log.Panicf("serialize the tx to []byte failed %v\n", err)
	}
	return buffer.Bytes()
}

//验证签名
func (tx *Transaction) Verify(prevTxs map[string]Transaction) bool {
	//检查能否找到指定交易哈希
	for _, vin := range tx.Vins {
		if prevTxs[hex.EncodeToString(vin.TxHash)].TxHash == nil {
			log.Panicf("VERIFY ERROR:transaction verify failed!\n")
		}
	}
	//提取相同的交易签名属性
	txCopy := tx.TrimmedCopy()
	//使用相同的椭圆
	curve := elliptic.P256()

	//遍历tx输入,对每一笔输出进行验证
	for vinId, vin := range tx.Vins {
		prevTx := prevTxs[hex.EncodeToString(vin.TxHash)]
		//找到发送者（当前输入引用的哈希输出的哈希)
		txCopy.Vins[vinId].PublicKey = prevTx.Vouts[vin.Vout].Ripemd160Hash
		// 由需要验证的数据生成的交易哈希，必须要与签名时的数据完全一致
		txCopy.TxHash = txCopy.Hash()
		//在比特币中签名是一个数值对 r，x 代表签名
		//所以要从输入的signature中获取
		//获取 r, s, rs 长度相等
		r := big.Int{}
		s := big.Int{}
		sigLen := len(vin.Signature)
		r.SetBytes(vin.Signature[:(sigLen / 2)])
		s.SetBytes(vin.Signature[(sigLen)/2:])

		//获取公钥
		//公钥由新，y 组成
		x := big.Int{}
		y := big.Int{}

		pubKyeLen := len(vin.PublicKey)
		x.SetBytes(vin.PublicKey[:pubKyeLen/2])
		y.SetBytes(vin.PublicKey[pubKyeLen/2:])

		rawPublicKey := ecdsa.PublicKey{curve, &x, &y}
		if !ecdsa.Verify(&rawPublicKey, txCopy.TxHash, &r, &s) {
			return false
		}

	}

	//调用签名核心函数
	return     true
}
