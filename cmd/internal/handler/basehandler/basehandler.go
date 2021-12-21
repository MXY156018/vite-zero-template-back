package basehandler

import (
	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-template/cmd/internal/logic/base"
	"go-zero-template/cmd/internal/svc"
	"net/http"
)

func CaptchaHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := base.NewCaptchaLogic(r.Context(), ctx)
		resp, err := l.Captcha()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
