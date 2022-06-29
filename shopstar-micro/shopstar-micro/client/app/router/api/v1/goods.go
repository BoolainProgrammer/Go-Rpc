package v1

import (
	"github.com/gin-gonic/gin"
)

func Goods(gin *gin.Engine)  {
	group := gin.Group("goods")
	{
		group.GET("/attrs", v1.Goods.AttributeList)
		group.GET("/assemblyattr", v1.GoodsSku.AssemblyAttr)
		group.POST("/skutbdetail", v1.GoodsSku.Tbdetail)
	}
}

