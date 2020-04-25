package BLC

//网络服务常量管理
//命令长度

const PROTOCOL = "tcp"
const COMMAND_LENGTH = 12

//命令的分类
const (
	//验证当前节点末端是否是最新区块
	CMD_VERSION = "version"
	///从最长的链上获取区块
	CMD_GETBLOKCS = "getblocks"
	//向其他的节点展示当前节点有哪些区块
	CMD_INV = "inv"
	//请求指定区块
	CMD_GETDATA = "getdata"
	//接收到区块之后进行处理
	CMD_BLOCK = "block"
	)


