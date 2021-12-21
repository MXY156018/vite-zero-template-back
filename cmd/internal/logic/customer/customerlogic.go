package customer

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/svc"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/model"
	"go-zero-template/cmd/utils"
	"go.uber.org/zap"
	"net/http"
)

type CustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) CustomerLogic {
	return CustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//CreateExaCustomer 创建客户
func (c *CustomerLogic) CreateExaCustomer(req types.ExaCustomer, r *http.Request) (*types.Result, error) {
	if err := utils.Verify(req, utils.CustomerVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	req.SysUserID = utils.GetUserID(r)
	req.SysUserAuthorityID = utils.GetUserAuthorityId(r)
	if err := model.CustomerServiceApp.CreateExaCustomer(req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "创建失败!"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "创建成功!"}, nil
	}
}

//UpdateExaCustomer 更新客户信息
func (c *CustomerLogic) UpdateExaCustomer(req types.ExaCustomer) (*types.Result, error) {
	if err := utils.Verify(req.GVA_MODEL, utils.IdVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	if err := utils.Verify(req, utils.CustomerVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	if err := model.CustomerServiceApp.UpdateExaCustomer(&req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "更新失败!"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "更新成功!"}, nil
	}
}

//DeleteExaCustomer 删除客户
func (c *CustomerLogic) DeleteExaCustomer(req types.ExaCustomer) (*types.Result, error) {
	if err := utils.Verify(req.GVA_MODEL, utils.IdVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	if err := model.CustomerServiceApp.DeleteExaCustomer(req); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "删除失败!"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "删除成功!"}, nil
	}
}

//GetExaCustomer 获取单一客户信息
func (c *CustomerLogic) GetExaCustomer(req types.ExaCustomer) (*types.Result, error) {
	if err := utils.Verify(req.GVA_MODEL, utils.IdVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	if err, data := model.CustomerServiceApp.GetExaCustomer(req.ID); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "获取失败!"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "获取成功!", Data: types.ExaCustomerResponse{Customer: data}}, nil
	}
}
func (c *CustomerLogic) GetExaCustomerList(req types.PageInfo, r *http.Request) (*types.Result, error) {
	if err := utils.Verify(req, utils.PageInfoVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	err, customerList, total := model.CustomerServiceApp.GetCustomerInfoList(utils.GetUserAuthorityId(r), req)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "获取失败!"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "获取成功!",
			Data: types.PageResult{
				List:     customerList,
				Total:    total,
				Page:     req.Page,
				PageSize: req.PageSize,
			},
		}, nil
	}
}
