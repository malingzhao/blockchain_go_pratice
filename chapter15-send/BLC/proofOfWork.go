package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

/*
共识管理文件
实现Pow以及相关功能
*/

//实现目标难度
const targetBit = 16

//工作量证明结构
type ProofOfWork struct {
	//需要共识验证的区块
	Block *Block
	// 目标难度的哈希(大数据存储)
	target *big.Int
}

//创建一个Pow的对象
func NewProofOfWork(block *Block) *ProofOfWork {
	//如何计算target
	target := big.NewInt(1)

	//数据总长度为8位
	//需求： 需要满足前两位为0 才能解决问题
	//1 * 2 << (8 - 2) = 64
	//0100 0000
	//0011 1111 = 63
	//32  * 8
	target = target.Lsh(target, 256-targetBit)

	return &ProofOfWork{Block: block, target: target}
}

//执行Pow比较哈希
//返回哈希值，以及碰撞的次数
func (proofOfWork *ProofOfWork) Run() ([]byte, int) {
	//碰撞次数
	var nonce = 0
	var hashInt big.Int
	//生成的哈希值
	var hash [32]byte
	//无限循环 生成符合条件的哈希值
	for {
		//生成准备数据
		dataBytes :=proofOfWork.prepareData(int64(nonce))
		hash= sha256.Sum256(dataBytes)
		hashInt.SetBytes(hash[:])
		//检测生成的哈希值是否符合条件
		// x<y -1 x>y 1 x=y
		//找到了符合条件的哈希值中断循环
		if proofOfWork.target.Cmp(&hashInt) == 1{
			break
		}
		nonce++
	}
	fmt.Printf("\n碰撞次数：%d\n",nonce)
	return hash[:],nonce
}

func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	var data []byte
	//拼接区块属性进行哈希计算
	timeStampBytes := IntToHex(pow.Block.TimeStamp)
	heightBytes := IntToHex(pow.Block.Height)
	data = bytes.Join([][]byte{
		heightBytes,
		timeStampBytes,
		pow.Block.PrevBlockHash,
		pow.Block.HashTransaction(),
		IntToHex(nonce),
		IntToHex(targetBit),
	}, []byte{})

	return data
}