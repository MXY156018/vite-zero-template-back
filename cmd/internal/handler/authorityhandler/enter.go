package authorityhandler

import (
	"github.com/zeromicro/go-zero/rest"
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
					Path:    "/authority/createAuthority",
					Handler: CreateAuthorityHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/authority/deleteAuthority",
					Handler: DeleteAuthorityHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/authority/updateAuthority",
					Handler: UpdateAuthorityHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/authority/copyAuthority",
					Handler: CopyAuthorityHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/authority/getAuthorityList",
					Handler: GetAuthorityListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/authority/setDataAuthority",
					Handler: SetDataAuthorityHandler(serverCtx),
				},
			}...,
		),
	)
}
