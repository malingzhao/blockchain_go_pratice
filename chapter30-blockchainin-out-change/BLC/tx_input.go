package BLC

import "bytes"

/*
交易的输入管理
*/

//输入结构
type TxInput struct {
	//交易哈希(不是指当前的交易的哈希)
	TxHash []byte
	// 引用的上一笔交易的索引号
	Vout int
	//用户名
	//ScriptSig string
	//数字签名
	Signature string
	//公钥
	PublicKey []byte
}



////验证引用的地址是否匹配
//func (tx *TxInput) CheckPubKeyWithAddress(address string) bool {
//	return address == tx.ScriptSig
//}

//传递hash160 进行判断
func (in *TxInput) UnLoakRipemd160Hash(ripremd160Hash []byte) bool {
	//截取input的ripemd160hash
	inputRipemd160Hash :=Ripemd160Hash(in.PublicKey)

	return bytes.Compare(inputRipemd160Hash,ripremd160Hash)==0
}