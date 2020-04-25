package BLC

import (
	"bytes"
	"crypto/sha256"
	"time"
)

//区块基本结构与功能管理文件

//实现一个最基本的区块结构
type Block struct {
	TimeStamp     int64  //区块时间 代表区块时间
	Hash          []byte // 当前区块哈希
	PrevBlockHash []byte //前区块哈希
	Height        int64  //区块高度 识别区块在区块链中的位置
	Data          []byte //交易数据
	Nonce         int64 // 运行pow的时候生成的哈希变化值，也是代表pow运行的时候动态生成的数据
}

//新建区块
func NewBlock(height int64, prevBlockHash []byte, data []byte) *Block {
	var block Block
	block = Block{
		TimeStamp:     time.Now().Unix(),
		PrevBlockHash: prevBlockHash,
		Height:        height,
		Data:          data,}

	//不需要返回值
	block.SetHash()
	//替换setHash
	//通过Pow生成新的哈希
	pow := NewProofOfWork(&block)
	hash,nonce:=pow.Run()
	block.Hash = hash
	block.Nonce = int64(nonce)
	//执行工作量证明算法
	//生成哈希
	return &block
}

//func method  什么时候使用func  什么时候我们使用method
//为了不传参不写返回值所以
//计算区块哈希 代码的可读性  先考且屡设置代码的函数
func (b *Block) SetHash() {
	//调用sha256 实现哈希的生成
	//实现int -> hash
	timeStampBytes := IntToHex(b.TimeStamp)
	heightBytes := IntToHex(b.Height)
	blockBytes := bytes.Join([][]byte{
		heightBytes,
		timeStampBytes,
		b.PrevBlockHash,
		b.Data,
	}, []byte{})
	hash := sha256.Sum256(blockBytes)
	b.Hash = hash[:]

}
