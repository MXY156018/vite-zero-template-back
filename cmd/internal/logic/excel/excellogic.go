package excel

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/svc"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/model"
	"go-zero-template/cmd/utils"
	"go.uber.org/zap"
	"net/http"
)

type ExcelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExcelLogic(ctx context.Context, svcCtx *svc.ServiceContext) ExcelLogic {
	return ExcelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//ImportExcel 导入Excel
func (e *ExcelLogic) ImportExcel(r *http.Request) (*types.Result, error) {
	_, header, err := r.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "接收文件失败!"}, nil
	}
	_ = model.FileApp.SaveUploadedFile(header, global.GVA_CONFIG.Excel.Dir+"ExcelImport.xlsx")
	return &types.Result{Code: 0, Msg: "导入成功!"}, nil
}

//LoadExcel 加载Excel数据
func (e *ExcelLogic) LoadExcel() (*types.Result, error) {
	menus, err := model.ExcelServiceApp.ParseExcel2InfoList()
	if err != nil {
		global.GVA_LOG.Error("加载数据失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "加载数据失败!"}, nil
	}
	return &types.Result{
		Code: 0,
		Msg:  "加载数据成功",
		Data: types.PageResult{
			List:     menus,
			Total:    int64(len(menus)),
			Page:     1,
			PageSize: 999,
		},
	}, nil
}

//ExportExcel 导出Excel
func (e *ExcelLogic) ExportExcel(req types.ExcelInfo, w http.ResponseWriter, r *http.Request) (*types.Result, error) {
	filePath := global.GVA_CONFIG.Excel.Dir + req.FileName
	err := model.ExcelServiceApp.ParseInfoList2Excel(req.InfoList, filePath)
	if err != nil {
		global.GVA_LOG.Error("转换Excel失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "转换Excel失败!"}, nil
	}
	w.Header().Add("success", "true")
	http.ServeFile(w, r, filePath)
	return &types.Result{Code: 0, Msg: "导出成功!"}, nil
}
func (e *ExcelLogic) DownloadTemplate(w http.ResponseWriter, r *http.Request) (*types.Result, error) {
	fileName := r.URL.Query().Get("fileName")
	filePath := global.GVA_CONFIG.Excel.Dir + fileName
	ok, err := utils.PathExists(filePath)
	if !ok || err != nil {
		global.GVA_LOG.Error("文件不存在!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "文件不存在!"}, nil
	}
	w.Header().Add("success", "true")
	http.ServeFile(w, r, filePath)
	return &types.Result{Code: 0, Msg: "导出成功!"}, nil
}
