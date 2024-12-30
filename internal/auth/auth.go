package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	// token有效时间
	expireTime = 24 * 3600 * time.Second
	// 签发token和解析token的jwtKey
	jwtKey = "go-backend"
)

type Auth struct {
	// RegisteredClaims 首先嵌套jwt.RegisteredClaims
	jwt.RegisteredClaims
	// 应该要包含一些用户的信息。
	// 具体包含哪些用户的信息呢？
	Identity string
	Role     string
}

// NewToken 获取用户的token
func NewToken(identity, role string) (string, error) {
	// 在expireTime之后过期
	ex := time.Now().Add(expireTime)
	auth := Auth{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(ex), // 填写过期时间
		},
		Identity: identity,
		Role:     role,
	}
	// 得到Token的结构体
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, auth)
	// 转换string的token
	token, err := t.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

// ParseToken 解析token，即校验token是否有效
func ParseToken(token string) (*Auth, error) {
	auth := &Auth{}
	t, err := jwt.ParseWithClaims(token, auth, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	auth, ok := t.Claims.(*Auth)
	if !ok {
		return nil, errors.New("invalid token")
	}
	return auth, nil
}
