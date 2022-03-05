package base

import (
	"context"
	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/svc"
	"go-zero-template/cmd/internal/types"
	"go.uber.org/zap"
)

type CaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) CaptchaLogic {
	return CaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var store = base64Captcha.DefaultMemStore

func (l *CaptchaLogic) Captcha() (*types.Result, error) {
	driver := base64Captcha.NewDriverDigit(global.GVA_CONFIG.Captcha.ImgHeight, global.GVA_CONFIG.Captcha.ImgWidth, global.GVA_CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.GVA_LOG.Error("验证码获取失败!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "验证码获取失败",
		}, nil
	} else {
		return &types.Result{
			Code: 0,
			Data: types.SysCaptchaResponse{
				CaptchaId: id,
				PicPath:   b64s,
			},
			Msg: "验证码获取成功",
		}, nil
	}

}
