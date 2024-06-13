package svc

import (
	"github.com/kebin6/wolflamp-rpc/ent"
	_ "github.com/kebin6/wolflamp-rpc/ent/runtime"
	"github.com/kebin6/wolflamp-rpc/internal/config"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-core/rpc/coreclient"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config  config.Config
	DB      *ent.Client
	Redis   redis.UniversalClient
	CoreRpc coreclient.Core
	Trans   *i18n.Translator
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)

	return &ServiceContext{
		Config:  c,
		DB:      db,
		Redis:   c.RedisConf.MustNewUniversalRedis(),
		CoreRpc: coreclient.NewCore(zrpc.MustNewClient(c.CoreRpc)),
		Trans:   i18n.NewTranslator(i18n.LocaleFS),
	}
}
