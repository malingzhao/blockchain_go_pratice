package BLC

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"log"
)


const addressCheckSumLen = 4
//钱包和管理想关文件
//校验长度

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
	pubKey := append(priv.PublicKey.X.Bytes(), priv.PublicKey.Y.Bytes()...)
	return *priv, pubKey
}

//生成地质

// 实现双哈希
func Ripemd160Hash(pubKey []byte) []byte {
	//1. sha256
	hash256 := sha256.New()
	hash256.Write(pubKey)
	hash := hash256.Sum(nil)
	//2. ripremd160
	rmd160 := ripemd160.New()
	rmd160.Write(hash)
	rmd160.Write(hash)
	return rmd160.Sum(hash)
}

//生成校验和
func CheckSum(input []byte) []byte {
	first_hash := sha256.Sum256(input)
	second_hash := sha256.Sum256(first_hash[:])
	return second_hash[:addressCheckSumLen]
}

//通过钱包（公钥）来获取地址
func (w *Wallet) GetAddress() []byte{
	//1. 获取hsh160
	ripemd160Hash := Ripemd160Hash(w.PublicKey)
	//2. 获取校验和
	checkSumBytes := CheckSum(ripemd160Hash)
	//3.地址组成成员 拼接
	addressBytes :=append(ripemd160Hash,checkSumBytes...)
	//4. base58编码
	b58Bytes :=Base58Encode(addressBytes)


	return  b58Bytes
}

//判断地址的有效性
func IsValidForAddressBytes(addressBytes []byte) bool {
	//1. 通过base58Decode进行解码(长度为24)
	pubkeyCheckSum :=Base58Decode(addressBytes)
	//2. 拆分进行校验和校验
	checkSumBytes := pubkeyCheckSum[len(pubkeyCheckSum)-addressCheckSumLen:]

	//传入ripedmdhash160函数 生成校验和
	ripemd160hash := pubkeyCheckSum[:len(checkSumBytes)-addressCheckSumLen]
	//3. 生成
	checkBytes :=CheckSum(ripemd160hash)

	//4.比较
	if bytes.Compare(checkBytes,checkBytes)== 0 {
		return  true
	}
	return false
}