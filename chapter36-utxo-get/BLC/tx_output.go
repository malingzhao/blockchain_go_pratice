package BLC

import "bytes"

/*
交易的输出管理
*/
type TxOutput struct {
	//金额(大写才能导出金额)
	Value int
	//ScriptPubKey string
	//用户名(UTXO的所有者)
	Ripemd160Hash []byte
}

////验证当前UTXO是否属于指定的第地址  查找属于自己的UTXO
//func (txOutput *TxOutput) CheckPubKeyWithAddress(address string) bool {
//	return address == txOutput.ScriptPubKey
//}

//output身份验证
func (TxOutput *TxOutput) UnLockScriptPubKeyWithAddress(address string) bool {
	//转换
	hash160 := StringToHash160(address)
	return bytes.Compare(hash160, TxOutput.Ripemd160Hash) == 0
}
