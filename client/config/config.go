package config

import (
	"toboefa/core/cache"
	"toboefa/core/log"
	"toboefa/core/model"
	"toboefa/core/rpc/client"
)

 

type Config struct {
	App

	Captche `mapstructure:"captcha"`
	RpcClient *client.Config `mapstructure:"rpc_client"`
	Mysql *model.Config `mapstructure:"mysql"`
	Cache *cache.Config `mapstructure:"cache"`
	Log *log.Config `mapstructure:"log"`

}
