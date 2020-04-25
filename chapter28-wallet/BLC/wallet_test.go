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


func TestWallet_GetAddress(t *testing.T) {
	wallet :=NewWallet()
	address := wallet.GetAddress()
	fmt.Printf("the address of coin is [%s]\n",address)
}

func TestIsValidForAddressBytes(t *testing.T) {
	wallet :=NewWallet()
	address := wallet.GetAddress()
	fmt.Printf("the validation of current address is %v", IsValidForAddressBytes([]byte(address)))
}
