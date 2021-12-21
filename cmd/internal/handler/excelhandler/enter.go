package excelhandler

import (
	"github.com/tal-tech/go-zero/rest"
	"go-zero-template/cmd/internal/svc"
	"net/http"
)

func RegisterHandlersAutocode(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Jwt, serverCtx.Casbin},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/excel/importExcel",
					Handler: ImportExcelHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/excel/loadExcel",
					Handler: LoadExcelHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/excel/exportExcel",
					Handler: ExportExcelHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/excel/downloadTemplate",
					Handler: DownloadTemplateHandler(serverCtx),
				},
			}...,
		),
	)
}
