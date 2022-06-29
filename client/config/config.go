package config

import (
	"toboefa/core/cache"
	"toboefa/core/log"
	"toboefa/core/model"

)

 

type Config struct {
	App

	Captche `mapstructure:"captcha"`

	Mysql *model.Config `mapstructure:"mysql"`
	Cache *cache.Config `mapstructure:"cache"`
	Log *log.Config `mapstructure:"log"`

}
