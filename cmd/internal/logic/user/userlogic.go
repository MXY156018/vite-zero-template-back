package user

import (
	"context"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/middleware"
	"go-zero-template/cmd/internal/svc"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/model"
	"go-zero-template/cmd/utils"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserLogic {
	return UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//UserRegister 用户注册账号
func (u *UserLogic) UserRegister(req types.Register) (*types.Result, error) {
	if err := utils.Verify(req, utils.RegisterVerify); err != nil {
		return &types.Result{
			Code: 7,
			Msg:  err.Error(),
		}, nil
	}
	var authorities []types.SysAuthority
	for _, v := range req.AuthorityIds {
		authorities = append(authorities, types.SysAuthority{
			AuthorityId: v,
		})
	}
	user := &types.SysUser{Username: req.Username, NickName: req.NickName, Password: req.Password, HeaderImg: req.HeaderImg, AuthorityId: req.AuthorityId, Authorities: authorities}
	err, userReturn := model.Register(*user)
	if err != nil {
		return &types.Result{
			Code: 7,
			Msg:  "注册失败",
			Data: userReturn,
		}, nil
	}
	return &types.Result{
		Code: 0,
		Msg:  "注册成功",
		Data: userReturn,
	}, nil

}

//UserChangePassword 用户修改密码
func (u *UserLogic) UserChangePassword(req types.ChangePasswordStruct) (*types.Result, error) {
	if err := utils.Verify(req, utils.ChangePasswordVerify); err != nil {
		return &types.Result{
			Code: 7,
			Msg:  err.Error(),
		}, nil
	}
	b := &types.SysUser{Username: req.Username, Password: req.Password}
	if err, _ := model.ChangePassword(b, req.NewPassword); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "修改失败,原密码与当前账户不符",
		}, nil
	}
	return &types.Result{
		Code: 0,
		Msg:  "修改成功",
	}, nil
}

//UserList 分页获取用户列表
func (u *UserLogic) UserList(req types.PageInfo) (*types.Result, error) {
	if err := utils.Verify(req, utils.PageInfoVerify); err != nil {
		return &types.Result{
			Code: 7,
			Msg:  err.Error(),
		}, nil
	}
	err, list, total := model.GetUserInfoList(req)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  err.Error(),
		}, nil
	}
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

//UserSetUserAuthority 修改用户权限
func (u *UserLogic) UserSetUserAuthority(req types.SetUserAuth, r *http.Request) (*types.Result, error) {
	if UserVerifyErr := utils.Verify(req, utils.SetUserAuthorityVerify); UserVerifyErr != nil {
		return &types.Result{
			Code: 7,
			Msg:  UserVerifyErr.Error(),
		}, nil
	}
	userID := utils.GetUserID(r)
	uuid := utils.GetUserUuid(r)
	if err := model.SetUserAuthority(userID, uuid, req.AuthorityId); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  err.Error(),
		}, nil
	} else {
		claims := utils.GetUserInfo(r)
		j := &middleware.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
		println(req.AuthorityId)
		println(claims.AuthorityId)
		fmt.Printf("%v\n",claims)
		claims.AuthorityId = req.AuthorityId
		if token, err := j.CreateToken(*claims); err != nil {
			global.GVA_LOG.Error("修改失败!", zap.Any("err", err))
			return &types.Result{
				Code: 7,
				Msg:  err.Error(),
			}, nil
		} else {
			r.Header.Set("new-token", token)
			r.Header.Set("net-expires-at", strconv.FormatInt(claims.ExpiresAt, 10))
			return &types.Result{
				Code: 0,
				Msg:  "修改成功",
			}, nil
		}
	}
}

//DeleteUser 删除用户
func (u *UserLogic) DeleteUser(req types.GetById, r *http.Request) (*types.Result, error) {
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		return &types.Result{
			Code: 7,
			Msg:  err.Error(),
		}, nil
	}
	jwtId := utils.GetUserID(r)
	if jwtId == uint(req.ID) {
		return &types.Result{
			Code: 7,
			Msg:  "删除失败，自杀失败",
		}, nil
	}
	if err := model.DeleteUser(req.ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "删除失败",
		}, nil
	}
	return &types.Result{
		Code: 0,
		Msg:  "删除成功",
	}, nil
}

//SetUserInfo 设置用户信息
func (u *UserLogic) SetUserInfo(req types.SysUser) (*types.Result, error) {
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	err, ReqUser := model.SetUserInfo(req)
	if err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "设置失败"}, nil
	}
	return &types.Result{Code: 0, Msg: "设置成功", Data: ReqUser}, nil
}

//SetUserAuthorities 设置用户权限组
func (u *UserLogic) SetUserAuthorities(req types.SetUserAuthorities) (*types.Result, error) {
	if err := model.SetUserAuthorities(req.ID, req.AuthorityIds); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "修改失败"}, nil
	}
	return &types.Result{Code: 0, Msg: "修改成功"}, nil
}

//GetUserInfo 获取用户信息
func (u *UserLogic) GetUserInfo(r *http.Request) (*types.Result, error) {
	uuid := utils.GetUserUuid(r)
	err, ReqUser := model.GetUserInfo(uuid)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "获取失败"}, nil
	}
	return &types.Result{Code: 0, Msg: "获取成功", Data: types.Sysuser{UserInfo: ReqUser}}, nil
}
