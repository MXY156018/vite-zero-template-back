package file

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/svc"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/model"
	"go-zero-template/cmd/utils"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
)

type FileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) FileLogic {
	return FileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//UploadFile 上传文件
func (F *FileLogic) UploadFile(r *http.Request) (*types.Result, error) {
	var file types.ExaFileUploadAndDownload
	noSave := r.URL.Query().Get("noSave")
	if noSave == "" {
		noSave = "0"
	}
	_, header, err := r.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "接收文件失败"}, nil
	}
	err, file = model.FileApp.UploadFile(header, noSave) // 文件上传后拿到文件路径
	if err != nil {
		global.GVA_LOG.Error("修改数据库链接失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "修改数据库链接失败"}, nil
	}
	return &types.Result{
		Code: 0,
		Msg:  "上传成功",
		Data: types.ExaFileResponse{
			File: file,
		},
	}, nil
}

//GetFileList 分页文件列表
func (F *FileLogic) GetFileList(req types.PageInfo) (*types.Result, error) {
	err, list, total := model.FileApp.GetFileRecordInfoList(req)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "获取失败"}, nil
	} else {
		return &types.Result{
			Code: 0,
			Msg:  "获取成功",
			Data: types.PageResult{
				List:     list,
				Total:    total,
				Page:     req.Page,
				PageSize: req.PageSize,
			},
		}, nil

	}
}

//DeleteFile 删除文件
func (F *FileLogic) DeleteFile(req types.ExaFileUploadAndDownload) (*types.Result, error) {
	if err := model.FileApp.DeleteFile(req); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "删除失败"}, nil
	}
	return &types.Result{Code: 0, Msg: "删除成功"}, nil
}

//BreakpointContinue 断点续传
func (F *FileLogic) BreakpointContinue(r *http.Request) (*types.Result, error) {
	fileMd5 := r.FormValue("fileMd5")
	fileName := r.FormValue("fileName")
	chunkMd5 := r.FormValue("chunkMd5")
	chunkNumber, _ := strconv.Atoi(r.FormValue("chunkNumber"))
	chunkTotal, _ := strconv.Atoi(r.FormValue("chunkTotal"))
	_, FileHeader, err := r.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "接收文件失败"}, nil
	}
	f, err := FileHeader.Open()
	if err != nil {
		global.GVA_LOG.Error("文件读取失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "文件读取失败"}, nil
	}
	defer f.Close()
	cen, _ := ioutil.ReadAll(f)
	if !utils.CheckMd5(cen, chunkMd5) {
		global.GVA_LOG.Error("检查md5失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "检查md5失败"}, nil
	}
	err, file := model.FileBreakContinueServiceApp.FindOrCreateFile(fileMd5, fileName, chunkTotal)
	if err != nil {
		global.GVA_LOG.Error("查找或创建记录失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "查找或创建记录失败"}, nil
	}
	err, pathc := utils.BreakPointContinue(cen, fileName, chunkNumber, chunkTotal, fileMd5)
	if err != nil {
		global.GVA_LOG.Error("断点续传失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "断点续传失败"}, nil
	}
	if err = model.FileBreakContinueServiceApp.CreateFileChunk(file.ID, pathc, chunkNumber); err != nil {
		global.GVA_LOG.Error("创建文件记录失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "创建文件记录失败"}, nil
	}
	return &types.Result{Code: 0, Msg: "切片创建成功"}, nil
}

//FindFile 查找文件
func (F *FileLogic) FindFile(r *http.Request) (*types.Result, error) {
	fileMd5 := r.URL.Query().Get("fileMd5")
	fileName := r.URL.Query().Get("fileName")
	chunkTotal, _ := strconv.Atoi(r.URL.Query().Get("chunkTotal"))
	err, file := model.FileBreakContinueServiceApp.FindOrCreateFile(fileMd5, fileName, chunkTotal)
	if err != nil {
		global.GVA_LOG.Error("查找失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "查找失败"}, nil
	} else {
		return &types.Result{
			Code: 0,
			Msg:  "切片创建成功",
			Data: types.FileResponse{
				File: file,
			},
		}, nil
	}
}

//BreakpointContinueFinish 创建文件  上传文件完成
func (F *FileLogic) BreakpointContinueFinish(r *http.Request) (*types.Result, error) {
	fileMd5 := r.URL.Query().Get("fileMd5")
	fileName := r.URL.Query().Get("fileName")

	err, filePath := utils.MakeFile(fileName, fileMd5)
	if err != nil {
		global.GVA_LOG.Error("文件创建失败!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "文件创建失败",
			Data: types.FilePathResponse{
				FilePath: filePath,
			},
		}, nil
	} else {
		return &types.Result{
			Code: 0,
			Msg:  "文件创建成功",
			Data: types.FilePathResponse{
				FilePath: filePath,
			},
		}, nil
	}
}

func (F *FileLogic) RemoveChunk(r *http.Request) (*types.Result, error) {
	fileMd5 := r.URL.Query().Get("fileMd5")
	fileName := r.URL.Query().Get("fileName")
	filePath := r.URL.Query().Get("filePath")
	err := utils.RemoveChunk(fileMd5)
	err = model.FileBreakContinueServiceApp.DeleteFileChunk(fileMd5, fileName, filePath)
	if err != nil {
		global.GVA_LOG.Error("缓存切片删除失败!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "缓存切片删除失败",
			Data: types.FilePathResponse{
				FilePath: filePath,
			},
		}, nil
	} else {
		return &types.Result{
			Code: 0,
			Msg:  "缓存切片删除成功",
			Data: types.FilePathResponse{
				FilePath: filePath,
			},
		}, nil
	}
}
