package BLC

import "crypto/sha256"

/*
树里面存储的就是节点
*/
//Merkle树实现管理
type MerkleTree struct {
	//根节点
	RootNode *MerkleNode
}

//merkle的节点的结构
type MerkleNode struct {
	//左子节点
	Left *MerkleNode
	//右子节点
	Right *MerkleNode
	Data  []byte
}

//创建MerkleNode 节点
func MakeMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	node := &MerkleNode{}
	//判断叶子节点
	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		node.Data = hash[:]
	} else {
		//非叶子节点
		prevHashes := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHashes)
		node.Data = hash[:]
	}

	//子节点的赋值的操作
	node.Left = left
	node.Right = right
	return node
}

//创建Merkle树
//txHashes 区块中的交易哈希列表
//Merkle节点根节点之外的其他层次的子节点的数量必须是偶数个 如果是奇数个 则将最后一个节点复制一份
func NewMerkleTree(txHashes [][]byte) *MerkleTree {
	//节点的列表 存储每一笔交易的哈希
	var nodes []MerkleNode
	//判断交易数据的条数 如果是奇数条的话 拷贝最后一份
	if len(txHashes)%2 != 0 {
		txHashes = append(txHashes, txHashes[len(txHashes)-1])
	}

	//遍历所有交易数据， 通过哈希生成叶子节点
	for _, data := range txHashes {
		node := MakeMerkleNode(nil, nil, data)
		nodes = append(nodes, *node)
	}

	/*
	假设6笔交易  len(txHashes) = 6
	i=0
	i=1
	=2
	 */

	//通过叶子节点创建父节点
	for i := 0; i < len(txHashes)/2; i++ {
		var parentNodes []MerkleNode // 父节点列表
		for j := 0; j < len(nodes); j += 2 {
			node := MakeMerkleNode(&nodes[j], &nodes[j+1], nil)
			parentNodes = append(parentNodes, *node)
		}
		if len(parentNodes)%2 != 0 {
			parentNodes = append(parentNodes, parentNodes[len(parentNodes)-1])
		}
		//最终只保存了根节点的哈希值
		nodes = parentNodes
	}


	mtree :=MerkleTree{&nodes[0]}
	return  &mtree

}
