package api

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/svc"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/model"
	"go-zero-template/cmd/utils"
	"go.uber.org/zap"
)

type ApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) ApiLogic {
	return ApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//CreateApi 新增api
func (a *ApiLogic) CreateApi(req types.SysApi) (*types.Result, error) {
	if err := utils.Verify(req, utils.ApiVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	if err := model.ApiServiceApp.CreateApi(req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "创建失败"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "创建成功"}, nil
	}
}

//DeleteApi 删除api
func (a *ApiLogic) DeleteApi(req types.SysApi) (*types.Result, error) {
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	if err := model.ApiServiceApp.DeleteApi(req); err != nil {
		global.GVA_LOG.Error("删除api失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "删除失败"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "删除成功"}, nil
	}
}

//GetApiList 获取apilist
func (a *ApiLogic) GetApiList(req types.SearchApiParams) (*types.Result, error) {
	fmt.Sprintf("%v\n", req)
	if err := utils.Verify(req.PageInfo, utils.PageInfoVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	err, list, total := model.ApiServiceApp.GetAPIInfoList(req.SysApi, req.PageInfo, req.OrderKey, req.Desc)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "获取失败"}, nil
	} else {
		return &types.Result{
			Code: 0,
			Msg:  "获取成功",
			Data: types.PageResult{
				List:     list,
				Total:    total,
				Page:     req.Page,
				PageSize: req.PageSize,
			},
		}, nil
	}
}

//GetApiById 获取单条Api消息
func (a *ApiLogic) GetApiById(req types.GetById) (*types.Result, error) {
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	err, api := model.ApiServiceApp.GetApiById(req.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "获取失败"}, nil
	} else {
		return &types.Result{
			Code: 0,
			Msg:  "获取成功",
			Data: types.SysAPIResponse{
				Api: api,
			},
		}, nil
	}
}

//UpdateApi 更新api
func (a *ApiLogic) UpdateApi(req types.SysApi) (*types.Result, error) {
	if err := utils.Verify(req, utils.ApiVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	if err := model.ApiServiceApp.UpdateApi(req); err != nil {
		global.GVA_LOG.Error("修改api失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "修改api失败"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "修改api成功"}, nil
	}
}

//GetAllApis 获取所有api
func (a *ApiLogic) GetAllApis() (*types.Result, error) {

	if err, apis := model.ApiServiceApp.GetAllApis(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "获取失败"}, nil
	} else {
		return &types.Result{
			Code: 0,
			Msg:  "获取成功",
			Data: types.SysAPIListResponse{
				Apis: apis,
			},
		}, nil
	}
}
func (a *ApiLogic) DeleteApisByIds(req types.IdsReq) (*types.Result, error) {
	if err := model.ApiServiceApp.DeleteApisByIds(req); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "删除失败"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "删除成功"}, nil
	}
}
