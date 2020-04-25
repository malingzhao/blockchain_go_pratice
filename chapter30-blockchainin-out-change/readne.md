## 实现多笔交易转账
1. 实现输入结构钱包功能结合
2. 实现输出结构与钱包功能结合
3. 调用边修改


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

