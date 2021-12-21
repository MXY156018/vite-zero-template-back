package basehandler

import (
	"github.com/tal-tech/go-zero/rest"
	"go-zero-template/cmd/internal/svc"
	"net/http"
)

func RegisterHandlersAutocode(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/base/captcha",
				Handler: CaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/base/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/jwt/jsonInBlacklist",
				Handler: JsonInBlacklistHandler(serverCtx),
			},
			//{
			//	Method:  http.MethodGet,
			//	Path:    "/resource/page/prefix",
			//	Handler: StaticHandler("/prefix", "/page"),
			//},
		},
	)
}
