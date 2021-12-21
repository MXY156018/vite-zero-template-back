package types

type SysMenu struct {
	SysBaseMenu
	MenuId      string                 `json:"menuId" gorm:"comment:菜单ID"`
	AuthorityId string                 `json:"-" gorm:"comment:角色ID"`
	Children    []SysMenu              `json:"children" gorm:"-"`
	Parameters  []SysBaseMenuParameter `json:"parameters" gorm:"foreignKey:SysBaseMenuID;references:MenuId"`
}
type SysMenusResponse struct {
	Menus []SysMenu `json:"menus"`
}
type SysMenuResponse struct {
	Menu []SysMenu `json:"menu"`
}
type SysBaseMenusResponse struct {
	Menus []SysBaseMenu `json:"menus"`
}

type SysBaseMenuResponse struct {
	Menu SysBaseMenu `json:"menu"`
}
type AddMenuAuthorityInfo struct {
	Menus       []SysBaseMenu `json:"menus"`
	AuthorityId string `json:"authorityId"` // 角色ID
}
type GetAuthorityId struct {
	AuthorityId string `json:"authorityId"` // 角色ID
}
