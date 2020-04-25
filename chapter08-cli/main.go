package main

import (
	"flag"
	"fmt"
)

//命令行原理
//定义一个字符串变量
var species = flag.String("species","go","the usage of flag")
//定义一个int字符
var  num = flag.Int("ins",1,"ins nums")
func main(){
	//解析，在flag各种类型参数生效之前需要对参数进行解析
	flag.Parse()
	var s []int = []int{1,2,3,4,5}
	fmt.Println(s[2:4]) //左闭右开的区间
	fmt.Println("a string flag", *species)
	fmt.Println("ins num", *num)

}