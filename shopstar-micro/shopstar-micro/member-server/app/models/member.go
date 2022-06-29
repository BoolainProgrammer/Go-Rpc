package models

/**
 * @Com www.sixstaredu.com
 * @Author 六星教育-shineyork老师
 */

type Member struct {

	Id        string
	Username  string
	Password  string
	Status    int
	Nickname  string
	Headerimg string
	Phone     string

	Address []Address `gorm:"foreignKey:UserId"`

}

func (model *Member) TableName() string {
	return "member"
}
