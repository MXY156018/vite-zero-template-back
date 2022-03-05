package jwthandler

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
				Path:    "/jwt/jsonInBlacklist",
				Handler: JsonInBlacklistHandler(serverCtx),
			},
		},
	)
}
