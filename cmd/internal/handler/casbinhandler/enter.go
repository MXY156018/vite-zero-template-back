package casbinhandler

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
					Path:    "/casbin/updateCasbin",
					Handler: UpdateCasbinHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/casbin/getPolicyPathByAuthorityId",
					Handler: GetPolicyPathByAuthorityIdHandler(serverCtx),
				},
			}...,
		),
	)
}
