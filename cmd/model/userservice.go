package model

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/utils"
	"gorm.io/gorm"
)

func Register(u types.SysUser) (err error, userInter types.SysUser) {
	var user types.SysUser
	if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = global.GVA_DB.Create(&u).Error
	return err, u
}
func ChangePassword(u *types.SysUser, newPassword string) (err error, userInter *types.SysUser) {
	var user types.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}
func GetUserInfoList(info types.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&types.SysUser{})
	var userList []types.SysUser
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return err, userList, total
}

//SetUserAuthority 更改用户权限
func SetUserAuthority(id uint, uuid uuid.UUID, authorityId string) (err error) {
	assignErr := global.GVA_DB.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityId).First(&types.SysUserAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}
	err = global.GVA_DB.Where("uuid = ?", uuid).First(&types.SysUser{}).Update("authority_id", authorityId).Error
	return err
}
func FindUserByUuid(uuid string) (err error, user *types.SysUser) {
	var u types.SysUser
	if err = global.GVA_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}
func DeleteUser(id float64) (err error) {
	var user types.SysUser
	err = global.GVA_DB.Where("id = ?", id).Delete(&user).Error
	err = global.GVA_DB.Table("sys_user_authority").Where("sys_user_id = ?", id).Delete(&types.SysUserAuthority{}).Error
	return err
}
func SetUserInfo(reqUser types.SysUser) (err error, user types.SysUser) {
	err = global.GVA_DB.Updates(&reqUser).Error
	return err, reqUser
}
func SetUserAuthorities(id uint, authorityIds []string) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]types.SysUserAuthority{}, "sys_user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		var useAuthority []types.SysUserAuthority
		for _, v := range authorityIds {
			useAuthority = append(useAuthority, types.SysUserAuthority{
				SysUserId: id, SysAuthorityAuthorityId: v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		// 返回 nil 提交事务
		return nil
	})
}
func GetUserInfo(uuid uuid.UUID) (err error, user types.SysUser) {
	var reqUser types.SysUser
	err = global.GVA_DB.Preload("Authorities").Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
	return err, reqUser
}
