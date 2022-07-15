package request
/**
 * @Com www.sixstaredu.com
 * @Author 六星教育-shineyork老师
 */

type Register struct {
	Username string `form:"username"  binding:"required,character,max=16"`
	Password string `form:"password"  binding:"required,character,min=4,max=16"`
	Phone    string `form:"phone"  binding:"required,phone"`
}

type Login struct {
	Username string `form:"username" binding:"required,character"`
	Password string `form:"password" binding:"required,character"`
}
type MemberAddress struct {
	Recipient string	`form:"recipient" binding:"required"`
	Phone string		`form:"phone" binding:"required,phone"`
	ProvinceId int		`form:"province_id" binding:"required"`
	CityId int			`form:"city_id" binding:"required"`
	DistrictId int		`form:"district_id" binding:"required"`
	Address string		`form:"address" binding:"required"`
	IsDefault byte		`form:"is_default" binding:"required"`
}