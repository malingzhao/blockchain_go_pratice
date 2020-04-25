package BLC

type INV struct {
	AddrFrom string //当前节点地址
	Hashes [][]byte //当前展示节点上的所有的区块  数据量是非常的大的 //当前节点的所有的区块的哈希列表
}

