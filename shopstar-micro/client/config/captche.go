package config


/**
 * @Com www.sixstaredu.com
 * @Author 六星教育-shineyork老师
 */
type Captche struct {
	KeyLong   int `mapstructure:"key_long"`
	ImgWidth  int `mapstructure:"img_width"`
	ImgHeight int `mapstructure:"img_height"`
}
