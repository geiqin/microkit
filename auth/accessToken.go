package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/geiqin/gotools/helper/xtime"
)

type AccessToken struct {
}

// 将 JWT 字符串解密为 CustomClaims 对象
func (a *AccessToken) Decode(tokenStr string) (*AccessClaims, error) {
	t, err := jwt.ParseWithClaims(tokenStr, &AccessClaims{}, func(token *jwt.Token) (interface{}, error) {
		return tokenCfg.PrivateKey, nil
	})
	// 解密转换类型并返回
	if claims, ok := t.Claims.(*AccessClaims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// 将 User 用户信息加密为 JWT 字符串
func (a *AccessToken) Encode(mode string, user *GrantUser, store *GrantStore, limit *AccessLimit) (string, error) {
	expireTime := xtime.GetAfterDay(tokenCfg.ExpireTime, xtime.DayType).Unix()
	claims := AccessClaims{
		mode,
		user,
		store,
		limit,
		jwt.StandardClaims{
			Issuer:    tokenCfg.Issuer, //签发者
			ExpiresAt: expireTime,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(tokenCfg.PrivateKey)
}
