package BLC

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
)

//钱包和管理想关文件

type Wallet struct {
	//1. 私钥
	PrivateKey ecdsa.PrivateKey
	//2. 公钥  验证交易的时候最简单的方法验证 写签名的时候会发现
	PublicKey []byte
}

//创建一个钱包
func NewWallet() *Wallet {
	privateKye, publicKey := newKeyPair()
	//公钥 私钥 赋值\
	return &Wallet{PrivateKey: privateKye, PublicKey: publicKey}
}
//通过钱包生成公钥-私钥对
func newKeyPair() (ecdsa.PrivateKey, []byte) {
	//1.  获取一个椭圆
	curve := elliptic.P256()
	//2. 通过椭圆生成私钥
	priv, err := ecdsa.GenerateKey(curve, rand.Reader)
	if nil != err {
		log.Panicf("ecdsa generate private key failed:%v\n", err)
	}
	//3. 通过私钥生成公钥

	pubKey :=append(priv.PublicKey.X.Bytes(),priv.PublicKey.Y.Bytes()...)
	return *priv, pubKey
}
