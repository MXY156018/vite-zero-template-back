package authorityhandler

import (
	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-template/cmd/internal/logic/authority"
	"go-zero-template/cmd/internal/svc"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/utils"
	"net/http"
)

func CreateAuthorityHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysAuthority
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := authority.NewAuthorityLogic(r.Context(), ctx)
		resp, err := l.Authority(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func DeleteAuthorityHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysAuthority
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := authority.NewAuthorityLogic(r.Context(), ctx)
		resp, err := l.DeleteAuthority(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func UpdateAuthorityHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysAuthority
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := authority.NewAuthorityLogic(r.Context(), ctx)
		resp, err := l.UpdateAuthority(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

//CopyAuthorityHandler 复制角色
func CopyAuthorityHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysAuthorityCopyResponse
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := authority.NewAuthorityLogic(r.Context(), ctx)
		resp, err := l.CopyAuthority(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func GetAuthorityListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageInfo
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := authority.NewAuthorityLogic(r.Context(), ctx)
		resp, err := l.GetAuthorityList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func SetDataAuthorityHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysAuthority
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := authority.NewAuthorityLogic(r.Context(), ctx)
		resp, err := l.SetDataAuthority(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}