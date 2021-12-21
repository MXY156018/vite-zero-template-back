package systemhandler

import (
	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-template/cmd/internal/logic/system"
	"go-zero-template/cmd/internal/svc"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/utils"
	"net/http"
)

func GetSystemConfigHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := system.NewSystemLogic(r.Context(), ctx)
		resp, err := l.GetSystemConfig()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func SetSystemConfigHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.System
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := system.NewSystemLogic(r.Context(), ctx)
		resp, err := l.SetSystemConfig(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func GetServerInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.System
		//if err := utils.Bind(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		l := system.NewSystemLogic(r.Context(), ctx)
		resp, err := l.GetServerInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func ReloadSystemHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.System
		//if err := utils.Bind(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		l := system.NewSystemLogic(r.Context(), ctx)
		resp, err := l.ReloadSystem()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}