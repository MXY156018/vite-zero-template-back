package types

import "go-zero-template/cmd/global"

// file struct, 文件结构体
type ExaFile struct {
	global.GVA_MODEL
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
}

// file chunk struct, 切片结构体
type ExaFileChunk struct {
	global.GVA_MODEL
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string
}
type FilePathResponse struct {
	FilePath string `json:"filePath"`
}

type FileResponse struct {
	File ExaFile `json:"file"`
}
