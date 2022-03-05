package sysoperationrecord

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/svc"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/model"
	"go-zero-template/cmd/utils"
	"go.uber.org/zap"
)

type SysOperationRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysOperationRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) SysOperationRecordLogic {
	return SysOperationRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (u *SysOperationRecordLogic) CreateSysOperationRecord(req types.SysOperationRecord) (*types.Result, error) {
	if err := model.OperationRecordServiceApp.CreateSysOperationRecord(req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "创建失败"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "创建成功"}, nil
	}
}
func (u *SysOperationRecordLogic) DeleteSysOperationRecord(req types.SysOperationRecord) (*types.Result, error) {
	if err := model.OperationRecordServiceApp.DeleteSysOperationRecord(req); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "删除失败"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "删除成功"}, nil
	}
}
func (u *SysOperationRecordLogic) DeleteSysOperationRecordByIds(req types.IdsReq) (*types.Result, error) {
	if err := model.OperationRecordServiceApp.DeleteSysOperationRecordByIds(req); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "批量删除失败"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "批量删除成功"}, nil
	}
}
func (u *SysOperationRecordLogic) FindSysOperationRecord(req types.SysOperationRecord) (*types.Result, error) {
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	if err, resysOperationRecord := model.OperationRecordServiceApp.GetSysOperationRecord(req.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "查询失败"}, nil
	} else {
		type ResysOperationRecord struct {
			ResysOperationRecord types.SysOperationRecord `json:"resysOperationRecord"`
		}
		return &types.Result{Code: 0, Msg: "查询成功", Data: ResysOperationRecord{ResysOperationRecord: resysOperationRecord}}, nil
	}
}
func (u *SysOperationRecordLogic) GetSysOperationRecordList(req types.SysOperationRecordSearch) (*types.Result, error) {

	if err, list, total := model.OperationRecordServiceApp.GetSysOperationRecordInfoList(req); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "获取失败"}, nil
	} else {

		return &types.Result{Code: 0, Msg: "查询成功", Data: types.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}}, nil
	}
}
