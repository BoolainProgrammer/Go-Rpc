package v1

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"toboefa/app/transform/request"
	"toboefa/app/transform/response"
	"toboefa/global"
)

type Menu struct {
}

func (*Menu) AddMenu(ctx *gin.Context) {
	req := new(request.Menu)
	err := ctx.ShouldBind(req)
	if err != nil {
		response.HandleValidatorError(ctx, err)
		return
	}
	err = MenuServices.AddMenu(req)
	if err != nil {
		response.FailWithErr(err, ctx)
		return
	}
	response.Ok(ctx)
	return
}

func (*Menu) SaveMenu(ctx *gin.Context) {
	req := new(request.Menu)
	err := ctx.ShouldBind(req)
	if err != nil {
		response.HandleValidatorError(ctx, err)
		return
	}
	if req.Id < 1 {
		response.FailWithErr(global.ErrGormRecord, ctx)
		return
	}
	err = MenuServices.AddMenu(req)
	if err != nil {
		response.FailWithErr(err, ctx)
		return
	}
	response.Ok(ctx)
	return
}


func (*Menu) FindMenuList(ctx *gin.Context) {
	menuName := ctx.Query("menuName")
	isDisplay := ctx.Query("isDisplay")
	list, err := MenuServices.FindMenuList(menuName, isDisplay)
	if err != nil {
		response.FailWithErr(err, ctx)
		return
	}
	response.OkWithData(list, ctx)
}

func (*Menu) FindMenuById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	if id < 1 {
		response.FailWithErr(global.ErrGormRecord, ctx)
		return
	}
	list, err := MenuServices.FindMenu(&request.Menu{Id: id})
	if err != nil {
		response.FailWithErr(err, ctx)
		return
	}
	response.OkWithData(list, ctx)
}
