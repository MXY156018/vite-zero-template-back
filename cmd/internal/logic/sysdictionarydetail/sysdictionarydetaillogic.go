package sysdictionarydetail

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/svc"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/model"
	"go.uber.org/zap"
)

type SysDictionaryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysDictionaryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) SysDictionaryDetailLogic {
	return SysDictionaryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//CreateSysDictionaryDetail 新建SysDictionaryDetail
func (s *SysDictionaryDetailLogic) CreateSysDictionaryDetail(req types.SysDictionaryDetail) (*types.Result, error) {
	if err := model.DictionaryDetailServiceApp.CreateSysDictionaryDetail(req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "创建失败!"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "创建成功"}, nil
	}
}

//DeleteSysDictionaryDetail 删除SysDictionaryDetail
func (s *SysDictionaryDetailLogic) DeleteSysDictionaryDetail(req types.SysDictionaryDetail) (*types.Result, error) {
	if err := model.DictionaryDetailServiceApp.DeleteSysDictionaryDetail(req); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "删除失败!"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "删除成功"}, nil
	}
}

//UpdateSysDictionaryDetail 更新SysDictionaryDetail
func (s *SysDictionaryDetailLogic) UpdateSysDictionaryDetail(req types.SysDictionaryDetail) (*types.Result, error) {
	if err := model.DictionaryDetailServiceApp.UpdateSysDictionaryDetail(&req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "更新失败!"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "更新成功"}, nil
	}
}

//FindSysDictionaryDetail 用id查询SysDictionaryDetail
func (s *SysDictionaryDetailLogic) FindSysDictionaryDetail(req types.SysDictionaryDetail) (*types.Result, error) {
	if err, resysDictionaryDetail := model.DictionaryDetailServiceApp.GetSysDictionaryDetail(req.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "查询失败!"}, nil
	} else {
		type DictionaryDetail struct {
			ResysDictionaryDetail types.SysDictionaryDetail `json:"resysDictionaryDetail"`
		}
		return &types.Result{Code: 0, Msg: "更新成功", Data: DictionaryDetail{ResysDictionaryDetail: resysDictionaryDetail}}, nil
	}
}

//GetSysDictionaryDetailList 分页获取SysDictionaryDetail列表
func (s *SysDictionaryDetailLogic) GetSysDictionaryDetailList(req types.SysDictionaryDetailSearch) (*types.Result, error) {
	if err, list, total := model.DictionaryDetailServiceApp.GetSysDictionaryDetailInfoList(req); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "获取失败!"}, nil
	} else {
		type DictionaryDetail struct {
			ResysDictionaryDetail types.SysDictionaryDetail `json:"resysDictionaryDetail"`
		}
		return &types.Result{Code: 0, Msg: "获取成功",
			Data: types.PageResult{
				List:     list,
				Total:    total,
				Page:     req.Page,
				PageSize: req.PageSize,
			},
		}, nil
	}
}
