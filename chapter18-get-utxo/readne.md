#1.实现UTXO查找的内部实现

bc.exe createblockchain -address mage
bc.exe printchain
bc.exe getbalance


bc.exe send -from "[\"troytan\"]" -to "[\"Alice\"]" -amount "[\"100\"]

1. 实现查找数据库中所有指定地址的以花费输出
2. 实现Coinbase交易判断函数
3. 实现指定地址UTXO的函数
