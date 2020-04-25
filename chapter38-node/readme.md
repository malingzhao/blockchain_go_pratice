## 网络实现
网路是非常重要的东西
存储数据 
分区
各种东西考虑进去的话
添加不同的节点是为了数据的同步
1. 模拟两个节点， 一个节点进行创建、转账等操作 另外的一个节点进行数据的同步
2. 通过不同的端口模拟不同的节点
3. 节点与程序进行关联
4. 钱包文件与blockchain数据库文件与节点的id号进行关联
windows系统使用的是set命令设置环境变量
通过各种系统的判断来使用不同的操作系统

设置相应的id



./bc.exe createblockchain -address  "4y5KNwSVfcrZRPaJgdesfBg2tZU9ZBhx1iq7SUHfnSQQrJPwELM81d3YWgwyKNCdi3AMXyJAUhRcK"
./bc.exe getbalance -address "4y5KNwSVfcrZRPaJgdesfBg2tZU9ZBhx1iq7SUHfnSQQrJPwELM81d3YWgwyKNCdi3AMXyJAUhRcK"

./bc.exe send -from "[\"4y5KNwSVfcrZRPaJgdesfBg2tZU9ZBhx1iq7SUHfnSQQrJPwELM81d3YWgwyKNCdi3AMXyJAUhRcK\"]" -to "[\"5GFdJB64m8FP9egbtWTzvjGqUpR1GwjLbEtUQhZeyTMsvd9TGyAt38tQGcs2ttBXrGnGSoDfsrF1L\"]" -amount "[\"2\"]"