package types

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"go-zero-template/cmd/global"
	"time"
)

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg,omitempty"`
}
type SysCaptchaResponse struct {
	CaptchaId string `json:"captchaId"`
	PicPath   string `json:"picPath"`
}
type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}
type Meta struct {
	KeepAlive   bool   `json:"keepAlive" gorm:"comment:是否缓存"`           // 是否缓存
	DefaultMenu bool   `json:"defaultMenu" gorm:"comment:是否是基础路由（开发中）"` // 是否是基础路由（开发中）
	Title       string `json:"title" gorm:"comment:菜单名"`                // 菜单名
	Icon        string `json:"icon" gorm:"comment:菜单图标"`                // 菜单图标
	CloseTab    bool   `json:"closeTab" gorm:"comment:自动关闭tab"`         // 自动关闭tab
}
type SysBaseMenuParameter struct {
	global.GVA_MODEL
	SysBaseMenuID uint
	Type          string `json:"type" gorm:"comment:地址栏携带参数为params还是query"` // 地址栏携带参数为params还是query
	Key           string `json:"key" gorm:"comment:地址栏携带参数的key"`            // 地址栏携带参数的key
	Value         string `json:"value" gorm:"comment:地址栏携带参数的值"`            // 地址栏携带参数的值
}
type SysBaseMenu struct {
	global.GVA_MODEL
	MenuLevel     uint                              `json:"-"`
	ParentId      string                            `json:"parentId" gorm:"comment:父菜单ID"`     // 父菜单ID
	Path          string                            `json:"path" gorm:"comment:路由path"`        // 路由path
	Name          string                            `json:"name" gorm:"comment:路由name"`        // 路由name
	Hidden        bool                              `json:"hidden" gorm:"comment:是否在列表隐藏"`     // 是否在列表隐藏
	Component     string                            `json:"component" gorm:"comment:对应前端文件路径"` // 对应前端文件路径
	Sort          int                               `json:"sort" gorm:"comment:排序标记"`          // 排序标记
	Meta          `json:"meta" gorm:"comment:附加属性"` // 附加属性
	SysAuthoritys []SysAuthority                    `json:"authoritys" gorm:"many2many:sys_authority_menus;"`
	Children      []SysBaseMenu                     `json:"children" gorm:"-"`
	Parameters    []SysBaseMenuParameter            `json:"parameters"`
}

type SysAuthority struct {
	CreatedAt       time.Time      // 创建时间
	UpdatedAt       time.Time      // 更新时间
	DeletedAt       *time.Time     `sql:"index"`
	AuthorityId     string         `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	AuthorityName   string         `json:"authorityName" gorm:"comment:角色名"`                                    // 角色名
	ParentId        string         `json:"parentId" gorm:"comment:父角色ID"`                                       // 父角色ID
	DataAuthorityId []SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id"`
	Children        []SysAuthority `json:"children" gorm:"-"`
	SysBaseMenus    []SysBaseMenu  `json:"menus" gorm:"many2many:sys_authority_menus;"`
	DefaultRouter   string         `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"` // 默认菜单(默认dashboard)
}
type SysUser struct {
	global.GVA_MODEL
	UUID        uuid.UUID      `json:"uuid" gorm:"comment:用户UUID"`                                                    // 用户UUID
	Username    string         `json:"userName" gorm:"comment:用户登录名"`                                                 // 用户登录名
	Password    string         `json:"-"  gorm:"comment:用户登录密码"`                                                      // 用户登录密码
	NickName    string         `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                     // 用户昵称
	HeaderImg   string         `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"` // 用户头像
	Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId string         `json:"authorityId" gorm:"default:888;comment:用户角色ID"`     // 用户角色ID
	SideMode    string         `json:"sideMode" gorm:"default:dark;comment:用户角色ID"`       // 用户侧边主题
	ActiveColor string         `json:"activeColor" gorm:"default:#1890ff;comment:用户角色ID"` // 活跃颜色
	BaseColor   string         `json:"baseColor" gorm:"default:#fff;comment:用户角色ID"`      // 基础颜色
	Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
}
type CustomClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	NickName    string
	AuthorityId string
	BufferTime  int64
	jwt.StandardClaims
}
type LoginResponse struct {
	User      SysUser `json:"user"`
	Token     string  `json:"token"`
	ExpiresAt int64   `json:"expiresAt"`
}
