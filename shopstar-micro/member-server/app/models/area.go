package models

/**
 * @Com www.sixstaredu.com
 * @Author 六星教育-shineyork老师
 */

// 地域
type Area struct {
	AreaId   int
	AreaName string
	Sort     int
}

func (model *Area) TableName() string {
	return "sys_area"
}

// 省份
type Province struct {
	ProvinceId   int
	Area_id      int
	ProvinceName string
	Sort         int
}
func (model *Province) TableName() string {
	return "sys_province"
}
// 城市
type City struct {
	CityId     int
	ProvinceId int
	CityName   string
	Zipcode    string
	Sort       int
}
func (model *City) TableName() string {
	return "sys_city"
}
// 地区
type District struct {
	DistrictId   int
	CityId       int
	DistrictName string
	Sort         int
}
func (model *District) TableName() string {
	return "sys_district"
}
