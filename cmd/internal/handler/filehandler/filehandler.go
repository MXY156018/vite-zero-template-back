package filehandler

import (
	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-template/cmd/internal/logic/file"
	"go-zero-template/cmd/internal/svc"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/utils"
	"net/http"
)

func UploadFileHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.Register
		//if err := utils.Bind(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		l := file.NewFileLogic(r.Context(), ctx)
		resp, err := l.UploadFile(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func GetFileListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageInfo
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := file.NewFileLogic(r.Context(), ctx)
		resp, err := l.GetFileList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func DeleteFileHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExaFileUploadAndDownload
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := file.NewFileLogic(r.Context(), ctx)
		resp, err := l.DeleteFile(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func BreakpointContinueHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.ExaFileUploadAndDownload
		//if err := utils.Bind(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		l := file.NewFileLogic(r.Context(), ctx)
		resp, err := l.BreakpointContinue(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func FindFileHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.ExaFileUploadAndDownload
		//if err := utils.Bind(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		l := file.NewFileLogic(r.Context(), ctx)
		resp, err := l.FindFile(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func BreakpointContinueFinishHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.ExaFileUploadAndDownload
		//if err := utils.Bind(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		l := file.NewFileLogic(r.Context(), ctx)
		resp, err := l.BreakpointContinueFinish(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func RemoveChunkHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.ExaFileUploadAndDownload
		//if err := utils.Bind(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		l := file.NewFileLogic(r.Context(), ctx)
		resp, err := l.RemoveChunk(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
