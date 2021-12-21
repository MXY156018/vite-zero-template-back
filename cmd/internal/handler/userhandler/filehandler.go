package userhandler

import (
	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-template/cmd/internal/logic/user"
	"go-zero-template/cmd/internal/svc"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/utils"
	"net/http"
)

func UserRegisterHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Register
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := user.NewUserLogic(r.Context(), ctx)
		resp, err := l.UserRegister(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func UserChangePasswordHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChangePasswordStruct
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := user.NewUserLogic(r.Context(), ctx)
		resp, err := l.UserChangePassword(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func UserListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageInfo
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := user.NewUserLogic(r.Context(), ctx)
		resp, err := l.UserList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func UserSetUserAuthorityHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req types.SetUserAuth
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := user.NewUserLogic(r.Context(), ctx)
		resp, err := l.UserSetUserAuthority(req, r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func UserDeleteUserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req types.GetById
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := user.NewUserLogic(r.Context(), ctx)
		resp, err := l.DeleteUser(req, r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func UserSetUserInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req types.SysUser
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := user.NewUserLogic(r.Context(), ctx)
		resp, err := l.SetUserInfo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func UserSetUserAuthoritiesHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetUserAuthorities
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := user.NewUserLogic(r.Context(), ctx)
		resp, err := l.SetUserAuthorities(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func UserGetUserInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := user.NewUserLogic(r.Context(), ctx)
		resp, err := l.GetUserInfo(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
