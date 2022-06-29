package menuApeRelation

type MenuApeRelation struct {
	Id           int `json:"id"`
	MenuId       int `json:"menuId"`
	ApiServiceId int `json:"apiServiceId"`
	CreatedUser  string `json:"createdUser"`
	CreatedTime  string `json:"createdTime"`
	UpdatedUser  string `json:"updatedUser"`
	UpdatedTime  string `json:"updatedTime"`
}

func (model *MenuApeRelation) TableName() string {
	return "t_menu_api_relation"
}
