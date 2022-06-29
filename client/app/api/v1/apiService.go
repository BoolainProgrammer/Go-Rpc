package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"toboefa/app/transform/request"
	"toboefa/app/transform/response"
	"toboefa/global"
)

type ApiService struct{}

func (*ApiService) AddApiService(ctx *gin.Context) {
	req := new(request.ApiService)
	if err := ctx.ShouldBind(req); err != nil {
		response.HandleValidatorError(ctx, err)
		return
	}
	req.UpdatedUser = "admin"
	req.CreatedUser = "admin"
	err := ApiServices.AddApiService(req)
	if err != nil {
		response.FailWithErr(err, ctx)
		zap.Any("error", err.Error())
		return
	}
	response.Ok(ctx)
}

func (*ApiService) FindApiServiceList(ctx *gin.Context) {
	moduleId, _ := strconv.Atoi(ctx.Query("moduleId"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	apiName := ctx.Query("apiName")
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	if pageSize < 1 {
		pageSize = 10
	}
	if page < 1 {
		page = 1
	}
	list, total := ApiServices.GetApiServiceList(moduleId, page, pageSize, apiName)
	response.OkWithData(&response.ApiList{
		List:     list,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, ctx)
	return
}
func (*ApiService) FindApiServiceById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	if id < 1 {
		response.FailWithErr(global.ErrGormRecord, ctx)
		return
	}
	list, err := ApiServices.GetApiServiceById(id)
	if err != nil {
		response.FailWithErr(err, ctx)
		return
	}
	response.OkWithData(list, ctx)
	return
}

func (*ApiService) SaveApiService(ctx *gin.Context) {
	req := new(request.ApiService)
	if err := ctx.ShouldBind(req); err != nil {
		fmt.Println(err.Error())
		response.HandleValidatorError(ctx, err)
		return
	}
	if req.Id < 1 {
		response.FailWithErr(global.ErrGormRecord, ctx)
		return
	}
	req.UpdatedUser = "admin"
	err := ApiServices.SaveApiService(req)
	if err != nil {
		response.FailWithErr(err, ctx)
		return
	}
	response.Ok(ctx)
	return
}
