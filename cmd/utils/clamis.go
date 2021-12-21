package utils

import (
	"encoding/json"
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/types"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// GetUserID 获取从jwt解析出来的用户ID
func GetUserID(r *http.Request) uint {
	claims := r.Header.Get("claims")
	if claims == "" {
		global.GVA_LOG.Error("从Gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件!")
		return 0
	} else {
		var cl *types.CustomClaims
		err := json.Unmarshal([]byte(claims), &cl)
		if err != nil {
			return 0
		}
		return cl.ID
	}
}

// GetUserUuid 获取从jwt解析出来的用户UUID
func GetUserUuid(r *http.Request) uuid.UUID {
	claims := r.Header.Get("claims")
	if claims == "" {
		global.GVA_LOG.Error("从Gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件!")
		return uuid.UUID{}
	} else {
		var cl *types.CustomClaims
		err := json.Unmarshal([]byte(claims), &cl)
		if err != nil {
			return uuid.UUID{}
		}
		return cl.UUID
	}
}

// GetUserAuthorityId 获取从jwt解析出来的用户角色id
func GetUserAuthorityId(r *http.Request) string {
	claims := r.Header.Get("claims")
	if claims == "" {
		global.GVA_LOG.Error("从Gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件!")
		return ""
	} else {
		var cl *types.CustomClaims
		err := json.Unmarshal([]byte(claims), &cl)
		if err != nil {
			return ""
		}
		return cl.AuthorityId
	}
}

// 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(r *http.Request) *types.CustomClaims {
	claims := r.Header.Get("claims")
	if claims == "" {
		global.GVA_LOG.Error("从Gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件!")
		return nil
	} else {
		var cl *types.CustomClaims
		err := json.Unmarshal([]byte(claims), &cl)
		if err != nil {
			return nil
		}
		return cl
	}
}
