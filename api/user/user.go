package user

import (
	"errors"
	"github.com/backend/internal/auth"
	"github.com/backend/model"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

var (
	AccountError         = errors.New("用户名重复")
	AccountNotFoundError = errors.New("用户账号不存在")
	PasswordError        = errors.New("密码错误")
	SystemError          = errors.New("系统错误")
	ParamsError          = errors.New("参数错误")
	EmailError           = errors.New("邮箱重复")
	BanError             = errors.New("无权限")
)

var (
	Default = "user"
	Admin   = "admin"
	Ban     = "ban"
)

type UserJson struct {
	UserAccount  string `json:"userAccount"`
	UserPassword string `json:"userPassword"`
	UserEmail    string `json:"userEmail"`
}

// AddUser /api/users POST
func AddUser(c *gin.Context) {
	var json UserJson
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": ParamsError.Error(),
			"code":    http.StatusNotAcceptable,
		})
		return
	}

	// 做基本数据的校验，他们不能为空
	if json.UserAccount == "" || json.UserPassword == "" || json.UserEmail == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": ParamsError.Error(),
			"code":    http.StatusNotAcceptable,
		})
		return
	}

	// 校验，校验账号是否已经存在
	user := &model.User{}
	if err := model.UseDB().Model(user).Where("account = ?", json.UserAccount).First(user).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"message": AccountError.Error(),
				"code":    http.StatusNotAcceptable,
			})
			return
		}
	}
	// 如果查出来是一样，那么不允许
	if user.Account == json.UserAccount {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": AccountError.Error(),
			"code":    http.StatusNotAcceptable,
		})
		return
	}

	// 对用户密码进行加密
	hashed, _ := bcrypt.GenerateFromPassword([]byte(json.UserPassword), bcrypt.DefaultCost)
	user.Identity = uuid.NewV4().String() // 生成uuid
	user.Account = json.UserAccount
	user.Password = string(hashed)
	user.Email = json.UserEmail
	user.Role = Default // 设置用户的默认角色为"user"

	if err := model.UseDB().Save(user).Error; err != nil {
		// TODO: log记录一下
		// err.(*mysql.MySQLError) 断言
		// 断言指的是，将这个表达式中的 "a.(type)" a 转换成 type类型
		// "断言之后的数据, ok := a.(type)"  ---> ok表示是否断言成功，平常建议用这种
		// "断言之后的数据 := a.(type)"
		if err.(*mysql.MySQLError).Number == 1062 {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"message": EmailError.Error(),
				"code":    http.StatusNotAcceptable,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": SystemError.Error(),
			"code":    http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": true,
		"code": http.StatusOK,
	})
}

// Login /api/login POST
func Login(c *gin.Context) {
	var json UserJson
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": ParamsError.Error(),
			"code":    http.StatusNotAcceptable,
		})
		return
	}

	if json.UserAccount == "" || json.UserPassword == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": ParamsError.Error(),
			"code":    http.StatusNotAcceptable,
		})
		return
	}

	// 登录逻辑
	user := &model.User{}
	if err := model.UseDB().Model(user).Where("account = ?", json.UserAccount).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": AccountNotFoundError.Error(),
				"code":    http.StatusNotFound,
			})
			return
		}
	}
	// 校验用户是否已经被拉入黑名单
	if user.Role == Ban {
		c.JSON(http.StatusForbidden, gin.H{
			"message": BanError.Error(),
			"code":    http.StatusForbidden,
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.UserPassword)); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": PasswordError.Error(),
			"code":    http.StatusNotAcceptable,
		})
		return
	}

	//a, _ := auth.ParseToken(c.Request.Header.Get("Authorization"))
	//fmt.Printf("auth: %v\n", a)
	// TODO: 没有写完
	// 返回token
	// 问题是： 什么是token？？？
	// 生成的token，让放在请求头，或者cookie或者session中
	token, err := auth.NewToken(user.Identity, user.Role)
	if err != nil {
		// TODO: log记录
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": SystemError.Error(),
			"code":    http.StatusInternalServerError,
		})
		return
	}
	c.Header("Authorization", token)

	c.JSON(http.StatusOK, gin.H{
		"data": true,
		"code": http.StatusOK,
	})
}

// Delete /api/users/:id DELETE
func Delete(c *gin.Context) {
	// /api/users/420727bd-9809-45cf-92db-cc5d1e256bdb
	identity := c.Param("id")
	// 只有管理员或者用户本人才可以进行注销操作
	// 校验身份
	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		// 没有权限
		c.JSON(http.StatusForbidden, gin.H{
			"message": BanError.Error(),
			"code":    http.StatusForbidden,
		})
		return
	}
	authStruct, err := auth.ParseToken(authorization)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": BanError.Error(),
			"code":    http.StatusForbidden,
		})
		return
	}
	// 如果是本人或者管理员
	// 有权限
	if authStruct.Identity == identity || authStruct.Role == Admin {
		// 实现软删除
		if err := model.UseDB().Delete(&model.User{}, "identity = ?", identity).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": SystemError.Error(),
				"code":    http.StatusInternalServerError,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": true,
			"code": http.StatusOK,
		})
		return
	}
	// 如果不是本人或者管理员
	// 没有权限
	c.JSON(http.StatusForbidden, gin.H{
		"message": BanError.Error(),
		"code":    http.StatusForbidden,
	})
}
