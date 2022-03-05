package menuhandler

import (
	"github.com/zeromicro/go-zero/rest"
	"go-zero-template/cmd/internal/svc"
	"net/http"
)

func RegisterHandlersAutocode(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Jwt, serverCtx.Casbin, serverCtx.OperateRecord},
			//[]rest.Middleware{serverCtx.OperateRecord},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/menu/getMenu",
					Handler: MenuGetMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/getMenuList",
					Handler: MenuGetMenuListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/addBaseMenu",
					Handler: MenuAddBaseMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/getBaseMenuTree",
					Handler: MenuGetBaseMenuTreeHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/addMenuAuthority",
					Handler: MenuAddMenuAuthorityHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/getMenuAuthority",
					Handler: MenuGetMenuAuthorityHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/deleteBaseMenu",
					Handler: MenuDeleteBaseMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/updateBaseMenu",
					Handler: MenuUpdateBaseMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/getBaseMenuById",
					Handler: MenuGetBaseMenuByIdHandler(serverCtx),
				},
			}...,
		),
	)
}
