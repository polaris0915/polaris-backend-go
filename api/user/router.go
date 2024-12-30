package user

import "github.com/gin-gonic/gin"

func InitManageUserRouter(r *gin.RouterGroup) {
	// 用户注册
	r.POST("users", AddUser)
	r.DELETE("users/:id", Delete)
	r.POST("login", Login)

}
