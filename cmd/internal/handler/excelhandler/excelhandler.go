package excelhandler

import (
	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-template/cmd/internal/logic/excel"
	"go-zero-template/cmd/internal/svc"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/utils"
	"net/http"
)

func ImportExcelHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.Register
		//if err := utils.Bind(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		l := excel.NewExcelLogic(r.Context(), ctx)
		resp, err := l.ImportExcel(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func LoadExcelHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.Register
		//if err := utils.Bind(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		l := excel.NewExcelLogic(r.Context(), ctx)
		resp, err := l.LoadExcel()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func ExportExcelHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExcelInfo
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := excel.NewExcelLogic(r.Context(), ctx)
		resp, err := l.ExportExcel(req,w,r)

		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func DownloadTemplateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.ExcelInfo
		//if err := utils.Bind(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		l := excel.NewExcelLogic(r.Context(), ctx)
		resp, err := l.DownloadTemplate(w,r)

		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
