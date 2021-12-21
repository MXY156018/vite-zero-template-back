package userhandler

import (
	"github.com/tal-tech/go-zero/rest"
	"go-zero-template/cmd/internal/svc"
	"net/http"
)

func RegisterHandlersAutocode(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Jwt, serverCtx.Casbin, serverCtx.OperateRecord},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/register",
					Handler: UserRegisterHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/changePassword",
					Handler: UserChangePasswordHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/getUserList",
					Handler: UserListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/setUserAuthority",
					Handler: UserSetUserAuthorityHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/user/deleteUser",
					Handler: UserDeleteUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/user/setUserInfo",
					Handler: UserSetUserInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/setUserAuthorities",
					Handler: UserSetUserAuthoritiesHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/getUserInfo",
					Handler: UserGetUserInfoHandler(serverCtx),
				},
			}...,
		),
	)
}
