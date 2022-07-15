package logic

import (
	"sixstar/shopstar-micro/member-server/global"
	"sixstar/shopstar-micro/member-server/app/models"
)

/**
 * @Com www.sixstaredu.com
 * @Author 六星教育-shineyork老师
 */

type SysArea struct {
}
var AreaLogic = new(SysArea)


func (logic *SysArea) GetAddress(provinceid, cityid, districtid int) (*models.Province, *models.City,
	*models.District) {
	var (
		province = new(models.Province)
		city = new(models.City)
		district = new(models.District)
	)

	global.DB.Where("province_id = ? ", provinceid).First(province)
	global.DB.Where("city_id = ? ", cityid).First(city)
	global.DB.Where("district_id = ? ", districtid).First(district)

	return province, city, district

}
