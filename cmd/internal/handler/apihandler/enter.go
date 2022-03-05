package apihandler

import (
	"github.com/zeromicro/go-zero/rest"
	"go-zero-template/cmd/internal/svc"
	"net/http"
)

func RegisterHandlersAutocode(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/createApi",
				Handler: CreateApiHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/deleteApi",
				Handler: DeleteApiHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/getApiList",
				Handler: GetApiListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/getApiById",
				Handler: GetApiByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/updateApi",
				Handler: UpdateApiHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/getAllApis",
				Handler: GetAllApisHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/deleteApisByIds",
				Handler: DeleteApisByIdsHandler(serverCtx),
			},
		},
	)
}
