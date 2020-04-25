#1. 实现生成普通交易
1. 实现生成交易普通函数
2. 修改挖矿函数调用NewSimpleTransaction()
3. 通过CLI接口实现普通转账
bc.exe send -from "[\"troytan\"]" -to "[\"Alice\"]" -amount "[\"100\"]