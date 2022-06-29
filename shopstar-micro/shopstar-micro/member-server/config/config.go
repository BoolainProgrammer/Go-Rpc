package config

import (
	"sixstar/shopstar-micro/member-server/core/cache"
	"sixstar/shopstar-micro/member-server/core/log"
	"sixstar/shopstar-micro/member-server/core/model"

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

	Mysql *model.Config `mapstructure:"mysql"`
	Cache *cache.Config `mapstructure:"cache"`
	Log *log.Config `mapstructure:"log"`

}
