package jwthandler

import (
	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-template/cmd/internal/logic/jwt"
	"go-zero-template/cmd/internal/svc"
	"net/http"
)

func JsonInBlacklistHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := jwt.NewJwtLogic(r.Context(), ctx)
		resp, err := l.JsonInBlacklist(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
