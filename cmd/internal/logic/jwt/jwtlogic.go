package jwt

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/svc"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/model"
	"go.uber.org/zap"
	"net/http"
)

type JwtLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJwtLogic(ctx context.Context, svcCtx *svc.ServiceContext) JwtLogic {
	return JwtLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//JsonInBlacklist jwt加入黑名单
func (a *JwtLogic) JsonInBlacklist(r *http.Request) (*types.Result, error) {
	token := r.Header.Get("x-token")
	jwt := types.JwtBlacklist{Jwt: token}
	if err := model.JwtServiceApp.JsonInBlacklist(jwt); err != nil {
		global.GVA_LOG.Error("jwt作废失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "jwt作废失败"}, nil
	}
	return &types.Result{Code: 7, Msg: "jwt作废成功"}, nil
}
