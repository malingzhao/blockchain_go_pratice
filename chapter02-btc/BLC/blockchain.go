package BLC

//区块链的管理文件

//区块链的基本结构
type BlockChain struct {
	Blocks []*Block //区块的切片
}


//初始化区块链
func CreateBlockChainWithGensisBlcok() *BlockChain {
	//生成创世区块
	block := CreateGensisBlock([]byte("init blockchain"))
	return &BlockChain{[]*Block{block}}
}

//添加区块到区块链中
func (bc *BlockChain) AddBlock(height int64, prevBlcokHash []byte, data []byte) {
	//var newBlock *Block
	newBlock := NewBlock(height, prevBlcokHash, data)
	bc.Blocks = append(bc.Blocks, newBlock)
}

//初始化区块
//生成创世区块
func CreateGensisBlock(data []byte) *Block{
	return  NewBlock(1, nil,data)
}