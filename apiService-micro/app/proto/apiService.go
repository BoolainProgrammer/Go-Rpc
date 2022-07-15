package proto

type AddApiServiceReq struct {
	Id            int    `json:"id"`
	ModuleId      int    `json:"moduleId" binding:"required"`
	ApiName       string `json:"apiName" binding:"required"`
	ApiDes        string `json:"apiDes" binding:"required"`
	HttpMethod    string `json:"httpMethod" binding:"required"`
	Url           string `json:"url" binding:"required"`
	Port          string `json:"port" binding:"required"`
	AudienceIdent string `json:"audienceIdent" binding:"required"`
	Path          string `json:"path" binding:"required"`
	IsCheck       string `json:"isCheck" binding:"required"`
	ParamType     string `json:"paramType" binding:"required"`
	IsParam       string `json:"isParam" binding:"required"`
	RequestParam  string `json:"requestParam" binding:"required"`
	Status        string `json:"status" binding:"required"`
	UpdatedUser   string
}
