package router

import (
	"github.com/backend/api/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InitRouter 处理所有路由的挂载
func InitRouter() {
	// 初始化gin的*Engine
	router := gin.Default()
	// 挂载url找不到的处理函数
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
	})

	// 分组
	root := router.Group("/api")
	{
		// Authorization required and not websocket request
		g := root.Group("/")
		{
			user.InitManageUserRouter(g)
		}
	}

	// TODO: 目前一定要Run()，以后需要更改，参考Nginx UI的处理方式
	router.Run()
}
