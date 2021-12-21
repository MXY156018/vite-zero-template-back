package types

import "go-zero-template/cmd/global"

type AutoHistoryByID struct {
	ID uint `json:"id"`
}

type DBReq struct {
	Database string `json:"database" gorm:"column:database"`
}

type TableReq struct {
	TableName string `json:"tableName"`
}

type ColumnReq struct {
	ColumnName    string `json:"columnName" gorm:"column:column_name"`
	DataType      string `json:"dataType" gorm:"column:data_type"`
	DataTypeLong  string `json:"dataTypeLong" gorm:"column:data_type_long"`
	ColumnComment string `json:"columnComment" gorm:"column:column_comment"`
}
type SysAutoCodeHistory struct {
	global.GVA_MODEL
	TableName     string `json:"tableName"`
	RequestMeta   string `gorm:"type:text" json:"requestMeta,omitempty"`   // 前端传入的结构化信息
	AutoCodePath  string `gorm:"type:text" json:"autoCodePath,omitempty"`  // 其他meta信息 path;path
	InjectionMeta string `gorm:"type:text" json:"injectionMeta,omitempty"` // 注入的内容 RouterPath@functionName@RouterString;
	StructName    string `json:"structName"`
	StructCNName  string `json:"structCNName"`
	ApiIDs        string `json:"apiIDs,omitempty"` // api表注册内容
	Flag          int    `json:"flag"`             // 表示对应状态 0 代表创建, 1 代表回滚 ...

}
type SysAutoHistory struct {
	PageInfo
}
