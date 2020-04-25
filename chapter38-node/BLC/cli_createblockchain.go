package BLC

//初始化区块链
func (cli *CLI) CreateBlockchain(address string,nodeId string) {
	//创建区块链对象
	bc := CreateBlockChainWithGensisBlcok(address,nodeId)
	defer bc.DB.Close()


	utxoSet := &UTXOSet{bc}
	//utxo的重置
	utxoSet.ResetUTXOSet()
}
