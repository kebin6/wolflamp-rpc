package config

import (
	"github.com/suyuan32/simple-admin-common/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DatabaseConf config.DatabaseConf
	RedisConf    config.RedisConf
	CoreRpc      zrpc.RpcClientConf
	GcicsConf    GcicsConf
}

type GcicsConf struct {
	Host       string
	AppId      string
	AppSecret  string
	NotifyUrl  string
	SuccessUrl string
	FailUrl    string
}
