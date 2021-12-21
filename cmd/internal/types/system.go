package types

import (
	"go-zero-template/cmd/internal/config"
	"time"
)

type SysUserAuthority struct {
	SysUserId               uint   `gorm:"column:sys_user_id"`
	SysAuthorityAuthorityId string `gorm:"column:sys_authority_authority_id"`
}
type JwtBlacklist struct {
	Status int    `json:"status,default= 1"`
	Jwt    string `json:"jwt"`
}
type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []string `json:"authorityIds"` // 角色ID
}
type SysOperationRecord struct {
	ID           uint      `json:"id,optional"`
	CreatedAt    time.Time `json:"created_at,optional"`
	Ip           string    `json:"ip,optional"`            // 请求ip
	Method       string    `json:"method,optional"`        // 请求方法
	Path         string    `json:"path,optional"`          // 请求路径
	Status       int       `json:"status,optional"`        // 状态
	Agent        string    `json:"agent,optional"`         // 代理
	ErrorMessage string    `json:"error_message,optional"` // 错误信息
	Body         string    `json:"body,optional"`          // 请求Body
	UserID       int       `json:"user_id,optional"`       // 用户id
	User         SysUser   `json:"user,optional"`
}
type CasbinModel struct {
	Ptype       string `json:"ptype" gorm:"column:ptype"`
	AuthorityId string `json:"rolename" gorm:"column:v0"`
	Path        string `json:"path" gorm:"column:v1"`
	Method      string `json:"method" gorm:"column:v2"`
}
type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

// api分页条件查询及排序结构体
type SearchApiParams struct {
	SysApi
	PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}
type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
type System struct {
	Config config.Server `json:"config"`
}
type Server1 struct {
	ServerInfo *Server `json:"server"`
}

type Server struct {
	Os   Os   `json:"os"`
	Cpu  Cpu  `json:"cpu"`
	Rrm  Rrm  `json:"ram"`
	Disk Disk `json:"disk"`
}

type Os struct {
	GOOS         string `json:"goos"`
	NumCPU       int    `json:"numCpu"`
	Compiler     string `json:"compiler"`
	GoVersion    string `json:"goVersion"`
	NumGoroutine int    `json:"numGoroutine"`
}

type Cpu struct {
	Cpus  []float64 `json:"cpus"`
	Cores int       `json:"cores"`
}

type Rrm struct {
	UsedMB      int `json:"usedMb"`
	TotalMB     int `json:"totalMb"`
	UsedPercent int `json:"usedPercent"`
}

type Disk struct {
	UsedMB      int `json:"usedMb"`
	UsedGB      int `json:"usedGb"`
	TotalMB     int `json:"totalMb"`
	TotalGB     int `json:"totalGb"`
	UsedPercent int `json:"usedPercent"`
}
