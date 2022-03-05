package customerhandler

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
					Path:    "/customer/customer",
					Handler: CreateExaCustomerHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/customer/customer",
					Handler: UpdateExaCustomerHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/customer/customer",
					Handler: DeleteExaCustomerHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/customer/customer",
					Handler: GetExaCustomerHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/customer/customerList",
					Handler: GetExaCustomerListHandler(serverCtx),
				},
			}...,
		),
	)
}
