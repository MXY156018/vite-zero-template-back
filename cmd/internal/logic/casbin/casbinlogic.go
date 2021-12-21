package casbin

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/svc"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/model"
	"go-zero-template/cmd/utils"
	"go.uber.org/zap"
)

type CasbinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCasbinLogic(ctx context.Context, svcCtx *svc.ServiceContext) CasbinLogic {
	return CasbinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (c *CasbinLogic) UpdateCasbin(req types.CasbinInReceive) (*types.Result, error) {
	if err := utils.Verify(req, utils.AuthorityIdVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	if err := model.CasbinServiceApp.UpdateCasbin(req.AuthorityId, req.CasbinInfos); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "更新失败!"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "更新成功!"}, nil
	}
}
func (c *CasbinLogic) GetPolicyPathByAuthorityId(req types.CasbinSearch) (*types.Result, error) {
	if err := utils.Verify(req, utils.AuthorityIdVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	paths := model.CasbinServiceApp.GetPolicyPathByAuthorityId(req.AuthorityId)
	return &types.Result{
		Code: 0,
		Msg:  "获取成功",
		Data: types.PolicyPathResponse{
			Paths: paths,
		},
	}, nil
}
