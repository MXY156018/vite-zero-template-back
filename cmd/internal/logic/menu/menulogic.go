package menu

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/svc"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/model"
	"go-zero-template/cmd/utils"
	"go.uber.org/zap"
	"net/http"
)

type MenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) MenuLogic {
	return MenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//GetMenu 获取菜单树
func (u *MenuLogic) GetMenu(r *http.Request) (*types.Result, error) {
	println(utils.GetUserAuthorityId(r))
	if err, menus := model.MenuServiceApp.GetMenuTree(utils.GetUserAuthorityId(r)); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "获取失败"}, nil
	} else {
		if menus == nil {
			menus = []types.SysMenu{}
		}
		return &types.Result{
			Code: 0,
			Msg:  "获取成功",
			Data: types.SysMenusResponse{
				Menus: menus,
			},
		}, nil
	}
}

//GetMenuList 分页获取基础menu列表
func (u *MenuLogic) GetMenuList(req types.PageInfo) (*types.Result, error) {
	if err := utils.Verify(req, utils.PageInfoVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	err, menuList, total := model.MenuServiceApp.GetInfoList()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "获取失败"}, nil
	} else {
		return &types.Result{
			Code: 0,
			Msg:  "获取成功",
			Data: types.PageResult{
				List:     menuList,
				Total:    total,
				Page:     req.Page,
				PageSize: req.PageSize,
			},
		}, nil
	}
}

//AddBaseMenu 新增菜单
func (u *MenuLogic) AddBaseMenu(req types.SysBaseMenu) (*types.Result, error) {
	if err := utils.Verify(req, utils.MenuVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	if err := utils.Verify(req.Meta, utils.MenuMetaVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	if err := model.MenuServiceApp.AddBaseMenu(req); err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "添加失败"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "添加成功"}, nil
	}
}

//GetBaseMenuTree 获取用户动态路由
func (u *MenuLogic) GetBaseMenuTree() (*types.Result, error) {
	if err, menus := model.MenuServiceApp.GetBaseMenuTree(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "获取失败"}, nil
	} else {
		return &types.Result{
			Code: 0,
			Msg:  "获取成功",
			Data: types.SysBaseMenusResponse{
				Menus: menus,
			},
		}, nil
	}
}

//AddMenuAuthority 增加menu和角色关联关系
func (u *MenuLogic) AddMenuAuthority(req types.AddMenuAuthorityInfo) (*types.Result, error) {
	if err := utils.Verify(req, utils.AuthorityIdVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	if err := model.MenuServiceApp.AddMenuAuthority(req.Menus, req.AuthorityId); err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "添加失败"}, nil
	}
	return &types.Result{Code: 0, Msg: "添加成功"}, nil
}

// @Tags AuthorityMenu
// @Summary 获取指定角色menu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetAuthorityId true "角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/GetMenuAuthority [post]
func (u *MenuLogic) GetMenuAuthority(req types.GetAuthorityId) (*types.Result, error) {
	if err := utils.Verify(req, utils.AuthorityIdVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	if err, menus := model.MenuServiceApp.GetMenuAuthority(&req); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "获取失败",
			Data: types.SysMenusResponse{
				Menus: menus,
			},
		}, nil
	} else {
		return &types.Result{
			Code: 0,
			Msg:  "获取成功",
			Data: types.SysMenusResponse{
				Menus: menus,
			},
		}, nil
	}
}

// @Tags Menu
// @Summary 删除菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "菜单id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /menu/deleteBaseMenu [post]
func (u *MenuLogic) DeleteBaseMenu(req types.GetById) (*types.Result, error) {
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	if err := model.DeleteBaseMenu(req.ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "删除失败"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "删除成功"}, nil
	}
}

// @Tags Menu
// @Summary 更新菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysBaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /menu/updateBaseMenu [post]
func (u *MenuLogic) UpdateBaseMenu(req types.SysBaseMenu) (*types.Result, error) {
	if err := utils.Verify(req, utils.MenuVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	if err := utils.Verify(req.Meta, utils.MenuMetaVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	if err := model.UpdateBaseMenu(req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "更新失败"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "更新成功"}, nil
	}
}
func (u *MenuLogic) GetBaseMenuById(req types.GetById) (*types.Result, error) {
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	if err, menu := model.GetBaseMenuById(req.ID); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "获取失败"}, nil
	} else {
		return &types.Result{
			Code: 0,
			Msg:  "获取成功",
			Data: types.SysBaseMenuResponse{
				Menu: menu,
			},
		}, nil
	}
}
