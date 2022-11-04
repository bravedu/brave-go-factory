package util

import (
	"github.com/go-pay/gopay/pkg/jwt"
	"time"
)

type JWT struct {
	SecretKey []byte
}

type CustomClaims struct {
	jwt.StandardClaims
	UID int `json:"uid"`
}

func NewJWT(secret string) *JWT {
	return &JWT{[]byte(secret)}
}

func (j *JWT) GenerateToken(uid int) (tokenStr string, err error) {
	claims := new(CustomClaims)
	claims.UID = uid
	claims.ExpiresAt = time.Now().Add(time.Hour * 168).Unix()
	claims.Issuer = "musenetwork.org"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用自定义字符串加密，解密同密
	tokenStr, err = token.SignedString(j.SecretKey)
	return
}

func (j *JWT) ParseToken(token string) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return j.SecretKey, nil
		},
	)
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
