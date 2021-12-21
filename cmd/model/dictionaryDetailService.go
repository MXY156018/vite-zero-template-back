package model

import (
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/types"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateSysDictionaryDetail
//@description: 创建字典详情数据
//@param: sysDictionaryDetail model.SysDictionaryDetail
//@return: err error

type DictionaryDetailService struct {
}

var DictionaryDetailServiceApp = new(DictionaryDetailService)

func (dictionaryDetailService *DictionaryDetailService) CreateSysDictionaryDetail(sysDictionaryDetail types.SysDictionaryDetail) (err error) {
	err = global.GVA_DB.Create(&sysDictionaryDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysDictionaryDetail
//@description: 删除字典详情数据
//@param: sysDictionaryDetail model.SysDictionaryDetail
//@return: err error

func (dictionaryDetailService *DictionaryDetailService) DeleteSysDictionaryDetail(sysDictionaryDetail types.SysDictionaryDetail) (err error) {
	err = global.GVA_DB.Delete(&sysDictionaryDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSysDictionaryDetail
//@description: 更新字典详情数据
//@param: sysDictionaryDetail *model.SysDictionaryDetail
//@return: err error

func (dictionaryDetailService *DictionaryDetailService) UpdateSysDictionaryDetail(sysDictionaryDetail *types.SysDictionaryDetail) (err error) {
	err = global.GVA_DB.Save(sysDictionaryDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysDictionaryDetail
//@description: 根据id获取字典详情单条数据
//@param: id uint
//@return: err error, sysDictionaryDetail model.SysDictionaryDetail

func (dictionaryDetailService *DictionaryDetailService) GetSysDictionaryDetail(id uint) (err error, sysDictionaryDetail types.SysDictionaryDetail) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysDictionaryDetail).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysDictionaryDetailInfoList
//@description: 分页获取字典详情列表
//@param: info request.SysDictionaryDetailSearch
//@return: err error, list interface{}, total int64

func (dictionaryDetailService *DictionaryDetailService) GetSysDictionaryDetailInfoList(info types.SysDictionaryDetailSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&types.SysDictionaryDetail{})
	var sysDictionaryDetails []types.SysDictionaryDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Label != "" {
		db = db.Where("label LIKE ?", "%"+info.Label+"%")
	}
	if info.Value != 0 {
		db = db.Where("value = ?", info.Value)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.SysDictionaryID != 0 {
		db = db.Where("sys_dictionary_id = ?", info.SysDictionaryID)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sysDictionaryDetails).Error
	return err, sysDictionaryDetails, total
}
