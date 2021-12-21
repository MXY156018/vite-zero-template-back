package system

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/svc"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/model"
	"go.uber.org/zap"
)

type SystemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemLogic(ctx context.Context, svcCtx *svc.ServiceContext) SystemLogic {
	return SystemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//GetSystemConfig 获取配置文件内容
func (a *SystemLogic) GetSystemConfig() (*types.Result, error) {
	if err, config := model.SystemConfigServiceApp.GetSystemConfig(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "获取失败"}, nil
	} else {
		return &types.Result{
			Code: 0,
			Msg:  "获取成功",
			Data: types.SysConfigResponse{
				Config: config,
			},
		}, nil
	}
}

//SetSystemConfig 设置配置文件内容
func (a *SystemLogic) SetSystemConfig(req types.System) (*types.Result, error) {
	if err := model.SystemConfigServiceApp.SetSystemConfig(req); err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "设置失败"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "设置成功"}, nil
	}
}

//GetServerInfo 获取服务器信息
func (a *SystemLogic) GetServerInfo() (*types.Result, error) {
	if server, err := model.SystemConfigServiceApp.GetServerInfo(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "获取失败"}, nil
	} else {
		return &types.Result{
			Code: 0,
			Msg:  "获取成功",
			Data: types.Server1{
				ServerInfo: server,
			},
		}, nil
	}
}
func (a *SystemLogic) ReloadSystem() (*types.Result, error) {
	err := model.SystemConfigServiceApp.Reload()
	if err != nil {
		global.GVA_LOG.Error("重启系统失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "重启系统失败"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "重启系统成功"}, nil
	}
}
