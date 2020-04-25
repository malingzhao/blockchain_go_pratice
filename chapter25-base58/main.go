package main

import (
	"blockchain_go_pratice/chapter25-base58/BLC"
	"fmt"
)

func main() {
	result := BLC.Base58Encode([]byte("this is the example"))
	fmt.Printf("result:%s\n", result)

	decodeResult:= BLC.Base58Decode([]byte("nj2SLMErZakmBni8xhSXtimREn1"))
	fmt.Printf("decodeResult:%s\n", decodeResult)

}