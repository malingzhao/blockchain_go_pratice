package BLC

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
	ScriptSig string
}



//验证引用的地址是否匹配
func (tx *TxInput) CheckPubKeyWithAddress(address string) bool {
	return address == tx.ScriptSig
}