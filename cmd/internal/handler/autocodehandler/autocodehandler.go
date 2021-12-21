package autocodehandler

import (
	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-template/cmd/internal/logic/autocode"
	"go-zero-template/cmd/internal/svc"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/utils"
	"net/http"
)

func DelSysHistoryHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AutoHistoryByID
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := autocode.NewAutoCodeLogic(r.Context(), ctx)
		resp, err := l.DelSysHistory(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func GetMetaHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AutoHistoryByID
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := autocode.NewAutoCodeLogic(r.Context(), ctx)
		resp, err := l.GetMeta(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func GetSysHistoryHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysAutoHistory
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := autocode.NewAutoCodeLogic(r.Context(), ctx)
		resp, err := l.GetSysHistory(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func RollBackHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AutoHistoryByID
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := autocode.NewAutoCodeLogic(r.Context(), ctx)
		resp, err := l.RollBack(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func PreviewTempHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AutoCodeStruct
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := autocode.NewAutoCodeLogic(r.Context(), ctx)
		resp, err := l.PreviewTemp(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func CreateTempHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AutoCodeStruct
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := autocode.NewAutoCodeLogic(r.Context(), ctx)
		resp, err := l.CreateTemp(req,w,r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}

}
func GetTablesHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.AutoCodeStruct
		//if err := utils.Bind(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		l := autocode.NewAutoCodeLogic(r.Context(), ctx)
		resp, err := l.GetTables()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func GetDBHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.AutoCodeStruct
		//if err := utils.Bind(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		l := autocode.NewAutoCodeLogic(r.Context(), ctx)
		resp, err := l.GetDB()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func GetColumnHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.AutoCodeStruct
		//if err := utils.Bind(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		l := autocode.NewAutoCodeLogic(r.Context(), ctx)
		resp, err := l.GetColumn(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}