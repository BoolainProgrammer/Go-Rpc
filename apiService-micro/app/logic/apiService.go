package logic

import (
	"gorm.io/gorm"
	"toboefa/app/models/apiService"
	"toboefa/app/proto"
	"toboefa/app/util"
	"toboefa/global"
)

type Api struct{}

var ApiServiceLogic = new(Api)

//AddApiService
//创建服务
func (logic *Api) AddApiService(req proto.AddApiServiceReq) error {
	_, err := ApiServiceLogic.GetApiServiceByApiName(req.ApiName)

	if err != global.ErrGormRecordNotFound {

		if err != nil {
			return global.ErrApiNameExist
		}
		return global.ErrCreateApi
	}
	_, errByUrl := ApiServiceLogic.GetApiServiceByUrl(&req)

	if errByUrl != global.ErrGormRecordNotFound {
		if errByUrl == nil {
			return global.ErrApiUrlExist
		}
		return global.ErrCreateApi
	}
	apiInfo :=unmarshalApiService(&req)
	apiInfo.CreatedUser = req.UpdatedUser
	apiInfo.CreatedTime = util.TimeNowStr()
	apiInfo.UpdatedUser = req.UpdatedUser
	apiInfo.UpdatedTime = util.TimeNowStr()
	err = global.DB.Create(apiInfo).Error
	if err != nil {
		return global.ErrCreateApi
	}
	return nil
}

func (logic *Api) GetApiServiceList(moduleId int, page int, pageSize int, apiName string) ([]*apiService.Api, int) {
	apiList := ([]*apiService.Api)(nil)
	tx := global.DB
	if moduleId > 0 {
		tx = tx.Where("module_id = ?", moduleId)
	}
	if len(apiName) > 0 {
		tx = tx.Where("api_name like ? ", "%"+apiName+"%")
	}
	var total int64
	tx.Model(&apiService.Api{}).Count(&total)
	if page > 0 {
		tx = tx.Limit(pageSize).Offset((page - 1) * pageSize).Order("create_time desc")
	}
	tx.Find(&apiList)
	return apiList, int(total)
}

func (logic *Api) SaveApiService(req *proto.AddApiServiceReq) error {
	apiInfo, err := ApiServiceLogic.GetApiServiceByApiName(req.ApiName)
	if err != nil && err != global.ErrGormRecordNotFound {
		return global.ErrUpdateApi
	}
	if apiInfo.Id != req.Id && apiInfo.Id > 0 {
		return global.ErrApiNameExist
	}
	apiData, err := ApiServiceLogic.GetApiServiceByUrl(req)

	if err != global.ErrGormRecordNotFound && err != nil {
		return global.ErrUpdateApi
	}
	if apiData.Id != req.Id && apiData.Id > 0 {
		return global.ErrApiUrlExist
	}
	saveApiData := unmarshalApiService(req)
	saveApiData.Id = req.Id
	saveApiData.UpdatedUser = req.UpdatedUser
	saveApiData.UpdatedTime = util.TimeNowStr()
	err = global.DB.Omit("created_time").Omit("created_user").Save(&saveApiData).Error
	if err != nil {
		return global.ErrUpdateApi
	}
	return nil
}

//GetApiServiceByApiName
//通过服务名称查询数据
func (logic *Api) GetApiServiceByApiName(apiName string) (*apiService.Api, error) {
	apiInfo := new(apiService.Api)
	err := global.DB.Where("api_name = ?", apiName).First(&apiInfo).Error
	if err == gorm.ErrRecordNotFound {
		return apiInfo, global.ErrGormRecordNotFound
	}
	if err != nil {
		return apiInfo, global.ErrGormRecord

	}
	return apiInfo, nil
}

//GetApiServiceById
//通过服务id查询数据
func (logic *Api) GetApiServiceById(id int) (*apiService.Api, error) {
	apiInfo := new(apiService.Api)
	err := global.DB.Where("id = ?", id).First(&apiInfo).Error
	if err == gorm.ErrRecordNotFound {
		return apiInfo, global.ErrGormRecordNotFound
	}
	if err != nil {
		return apiInfo, global.ErrGormRecord

	}
	return apiInfo, nil
}

//GetApiServiceByUrl
//通过url查询是数据
func (logic *Api) GetApiServiceByUrl(req *proto.AddApiServiceReq) (*apiService.Api, error) {
	apiInfo := new(apiService.Api)
	tx := global.DB
	if len(req.Url) > 0 {
		tx = tx.Where("url = ?", req.Url)
	}
	if len(req.HttpMethod) > 0 {
		tx = tx.Where("http_method = ?", req.HttpMethod)
	}
	if len(req.AudienceIdent) > 0 {
		tx = tx.Where("audience_ident = ?", req.AudienceIdent)
	}
	if len(req.Port) > 0 {
		tx = tx.Where("port = ?", req.Port)
	}
	if len(req.Path) > 0 {
		tx = tx.Where("path = ?", req.Path)
	}
	err := tx.First(apiInfo).Error
	if err == gorm.ErrRecordNotFound {
		return apiInfo, global.ErrGormRecordNotFound
	}
	if err != nil {
		return apiInfo, global.ErrGormRecord
	}
	return apiInfo, nil
}

func unmarshalApiService(req *proto.AddApiServiceReq) (apiInfo *apiService.Api) {
	apiInfo.ApiName = req.ApiName
	apiInfo.ApiDes = req.ApiDes
	apiInfo.HttpMethod = req.HttpMethod
	apiInfo.Url = req.Url
	apiInfo.ModuleId = req.ModuleId
	apiInfo.Port = req.Port
	apiInfo.Path = req.Path
	apiInfo.IsCheck = req.IsCheck
	apiInfo.ParamType = req.ParamType
	apiInfo.AudienceIdent = req.AudienceIdent
	apiInfo.RequestParam = req.RequestParam
	apiInfo.Status = req.Status
	return apiInfo
}
