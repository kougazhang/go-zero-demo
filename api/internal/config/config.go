package config

import (
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	// 手动添加
	// RpcClientConf 是 rpc 客户端的配置, 用来解析在 blog-api.yaml 中的配置
	User zrpc.RpcClientConf
}
