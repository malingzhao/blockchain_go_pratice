

## 测试方法
mac系统环境
export NODE_ID=4000
go build -o bc.exe main.go
./bc.exe createwallet
使用这个命令执行多次生成多个账号
./bc.exe accounts
./bc.exe createblockchain -address  34SnogDeAZb4qy4wV8BeWBS6pzs8DrKqpHQykLiGG9Lt5ZeiJkwyKhcqHNkurLU11EzaQkcH5mD3V
下面的是生成的地址 每个人的是不同的
./bc.exe send -from "[\"34SnogDeAZb4qy4wV8BeWBS6pzs8DrKqpHQykLiGG9Lt5ZeiJkwyKhcqHNkurLU11EzaQkcH5mD3V\"]" -to "[\"7iCt7umJfnfGaoe1PJaiEbbSJbtNJq1z4XyWCcUuJdwHjKQSqMGgmb2qbys7rFvy7sq5KGd9wzYwR\"]" -amount "[\"2\"]"

得到的数据库和wallet文件分别拷贝然后进行测试
