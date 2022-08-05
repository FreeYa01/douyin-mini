package util

import (
	"douyin-mini/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)
type Claims struct {
	UserID       int64
	jwt.StandardClaims
}
// GenerationToken 生成token,携带用户的姓名和id
func GenerationToken(UserID int64) (string,error) {
	claims := Claims{
		UserID: UserID,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt:time.Now().Add(12*time.Hour).Unix(),
			Issuer: "jwtF",
		},
	}
	token,err := jwt.NewWithClaims(jwt.SigningMethodHS256,claims).SignedString([]byte(setting.Set.JwtSecret))
	return token,err
}
// ParseToken 解析token
func ParseToken(token string)(*Claims,error)  {
	tokenClaims,err := jwt.ParseWithClaims(token,&Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 根据密钥进行解析
		return []byte(setting.Set.JwtSecret),nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if c,ok := tokenClaims.Claims.(*Claims);ok && tokenClaims.Valid{
			return c,nil
		}
	}
	return nil,err
}
