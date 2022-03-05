package sysdictionary

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

type SysDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) SysDictionaryLogic {
	return SysDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//CreateSysDictionary 创建SysDictionary
func (u *SysDictionaryLogic) CreateSysDictionary(req types.SysDictionary) (*types.Result, error) {
	if err := model.DictionaryServiceAPP.CreateSysDictionary(req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "创建失败"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "创建成功"}, nil
	}
}
func (u *SysDictionaryLogic) DeleteSysDictionary(req types.SysDictionary) (*types.Result, error) {
	if err := model.DictionaryServiceAPP.DeleteSysDictionary(req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "创建失败"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "创建成功"}, nil
	}
}
func (u *SysDictionaryLogic) UpdateSysDictionary(req types.SysDictionary) (*types.Result, error) {
	if err := model.DictionaryServiceAPP.UpdateSysDictionary(&req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "更新失败"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "更新成功"}, nil
	}
}
func (u *SysDictionaryLogic) FindSysDictionary(req types.SysDictionary) (*types.Result, error) {
	if err, sysDictionary := model.DictionaryServiceAPP.GetSysDictionary(req.Type, req.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "查询失败"}, nil
	} else {
		type ResysDictionary struct {
			ResysDictionary types.SysDictionary `json:"resysDictionary"`
		}
		return &types.Result{Code: 0, Msg: "查询成功", Data: ResysDictionary{ResysDictionary: sysDictionary}}, nil
	}
}
func (u *SysDictionaryLogic) GetSysDictionaryList(req types.SysDictionarySearch) (*types.Result, error) {
	if err := utils.Verify(req.PageInfo, utils.PageInfoVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	if err, list, total := model.DictionaryServiceAPP.GetSysDictionaryInfoList(req); err != nil {
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
