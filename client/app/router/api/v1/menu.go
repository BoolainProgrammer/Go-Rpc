package v1

import "github.com/gin-gonic/gin"

func menu(gin *gin.Engine) {
	group := gin.Group("menu")
	{
		group.POST("/add-menu",
			v1.Menu.AddMenu,
		)
		group.POST("/save-menu",
			v1.Menu.SaveMenu,
		)
		group.GET("/find-menu-list",
			v1.Menu.FindMenuList,
		)
		group.GET("/find-menu-by-id",
			v1.Menu.FindMenuById,
		)

	}
}

