package main

import (
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/ripemd160"
)

func main() {
	//sha256

	hash :=sha256.New()
   hash.Write([]byte("eth1808"))
	bytes:=hash.Sum(nil)
	fmt.Printf("sha256:%x\n",bytes)



	r160:=ripemd160.New()
	r160.Write(bytes)

	byteRipemd :=r160.Sum(nil)
	fmt.Printf("ripemd:%x\n",byteRipemd)
}
