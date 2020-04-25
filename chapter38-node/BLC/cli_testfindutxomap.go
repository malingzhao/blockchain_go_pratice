package BLC

//重置utxo table
func (cli *CLI) TestResetUTXO(nodeId string) {
	//获取对象
	blockchain := BlockChainObject(nodeId)
	defer blockchain.DB.Close()

	utxoSet := UTXOSet{Blockchain: blockchain}
	utxoSet.ResetUTXOSet()
}

//重置


//查找

func (cli *CLI) TestFindUTXOMap(){

}