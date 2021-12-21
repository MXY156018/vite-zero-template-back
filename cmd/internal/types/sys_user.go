package types

type Register struct {
	Username     string   `json:"userName"`
	Password     string   `json:"passWord"`
	NickName     string   `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg    string   `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
	AuthorityId  string   `json:"authorityId" gorm:"default:888"`
	AuthorityIds []string `json:"authorityIds"`
}
type ChangePasswordStruct struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}
type PageInfo struct {
	Page     int `json:"page,optional" form:"page"`     // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}
type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}
type SetUserAuth struct {
	AuthorityId string `json:"authorityId"` // 角色ID
}
type GetById struct {
	ID float64 `json:"id"` // 主键ID
}
type IdsReq struct {
	Ids []int `json:"ids"`
}
type Sysuser struct {
	UserInfo SysUser `json:"userInfo"`
}
