package service

import (
	v1 "awesomeProject/core/v1"
	"crypto/rsa"
)

type Jwt interface {
	GetRsaKeys() *v1.RsaKeys
	GenerateToken(userUUID string, duration int64, data interface{}) (token string, refreshToken string, err error)
	IsTokenValid(tokenString string) bool
	GetPrivateKey() *rsa.PrivateKey
	GetPublicKey() *rsa.PublicKey
}
