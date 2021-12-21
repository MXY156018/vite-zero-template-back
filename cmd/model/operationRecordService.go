package model

import (
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/types"
)

//@author: [granty1](https://github.com/granty1)
//@function: CreateSysOperationRecord
//@description: 创建记录
//@param: sysOperationRecord model.SysOperationRecord
//@return: err error

type OperationRecordService struct {
}

var OperationRecordServiceApp = new(OperationRecordService)

func (operationRecordService *OperationRecordService) CreateSysOperationRecord(sysOperationRecord types.SysOperationRecord) (err error) {
	err = global.GVA_DB.Create(&sysOperationRecord).Error
	return err
}

//@author: [granty1](https://github.com/granty1)
//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysOperationRecordByIds
//@description: 批量删除记录
//@param: ids request.IdsReq
//@return: err error

func (operationRecordService *OperationRecordService) DeleteSysOperationRecordByIds(ids types.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]types.SysOperationRecord{}, "id in (?)", ids.Ids).Error
	return err
}

//@author: [granty1](https://github.com/granty1)
//@function: DeleteSysOperationRecord
//@description: 删除操作记录
//@param: sysOperationRecord model.SysOperationRecord
//@return: err error

func (operationRecordService *OperationRecordService) DeleteSysOperationRecord(sysOperationRecord types.SysOperationRecord) (err error) {
	err = global.GVA_DB.Model(&sysOperationRecord).Where("id=?", sysOperationRecord.ID).Delete(&sysOperationRecord).Error
	return err
}

//@author: [granty1](https://github.com/granty1)
//@function: DeleteSysOperationRecord
//@description: 根据id获取单条操作记录
//@param: id uint
//@return: err error, sysOperationRecord model.SysOperationRecord

func (operationRecordService *OperationRecordService) GetSysOperationRecord(id uint) (err error, sysOperationRecord types.SysOperationRecord) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysOperationRecord).Error
	return
}

//@author: [granty1](https://github.com/granty1)
//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysOperationRecordInfoList
//@description: 分页获取操作记录列表
//@param: info systemReq.SysOperationRecordSearch
//@return: err error, list interface{}, total int64

func (operationRecordService *OperationRecordService) GetSysOperationRecordInfoList(info types.SysOperationRecordSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&types.SysOperationRecord{})
	var sysOperationRecords []types.SysOperationRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.Path != "" {
		db = db.Where("path LIKE ?", "%"+info.Path+"%")
	}
	err = db.Count(&total).Error
	err = db.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(&sysOperationRecords).Error
	return err, sysOperationRecords, total
}
