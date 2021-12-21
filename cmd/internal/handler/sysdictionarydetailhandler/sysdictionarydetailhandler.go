package sysdictionarydetailhandler

import (
	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-template/cmd/internal/logic/sysdictionarydetail"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/utils"

	"go-zero-template/cmd/internal/svc"
	"net/http"
)

func CreateSysDictionaryDetailHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysDictionaryDetail
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := sysdictionarydetail.NewSysDictionaryDetailLogic(r.Context(), ctx)
		resp, err := l.CreateSysDictionaryDetail(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func DeleteSysDictionaryDetailHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysDictionaryDetail
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := sysdictionarydetail.NewSysDictionaryDetailLogic(r.Context(), ctx)
		resp, err := l.DeleteSysDictionaryDetail(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func UpdateSysDictionaryDetailHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysDictionaryDetail
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := sysdictionarydetail.NewSysDictionaryDetailLogic(r.Context(), ctx)
		resp, err := l.UpdateSysDictionaryDetail(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func FindSysDictionaryDetailHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysDictionaryDetail
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := sysdictionarydetail.NewSysDictionaryDetailLogic(r.Context(), ctx)
		resp, err := l.FindSysDictionaryDetail(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func GetSysDictionaryDetailListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysDictionaryDetailSearch
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := sysdictionarydetail.NewSysDictionaryDetailLogic(r.Context(), ctx)
		resp, err := l.GetSysDictionaryDetailList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}