package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Register struct {
	UserAccount  string `json:"userAccount"`
	UserPassword string `json:"userPassword"`
	UserEmail    string `json:"userEmail"`
}

// AddUser /api/users POST
func AddUser(c *gin.Context) {
	var json Register
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Requested with wrong parameters",
			"code":    http.StatusNotAcceptable,
		})
		return
	}

	// 处理用户注册的数据
	// 校验用的账号，密码等等信息
	// 校验完成之后没有问题，就要将数据插入到数据库
}
