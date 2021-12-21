package model

import (
	"errors"
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/utils/upload"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

type FileUploadAndDownloadService struct {
}

var FileApp = new(FileUploadAndDownloadService)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Upload
//@description: 创建文件上传记录
//@param: file model.ExaFileUploadAndDownload
//@return: error

func (e *FileUploadAndDownloadService) Upload(file types.ExaFileUploadAndDownload) error {
	return global.GVA_DB.Create(&file).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: FindFile
//@description: 删除文件切片记录
//@param: id uint
//@return: error, model.ExaFileUploadAndDownload

func (e *FileUploadAndDownloadService) FindFile(id uint) (error, types.ExaFileUploadAndDownload) {
	var file types.ExaFileUploadAndDownload
	err := global.GVA_DB.Where("id = ?", id).First(&file).Error
	return err, file
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFile
//@description: 删除文件记录
//@param: file model.ExaFileUploadAndDownload
//@return: err error

func (e *FileUploadAndDownloadService) DeleteFile(file types.ExaFileUploadAndDownload) (err error) {
	var fileFromDb types.ExaFileUploadAndDownload
	err, fileFromDb = e.FindFile(file.ID)
	oss := upload.NewOss()
	if err = oss.DeleteFile(fileFromDb.Key); err != nil {
		return errors.New("文件删除失败")
	}
	err = global.GVA_DB.Where("id = ?", file.ID).Unscoped().Delete(&file).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFileRecordInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (e *FileUploadAndDownloadService) GetFileRecordInfoList(info types.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	var fileLists []types.ExaFileUploadAndDownload
	err = db.Find(&fileLists).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	return err, fileLists, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UploadFile
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader, noSave string
//@return: err error, file model.ExaFileUploadAndDownload

func (e *FileUploadAndDownloadService) UploadFile(header *multipart.FileHeader, noSave string) (err error, file types.ExaFileUploadAndDownload) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(err)
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := types.ExaFileUploadAndDownload{
			Url:  filePath,
			Name: header.Filename,
			Tag:  s[len(s)-1],
			Key:  key,
		}
		return e.Upload(f), f
	}
	return
}

//SaveUploadedFile 保存文件到指定目录
func (e *FileUploadAndDownloadService) SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
