package sysoperationrecordhandler

import (
	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-template/cmd/internal/logic/sysoperationrecord"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/utils"

	"go-zero-template/cmd/internal/svc"
	"net/http"
)

func CreateSysOperationRecordHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysOperationRecord
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := sysoperationrecord.NewSysOperationRecordLogic(r.Context(), ctx)
		resp, err := l.CreateSysOperationRecord(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func DeleteSysOperationRecordHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysOperationRecord
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := sysoperationrecord.NewSysOperationRecordLogic(r.Context(), ctx)
		resp, err := l.DeleteSysOperationRecord(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func DeleteSysOperationRecordByIdsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdsReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := sysoperationrecord.NewSysOperationRecordLogic(r.Context(), ctx)
		resp, err := l.DeleteSysOperationRecordByIds(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func FindSysOperationRecordHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysOperationRecord
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := sysoperationrecord.NewSysOperationRecordLogic(r.Context(), ctx)
		resp, err := l.FindSysOperationRecord(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func GetSysOperationRecordListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysOperationRecordSearch
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := sysoperationrecord.NewSysOperationRecordLogic(r.Context(), ctx)
		resp, err := l.GetSysOperationRecordList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
