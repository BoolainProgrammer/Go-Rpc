package config

import (
	"sixstar/shopstar-micro/client/core/cache"
	"sixstar/shopstar-micro/client/core/log"
	"sixstar/shopstar-micro/client/core/model"
	"sixstar/shopstar-micro/client/core/rpc/client"
)

/**
 * @Com www.sixstaredu.com
 * @Author 六星教育-shineyork老师
 */

type Config struct {
	App
	Sms
	Smsbao
	Captche `mapstructure:"captcha"`
	Pay



	RpcClient *client.Config `mapstructure:"rpc_client"`
	Mysql *model.Config `mapstructure:"mysql"`
	Cache *cache.Config `mapstructure:"cache"`
	Log *log.Config `mapstructure:"log"`

}
