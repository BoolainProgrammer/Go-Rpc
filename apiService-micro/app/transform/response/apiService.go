package response

import "toboefa/app/models/apiService"

type ApiList struct {
	List     []*apiService.Api `json:"list"`
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}
