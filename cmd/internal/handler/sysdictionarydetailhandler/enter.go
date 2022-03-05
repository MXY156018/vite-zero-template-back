package sysdictionarydetailhandler

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
					Path:    "/sysDictionaryDetail/createSysDictionaryDetail",
					Handler: CreateSysDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/sysDictionaryDetail/deleteSysDictionaryDetail",
					Handler: DeleteSysDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/sysDictionaryDetail/updateSysDictionaryDetail",
					Handler: UpdateSysDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/sysDictionaryDetail/findSysDictionaryDetail",
					Handler: FindSysDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/sysDictionaryDetail/getSysDictionaryDetailList",
					Handler: GetSysDictionaryDetailListHandler(serverCtx),
				},
			}...,
		),
	)
}
