package main

import (
	"fmt"
	"os"
)

func main() {

	os.Setenv("NODE_ID", "12")
	getenv := os.Getenv("NODE_ID")

	fmt.Println(getenv)
}
