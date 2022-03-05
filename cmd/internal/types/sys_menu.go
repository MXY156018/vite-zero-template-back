package types

import "go-zero-template/cmd/global"

func DefaultMenu() []SysBaseMenu {
	return []SysBaseMenu{{
		GVA_MODEL: global.GVA_MODEL{ID: 1},
		ParentId:  "0",
		Path:      "dashboard",
		Name:      "dashboard",
		Component: "view/dashboard/index.vue",
		Sort:      1,
		Meta: Meta{
			Title: "仪表盘",
			Icon:  "setting",
		},
	}}
}
