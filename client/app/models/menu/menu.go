package menu

type Menu struct {
	Id           int    `json:"id"`
	ParentId     int    `json:"parentId"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	Type         string `json:"type"`
	Content      string `json:"content"`
	DisplayOrder int    `json:"display_order"`
	IsDisplay    string    `json:"isDisplay"`
	Status       string `json:"status"`
	CreatedUser  string `json:"createdUser"`
	CreatedTime  string `json:"createdTime"`
	UpdatedUser  string `json:"updatedUser"`
	UpdatedTime  string `json:"updatedTime"`
	Child        []Menu `gorm:"foreignKey:Id"`
}

func (model *Menu) TableName() string {
	return "t_menu"
}
