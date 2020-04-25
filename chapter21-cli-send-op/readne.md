## 转账逻辑望山与UTXO查找优化
1. 实现查找可用UTXO的函数FindSpendableUTXO
2. 实现公共UTXO查询进行转账,修改NewSimpleTransaction



go build -o bc.exe main.go
bc.exe createblockchain -address mage 
bc.exe send -from "[\"mage\"]" -to "[\"Alice\"]" -amount "[\"2\"]
bc.exe getbalance -address mage 
bc.exe getbalance -address Alice
bc.exe printchain 

