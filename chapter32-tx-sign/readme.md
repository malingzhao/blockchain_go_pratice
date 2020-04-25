## 实现交易的签名
1. 在什么地方对交易进行签名
   在交易生成的时候对交易进行签名
2. 在什么地方对交易进行验证
   在交易打包进入区块之前进行验证
3. 是否需要对交易中的所有属性进行签名和验证（哪些属性是必须要进行交易签名的）


理论相关
交易的本质：
  交易实际上就是解锁指定地址的output 重新分配他们的值在加锁到新的output中
为了交易安全 必须加密（签名）的数据如下：
  1. 保存在已解锁的output中的公钥哈希 代表交易的发送者
  2. 保存新生成的output的公钥哈希，代表交易的签收者
  3. 新生成的output所包含的value
  
  
流程：
  1. 生成交易
  2. 对交易进行签名
     1. 判断该交易是不是coinbase 交易 如果是coinbase交易不签名
     2. 查找当前交易的输入所引用的交易（输出所在的交易）
     3. 提取需要签名的属性
     4.签名 
    
    
  3. 验证交易
  4. 打包





go build -o bc.exe main.go
bc.exe createblockchain -address mage 
bc.exe send -from "[\"mage\"]" -to "[\"Alice\"]" -amount "[\"2\"]
bc.exe getbalance -address mage 
bc.exe getbalance -address Alice
bc.exe printchain 



bc.exe send -from "[\"mage\"]" -to "[\"Alice\"]" -amount "[\"5\"]
bc.exe send -from "[\"mage\",\"Alice\"]" -to "[\"Alice\",\"mage\"]" -amount "[\"5\",\"2\"]"
//mage->alice 5 mage 5 alice 5
//alice-> mage 2 mage7 alcie 3

Hash:00006badb9af6474b8e415d6efb225aafc5616fb06e44760389949958feb8308


bc.exe send -from "[\"6hvjtdhRYNPSLRrdu9yFJKm16mL3pHAeqchVdtQzFKDnoZTBAWbvF5iXtnmsUuHWExpwAKDCgdJyK\"]" -to "[\"5GFdJB64m8FP9egbtWTzvjGqUpR1GwjLbEtUQhZeyTMsvd9TGyAt38tQGcs2ttBXrGnGSoDfsrF1L\"]" -amount "[\"5\"]

打印区块链完整信息.......
-----------------------------------
        Hash:0000cb3cae6772cc4e65625206efe5d33b327473587c8d9ec4cb48a02fd3bfa1
        PrevBlockHash:0000811d1db1bf5d4984e48fa2b63675eb7b64cb0e5bd34b6b2d972aa0a21c78
        TimeStamp:1587259968
        Height:2
        Nonce:26917
        Txs:[0xc000098050 0xc0000980a0]
                 tx-hash:c1449fabbc9c7c097d1d7df8a60aaaa90148165755218faaa1061e99aeaa9e67
                输入....
                vin-TxHash:f449245013e0d88c713617511436ef2f546ff3c600a44976c0f0a651f6d722f6
                vin-vout:0
                vin-scriptSig:mage
                输出...
                vout-value：5
                vout-scriptPubKey:Alice
                vout-value：5
                vout-scriptPubKey:mage
                 tx-hash:48a984b16dd36fdabb9efea9d1a28e8fbf5a5e518d2ce7e840c17178f47c21a9
                输入....
                vin-TxHash:c1449fabbc9c7c097d1d7df8a60aaaa90148165755218faaa1061e99aeaa9e67
                vin-vout:0
                vin-scriptSig:Alice
                输出...
                vout-value：2
                vout-scriptPubKey:mage
                vout-value：3
                vout-scriptPubKey:Alice
-----------------------------------
        Hash:0000811d1db1bf5d4984e48fa2b63675eb7b64cb0e5bd34b6b2d972aa0a21c78
        PrevBlockHash:
        TimeStamp:1587259194
        Height:1
        Nonce:52378
        Txs:[0xc0000980f0]
                 tx-hash:f449245013e0d88c713617511436ef2f546ff3c600a44976c0f0a651f6d722f6
                输入....
                vin-TxHash:
                vin-vout:-1
                vin-scriptSig:syetm reward
                输出...
                vout-value：10
                vout-scriptPubKey:mage

