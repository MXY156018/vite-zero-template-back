package simpleuploadhandler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-template/cmd/internal/logic/simpleupload"
	"go-zero-template/cmd/internal/svc"
	"net/http"
)

func SimpleUploaderUploadHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.Register
		//if err := utils.Bind(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		l := simpleupload.NewSimpleUploadLogic(r.Context(), ctx)
		resp, err := l.SimpleUploaderUpload(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func CheckFileMd5Handler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.Register
		//if err := utils.Bind(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		l := simpleupload.NewSimpleUploadLogic(r.Context(), ctx)
		resp, err := l.CheckFileMd5(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func MergeFileMd5Handler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.Register
		//if err := utils.Bind(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		l := simpleupload.NewSimpleUploadLogic(r.Context(), ctx)
		resp, err := l.MergeFileMd5(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
