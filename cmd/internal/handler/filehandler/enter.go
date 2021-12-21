package filehandler

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
					Path:    "/fileUploadAndDownload/upload",
					Handler: UploadFileHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/fileUploadAndDownload/getFileList",
					Handler: GetFileListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/fileUploadAndDownload/deleteFile",
					Handler: DeleteFileHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/fileUploadAndDownload/breakpointContinue",
					Handler: BreakpointContinueHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/fileUploadAndDownload/findFile",
					Handler: FindFileHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/fileUploadAndDownload/breakpointContinueFinish",
					Handler: BreakpointContinueFinishHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/fileUploadAndDownload/removeChunk",
					Handler: RemoveChunkHandler(serverCtx),
				},
			}...,
		),
	)
}
