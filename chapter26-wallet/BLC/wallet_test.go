package BLC

import (
	"fmt"
	"testing"
)

func TestNewWallet(t *testing.T) {
	wallet:=NewWallet()
	fmt.Printf("private key : %v\n",wallet.PrivateKey)
	fmt.Printf("private key : %v\n",wallet.PublicKey)
	fmt.Printf("private key : %v\n",wallet)
}
