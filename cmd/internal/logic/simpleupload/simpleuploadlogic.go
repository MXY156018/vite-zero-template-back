package simpleupload

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/svc"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/model"
	"go-zero-template/cmd/utils"
	"go.uber.org/zap"
	"net/http"
)

type SimpleUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSimpleUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) SimpleUploadLogic {
	return SimpleUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//SimpleUploaderUpload 上传功能
func (s *SimpleUploadLogic) SimpleUploaderUpload(r *http.Request) (*types.Result, error) {
	var chunk types.ExaSimpleUploader
	_, header, err := r.FormFile("file")
	chunk.Filename = r.PostForm.Get("filename")
	chunk.ChunkNumber = r.PostForm.Get("chunkNumber")
	chunk.CurrentChunkSize = r.PostForm.Get("currentChunkSize")
	chunk.Identifier = r.PostForm.Get("identifier")
	chunk.TotalSize = r.PostForm.Get("totalSize")
	chunk.TotalChunks = r.PostForm.Get("totalChunks")
	var chunkDir = "./chunk/" + chunk.Identifier + "/"
	hasDir, _ := utils.PathExists(chunkDir)
	if !hasDir {
		if err := utils.CreateDir(chunkDir); err != nil {
			global.GVA_LOG.Error("创建目录失败!", zap.Any("err", err))
		}
	}
	chunkPath := chunkDir + chunk.Filename + chunk.ChunkNumber
	err = model.FileApp.SaveUploadedFile(header, chunkPath)
	if err != nil {
		global.GVA_LOG.Error("切片创建失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "切片创建失败!"}, nil
	}
	chunk.CurrentChunkPath = chunkPath
	err = model.SimpleUploaderServiceApp.SaveChunk(chunk)
	if err != nil {
		global.GVA_LOG.Error("切片创建失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "切片创建失败!"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "切片创建成功!"}, nil
	}
}

//CheckFileMd5 文件完整度验证
func (s *SimpleUploadLogic) CheckFileMd5(r *http.Request) (*types.Result, error) {
	md5 := r.URL.Query().Get("md5")
	err, chunks, isDone := model.SimpleUploaderServiceApp.CheckFileMd5(md5)
	if err != nil {
		global.GVA_LOG.Error("md5读取失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "md5读取失败!"}, nil
	} else {
		return &types.Result{
			Code: 0,
			Msg:  "查询成功",
			Data: types.ExaSimpleSearch{
				Chunks: chunks,
				IsDone: isDone,
			},
		}, nil

	}
}
func (s *SimpleUploadLogic) MergeFileMd5(r *http.Request) (*types.Result, error) {
	md5 := r.URL.Query().Get("md5")
	fileName := r.URL.Query().Get("fileName")
	err := model.SimpleUploaderServiceApp.MergeFileMd5(md5, fileName)
	if err != nil {
		global.GVA_LOG.Error("md5读取失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "md5读取失败!"}, nil
	} else {
		return &types.Result{Code: 0, Msg: "合并成功!"}, nil
	}
}
