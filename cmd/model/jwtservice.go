package model

import (
	"errors"
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/types"
	"gorm.io/gorm"
	"time"
)

type JwtService struct {
}

var JwtServiceApp = new(JwtService)

func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	err := global.GVA_DB.Where("jwt = ?", jwt).First(&types.JwtBlacklist{}).Error
	isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	return !isNotFound
}
func (jwtService *JwtService) JsonInBlacklist(jwtList types.JwtBlacklist) (err error) {
	err = global.GVA_DB.Model(&types.JwtBlacklist{}).Create(&jwtList).Error
	return
}
func (jwtService *JwtService) GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.GVA_REDIS.Get(userName).Result()
	return err, redisJWT
}
func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.GVA_CONFIG.JWT.ExpiresTime) * time.Second
	err = global.GVA_REDIS.Set(userName, jwt, timer).Err()
	return err
}
