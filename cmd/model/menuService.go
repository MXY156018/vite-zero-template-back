package model

import (
	"errors"
	"fmt"
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/types"
	"gorm.io/gorm"
	"strconv"
)

type MenuService struct {
}

var MenuServiceApp = new(MenuService)

func (menuService *MenuService) GetMenuTree(authorityId string) (err error, menus []types.SysMenu) {
	err, menuTree := menuService.getMenuTreeMap(authorityId)
	fmt.Printf("%v\n", menuTree)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getChildrenList(&menus[i], menuTree)
	}
	return err, menus
}

func (menuService *MenuService) getMenuTreeMap(authorityId string) (err error, treeMap map[string][]types.SysMenu) {
	var allMenus []types.SysMenu
	treeMap = make(map[string][]types.SysMenu)
	err = global.GVA_DB.Table("authority_menu").Where("authority_id = ?", authorityId).Order("sort").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}
func (menuService *MenuService) getChildrenList(menu *types.SysMenu, treeMap map[string][]types.SysMenu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}
func (menuService *MenuService) GetInfoList() (err error, list interface{}, total int64) {
	var menuList []types.SysBaseMenu
	err, treeMap := menuService.getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = menuService.getBaseChildrenList(&menuList[i], treeMap)
	}
	return err, menuList, total
}
func (menuService *MenuService) getBaseMenuTreeMap() (err error, treeMap map[string][]types.SysBaseMenu) {
	var allMenus []types.SysBaseMenu
	treeMap = make(map[string][]types.SysBaseMenu)
	err = global.GVA_DB.Order("sort").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}
func (menuService *MenuService) getBaseChildrenList(menu *types.SysBaseMenu, treeMap map[string][]types.SysBaseMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}
func (menuService *MenuService) AddBaseMenu(menu types.SysBaseMenu) error {
	if !errors.Is(global.GVA_DB.Where("name = ?", menu.Name).First(&types.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	return global.GVA_DB.Create(&menu).Error
}
func (menuService *MenuService) GetBaseMenuTree() (err error, menus []types.SysBaseMenu) {
	err, treeMap := menuService.getBaseMenuTreeMap()
	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getBaseChildrenList(&menus[i], treeMap)
	}
	return err, menus
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: AddMenuAuthority
//@description: 为角色增加menu树
//@param: menus []model.SysBaseMenu, authorityId string
//@return: err error
func (menuService *MenuService) AddMenuAuthority(menus []types.SysBaseMenu, authorityId string) (err error) {
	var auth types.SysAuthority
	auth.AuthorityId = authorityId
	auth.SysBaseMenus = menus
	err = AuthorityServiceApp.SetMenuAuthority(&auth)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMenuAuthority
//@description: 查看当前角色树
//@param: info *request.GetAuthorityId
//@return: err error, menus []model.SysMenu

func (menuService *MenuService) GetMenuAuthority(info *types.GetAuthorityId) (err error, menus []types.SysMenu) {
	err = global.GVA_DB.Table("authority_menu").Where("authority_id = ? ", info.AuthorityId).Order("sort").Find(&menus).Error
	//sql := "SELECT authority_menu.keep_alive,authority_menu.default_menu,authority_menu.created_at,authority_menu.updated_at,authority_menu.deleted_at,authority_menu.menu_level,authority_menu.parent_id,authority_menu.path,authority_menu.`name`,authority_menu.hidden,authority_menu.component,authority_menu.title,authority_menu.icon,authority_menu.sort,authority_menu.menu_id,authority_menu.authority_id FROM authority_menu WHERE authority_menu.authority_id = ? ORDER BY authority_menu.sort ASC"
	//err = global.GVA_DB.Raw(sql, authorityId).Scan(&menus).Error
	return err, menus
}
