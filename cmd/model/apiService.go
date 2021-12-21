package model

import (
	"errors"
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/types"
	"gorm.io/gorm"
	"time"
)

type ApiService struct {
}

var ApiServiceApp = new(ApiService)

func (apiService *ApiService) CreateApi(api types.SysApi) (err error) {
	if !errors.Is(global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&types.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.GVA_DB.Create(&api).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteApi
//@description: 删除基础api
//@param: api model.SysApi
//@return: err error
func (apiService *ApiService) DeleteApi(api types.SysApi) (err error) {
	err = global.GVA_DB.Table("sys_apis").Where("id=?", api.ID).Update("deleted_at", time.Now()).Error
	CasbinServiceApp.ClearCasbin(1, api.Path, api.Method)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAPIInfoList
//@description: 分页获取数据,
//@param: api model.SysApi, info request.PageInfo, order string, desc bool
//@return: err error

func (apiService *ApiService) GetAPIInfoList(api types.SysApi, info types.PageInfo, order string, desc bool) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Table("sys_apis")
	var apiList []types.SysApi

	if api.Path != "" {
		db = db.Where("path LIKE ?", "%"+api.Path+"%")
	}
	if api.Description != "" {
		db = db.Where("description LIKE ?", "%"+api.Description+"%")
	}
	if api.Method != "" {
		db = db.Where("method = ?", api.Method)
	}
	if api.ApiGroup != "" {
		db = db.Where("api_group = ?", api.ApiGroup)
	}

	err = db.Count(&total).Error

	if err != nil {
		return err, apiList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			if desc {
				OrderStr = order + " desc"
			} else {
				OrderStr = order
			}
			err = db.Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Order("api_group").Find(&apiList).Error
		}
	}
	return err, apiList, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetApiById
//@description: 根据id获取api
//@param: id float64
//@return: err error, api model.SysApi
func (apiService *ApiService) GetApiById(id float64) (err error, api types.SysApi) {
	err = global.GVA_DB.Where("id = ?", id).First(&api).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateApi
//@description: 根据id更新api
//@param: api model.SysApi
//@return: err error

func (apiService *ApiService) UpdateApi(api types.SysApi) (err error) {
	var oldA types.SysApi
	err = global.GVA_DB.Where("id = ?", api.ID).First(&oldA).Error
	if oldA.Path != api.Path || oldA.Method != api.Method {
		if !errors.Is(global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&types.SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	} else {
		err = CasbinServiceApp.UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
		if err != nil {
			return err
		} else {
			err = global.GVA_DB.Save(&api).Error
		}
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAllApis
//@description: 获取所有的api
//@return: err error, apis []model.SysApi

func (apiService *ApiService) GetAllApis() (err error, apis []types.SysApi) {
	err = global.GVA_DB.Find(&apis).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteApis
//@description: 删除选中API
//@param: apis []model.SysApi
//@return: err error

func (apiService *ApiService) DeleteApisByIds(ids types.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]types.SysApi{}, "id in ?", ids.Ids).Error
	return err
}
func (apiService *ApiService) DeleteApiByIds(ids []string) (err error) {
	return global.GVA_DB.Delete(types.SysApi{}, ids).Error
}
