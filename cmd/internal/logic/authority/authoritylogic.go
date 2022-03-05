package authority

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

type AuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) AuthorityLogic {
	return AuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//Authority 创建角色
func (u *AuthorityLogic) Authority(req types.SysAuthority) (*types.Result, error) {
	if err := utils.Verify(req, utils.AuthorityVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	if err, authBack := model.AuthorityServiceApp.CreateAuthority(req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "创建失败!"}, nil
	} else {
		_ = model.MenuServiceApp.AddMenuAuthority(types.DefaultMenu(), req.AuthorityId)
		_ = model.CasbinServiceApp.UpdateCasbin(req.AuthorityId, types.DefaultCasbin())
		return &types.Result{Code: 0, Msg: "创建成功",
			Data: types.SysAuthorityResponse{
				Authority: authBack,
			},
		}, nil
	}
}

//DeleteAuthority 删除角色
func (u *AuthorityLogic) DeleteAuthority(req types.SysAuthority) (*types.Result, error) {
	if err := model.AuthorityServiceApp.DeleteAuthority(&req); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "删除失败!"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "删除成功!"}, nil
	}
}

//UpdateAuthority 更新角色信息
func (u *AuthorityLogic) UpdateAuthority(req types.SysAuthority) (*types.Result, error) {
	if err := utils.Verify(req, utils.AuthorityVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	if err, authority := model.AuthorityServiceApp.UpdateAuthority(req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "更新失败!"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "创建成功",
			Data: types.SysAuthorityResponse{
				Authority: authority,
			},
		}, nil
	}
}
func (u *AuthorityLogic) CopyAuthority(req types.SysAuthorityCopyResponse) (*types.Result, error) {
	if err := utils.Verify(req, utils.OldAuthorityVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	if err := utils.Verify(req.Authority, utils.AuthorityVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	if err, authBack := model.AuthorityServiceApp.CopyAuthority(req); err != nil {
		global.GVA_LOG.Error("拷贝失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "拷贝失败!"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "拷贝成功",
			Data: types.SysAuthorityResponse{
				Authority: authBack,
			},
		}, nil
	}
}
func (u *AuthorityLogic) GetAuthorityList(req types.PageInfo) (*types.Result, error) {
	if err := utils.Verify(req, utils.PageInfoVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}

	if err, list, total := model.AuthorityServiceApp.GetAuthorityInfoList(req); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "获取失败!"}, nil
	} else {
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
func (u *AuthorityLogic) SetDataAuthority(req types.SysAuthority) (*types.Result, error) {
	if err := utils.Verify(req, utils.AuthorityIdVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}

	if err := model.AuthorityServiceApp.SetDataAuthority(req); err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "设置失败!"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "设置成功"}, nil
	}
}
