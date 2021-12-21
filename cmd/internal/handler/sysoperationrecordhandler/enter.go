package sysoperationrecordhandler

import (
	"github.com/tal-tech/go-zero/rest"
	"go-zero-template/cmd/internal/svc"
	"net/http"
)

func RegisterHandlersAutocode(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Jwt, serverCtx.Casbin, serverCtx.OperateRecord},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/sysOperationRecord/createSysOperationRecord",
					Handler: CreateSysOperationRecordHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/sysOperationRecord/deleteSysOperationRecord",
					Handler: DeleteSysOperationRecordHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/sysOperationRecord/deleteSysOperationRecordByIds",
					Handler: DeleteSysOperationRecordByIdsHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/sysOperationRecord/findSysOperationRecord",
					Handler: FindSysOperationRecordHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/sysOperationRecord/getSysOperationRecordList",
					Handler: GetSysOperationRecordListHandler(serverCtx),
				},
			}...,
		),
	)
}
