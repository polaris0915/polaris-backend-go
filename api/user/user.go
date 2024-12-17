package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserJson struct {
	Name     string `json:"name" binding:"required,max=255"`
	Password string `json:"password" binding:"max=255"`
}

// AddUser /api/users POST
func AddUser(c *gin.Context) {
	var json UserJson
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Requested with wrong parameters",
			"code":    http.StatusNotAcceptable,
		})
		return
	}
}
