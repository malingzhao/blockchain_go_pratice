## Merkle树的实现
1. Merkle节点的实现
2. Merkle树的实现
3. 交易哈希与Merkle树的联结

go build -o bc.exe main.go

bc.exe accounts




bc.exe send -from "[\"mage\"]" -to "[\"Alice\"]" -amount "[\"5\"]
bc.exe send -from "[\"mage\",\"Alice\"]" -to "[\"Alice\",\"mage\"]" -amount "[\"5\",\"2\"]"
//mage->alice 5 mage 5 alice 5
//alice-> mage 2 mage7 alcie 3

Hash:00006badb9af6474b8e415d6efb225aafc5616fb06e44760389949958feb8308


bc.exe createblockchain -address 4y5KNwSVfcrZRPaJgdesfBg2tZU9ZBhx1iq7SUHfnSQQrJPwELM81d3YWgwyKNCdi3AMXyJAUhRcK
bc.exe getbalance -address 4y5KNwSVfcrZRPaJgdesfBg2tZU9ZBhx1iq7SUHfnSQQrJPwELM81d3YWgwyKNCdi3AMXyJAUhRcK
bc.exe getbalance -address 5
GFdJB64m8FP9egbtWTzvjGqUpR1GwjLbEtUQhZeyTMsvd9TGyAt38tQGcs2ttBXrGnGSoDfsrF1L
./bc.exe send -from "[\"4y5KNwSVfcrZRPaJgdesfBg2tZU9ZBhx1iq7SUHfnSQQrJPwELM81d3YWgwyKNCdi3AMXyJAUhRcK\"]" -to "[\"5GFdJB64m8FP9egbtWTzvjGqUpR1GwjLbEtUQhZeyTMsvd9TGyAt38tQGcs2ttBXrGnGSoDfsrF1L\"]" -amount "[\"2\"]
----------------------------------
        Hash:0000c08badd2a13f779d95cb2aaeab7d4d51b7ba63ba413983abf2982c0f27f0
        PrevBlockHash:0000d397d1abb76f8433a234ee9947ad3e103effedabe85afacf59961aa48524
        TimeStamp:1587297008
        Height:3
        Nonce:19099
        Txs:[0xc0000a20a0]
                 tx-hash:0dffdbb5ab1e5719280c7efa365ca36ca4178b202d6994d080addbf132590baa
                输入....
                vin-TxHash:55d1afac7eb30b8bb81074190442d1031b3b9d0e4d8d4e3fc220f7542fe10459
                vin-vout:1
                vin-scriptSig:405fe4a0d7db7aac883af2c1c7a9adac34e2f0df609b01d16eee26f6003310a786a0517eeafaa0af3235ce2d75215c9b9dd9f0ee8f8e387977d47df3f
81aae98
                vin-scriptSig:1402303e69f70f476bcbf3b352a3c3c19d5054e6048c4dc629c32ee085a64b81c97e66db693ca48f58647eded0bd4cd15d460adfdbfb0acc697057646
f79ce9e
                输出...
                vout-value：2
                vout-scriptPubKey:9d69d0acf13bc3353d7c7dc06edeef882d440799cf4779fa1dc88744ec1d9bf6cce870dd357e226bba95803695cea2f21a488e50
                vout-value：3
                vout-scriptPubKey:927a2434b273e11102cfa810a6039d24bd8f7f862d2d9b58a785bb8656bfc9eed1615bd5d523441a1da9c0e493fa5d8639eee7bc
-----------------------------------
        Hash:0000d397d1abb76f8433a234ee9947ad3e103effedabe85afacf59961aa48524
        PrevBlockHash:000068b60c2daa7cb46d3351e0988d99dc2582f5579ed8cf1b7e60bac050bc51
        TimeStamp:1587296530
        Height:2
        Nonce:244491
        Txs:[0xc0000a2140]
                 tx-hash:55d1afac7eb30b8bb81074190442d1031b3b9d0e4d8d4e3fc220f7542fe10459
                输入....
                vin-TxHash:834dfc1836cdffdc8d89d4def743377d73050cd1e3a3bc9c307177d77f2cd53a
                vin-vout:0
                vin-scriptSig:405fe4a0d7db7aac883af2c1c7a9adac34e2f0df609b01d16eee26f6003310a786a0517eeafaa0af3235ce2d75215c9b9dd9f0ee8f8e387977d47df3f
81aae98
                vin-scriptSig:60ef719bcd824184b94437478c51d3a381d265e6943553a630ce0170152c1b03a33f58ff61182de3215f74e898b874d25b4dd89f624f19b2fa6f9b781
0bd06b4
                输出...
                vout-value：5
                vout-scriptPubKey:9d69d0acf13bc3353d7c7dc06edeef882d440799cf4779fa1dc88744ec1d9bf6cce870dd357e226bba95803695cea2f21a488e50
                vout-value：5
                vout-scriptPubKey:927a2434b273e11102cfa810a6039d24bd8f7f862d2d9b58a785bb8656bfc9eed1615bd5d523441a1da9c0e493fa5d8639eee7bc
-----------------------------------
        Hash:000068b60c2daa7cb46d3351e0988d99dc2582f5579ed8cf1b7e60bac050bc51
        PrevBlockHash:
        TimeStamp:1587295598
        Height:1
        Nonce:24125
        Txs:[0xc0000a21e0]
                 tx-hash:834dfc1836cdffdc8d89d4def743377d73050cd1e3a3bc9c307177d77f2cd53a
                输入....
                vin-TxHash:
                vin-vout:-1
                vin-scriptSig:
                vin-scriptSig:
                输出...
                vout-value：10
                vout-scriptPubKey:927a2434b273e11102cfa810a6039d24bd8f7f862d2d9b58a785bb8656bfc9eed1615bd5d523441a1da9c0e493fa5d8639eee7bc



./bc.exe createblockchain -address  "4y5KNwSVfcrZRPaJgdesfBg2tZU9ZBhx1iq7SUHfnSQQrJPwELM81d3YWgwyKNCdi3AMXyJAUhRcK"
./bc.exe getbalance -address "4y5KNwSVfcrZRPaJgdesfBg2tZU9ZBhx1iq7SUHfnSQQrJPwELM81d3YWgwyKNCdi3AMXyJAUhRcK"

./bc.exe send -from "[\"4y5KNwSVfcrZRPaJgdesfBg2tZU9ZBhx1iq7SUHfnSQQrJPwELM81d3YWgwyKNCdi3AMXyJAUhRcK\"]" -to "[\"5GFdJB64m8FP9egbtWTzvjGqUpR1GwjLbEtUQhZeyTMsvd9TGyAt38tQGcs2ttBXrGnGSoDfsrF1L\"]" -amount "[\"2\"]"