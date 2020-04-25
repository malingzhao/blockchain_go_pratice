package BLC

/*
交易的输出管理
*/
type TxOutput struct {
	//金额
	value int
	//用户名(UTXO的所有者)
	ScriptPubKey string
}
