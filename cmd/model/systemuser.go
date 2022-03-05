package model

import (
	"fmt"
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/types"
)

func CreateSysOperationRecord(sysOperationRecord types.SysOperationRecord) (err error) {
	fmt.Sprintf("%v\n", sysOperationRecord)
	err = global.GVA_DB.Create(&sysOperationRecord).Error
	return err
}
