package request

type Menu struct {
	Id             int    `json:"id"`
	ParentId       int    `json:"parentId"`
	Name           string `json:"name" binding:"required"`
	Code           string `json:"code" binding:"required"`
	Type           string `json:"type" binding:"required"`
	Content        string `json:"content" binding:"required"`
	DisplayOrder   int    `json:"displayOrder" binding:"required"`
	IsDisplay      string `json:"isDisplay" binding:"required"`
	Status         string `json:"status" binding:"required"`
	ApiServiceList string `json:"apiServiceList" binding:"required"`
	UpdatedUser    string
	CreatedUser    string
}
