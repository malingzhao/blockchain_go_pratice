package BLC

/*
交易的输出管理
*/
type TxOutput struct {
	//金额(大写才能导出金额)
	Value int
	//用户名(UTXO的所有者)
	ScriptPubKey string
}




//验证当前UTXO是否属于指定的第地址  查找属于自己的UTXO
func(txOutput *TxOutput) CheckPubKeyWithAddress(address string) bool{
	return address == txOutput.ScriptPubKey
}




