package sysdictionaryhandler

import (
	"github.com/tal-tech/go-zero/rest"
	"go-zero-template/cmd/internal/svc"
	"net/http"
)
// 字典管理
func RegisterHandlersAutocode(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Jwt, serverCtx.Casbin, serverCtx.OperateRecord},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/sysDictionary/createSysDictionary",
					Handler: CreateSysDictionaryHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/sysDictionary/deleteSysDictionary",
					Handler: DeleteSysDictionaryHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/sysDictionary/updateSysDictionary",
					Handler: UpdateSysDictionaryHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/sysDictionary/findSysDictionary",
					Handler: FindSysDictionaryHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/sysDictionary/getSysDictionaryList",
					Handler: GetSysDictionaryListHandler(serverCtx),
				},
			}...,
		),
	)
}
