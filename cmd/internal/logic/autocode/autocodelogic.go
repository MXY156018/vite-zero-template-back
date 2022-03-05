package autocode

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/svc"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/model"
	"go-zero-template/cmd/utils"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"os"
)

type AutoCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAutoCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) AutoCodeLogic {
	return AutoCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//DelSysHistory 删除回滚记录
func (c *AutoCodeLogic) DelSysHistory(req types.AutoHistoryByID) (*types.Result, error) {
	err := model.AutoCodeHistoryServiceApp.DeletePage(req.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	return &types.Result{Code: 0, Msg: "删除成功"}, nil
}

//GetMeta 根据id获取meta信息
func (c *AutoCodeLogic) GetMeta(req types.AutoHistoryByID) (*types.Result, error) {
	type Meta1 struct {
		Meta string `json:"meta"`
	}
	v, err := model.AutoCodeHistoryServiceApp.GetMeta(req.ID)
	if err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	return &types.Result{
		Code: 0,
		Msg:  "获取成功",
		Data: Meta1{
			Meta: v,
		},
	}, nil
}

//GetSysHistory 获取回滚记录分页
func (c *AutoCodeLogic) GetSysHistory(req types.SysAutoHistory) (*types.Result, error) {
	err, list, total := model.AutoCodeHistoryServiceApp.GetSysHistoryPage(req.PageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "获取失败!"}, nil
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

//RollBack 回滚
func (c *AutoCodeLogic) RollBack(req types.AutoHistoryByID) (*types.Result, error) {
	if err := model.AutoCodeHistoryServiceApp.RollBack(req.ID); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	return &types.Result{Code: 0, Msg: "回滚成功"}, nil
}

//PreviewTemp 获取自动创建代码预览
func (c *AutoCodeLogic) PreviewTemp(req types.AutoCodeStruct) (*types.Result, error) {
	if err := utils.Verify(req, utils.AutoCodeVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	autoCode, err := model.AutoCodeServiceApp.PreviewTemp(req)

	type AutoCode struct {
		AutoCode map[string]string `json:"autoCode"`
	}
	if err != nil {
		global.GVA_LOG.Error("预览失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "预览失败!"}, nil
	} else {
		return &types.Result{
			Code: 0,
			Msg:  "预览成功",
			Data: AutoCode{
				AutoCode: autoCode,
			},
		}, nil
	}
}

//CreateTemp 创建自动化代码
func (c *AutoCodeLogic) CreateTemp(req types.AutoCodeStruct, w http.ResponseWriter, r *http.Request) (*types.Result, error) {
	if err := utils.Verify(req, utils.AutoCodeVerify); err != nil {
		return &types.Result{Code: 7, Msg: err.Error()}, nil
	}
	var apiIds []uint
	if req.AutoCreateApiToSql {
		if ids, err := model.AutoCodeServiceApp.AutoCreateApi(&req); err != nil {
			global.GVA_LOG.Error("自动化创建失败!请自行清空垃圾数据!", zap.Any("err", err))
			w.Header().Add("success", "false")
			w.Header().Add("msg", url.QueryEscape("自动化创建失败!请自行清空垃圾数据!"))
			return &types.Result{Code: 7, Msg: "自动化创建失败!请自行清空垃圾数据!"}, nil
		} else {
			apiIds = ids
		}
	}
	err := model.AutoCodeServiceApp.CreateTemp(req, apiIds...)
	if err != nil {
		if errors.Is(err, types.AutoMoveErr) {
			w.Header().Add("success", "false")
			w.Header().Add("msgtype", "success")
			w.Header().Add("msg", url.QueryEscape(err.Error()))
		} else {
			w.Header().Add("success", "false")
			w.Header().Add("msg", url.QueryEscape(err.Error()))
			_ = os.Remove("./ginvueadmin.zip")
		}
	} else {
		w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "ginvueadmin.zip")) // fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("success", "true")
		http.ServeFile(w, r, "./ginvueadmin.zip")
		_ = os.Remove("./ginvueadmin.zip")
	}
	return &types.Result{Code: 0}, nil
}

//GetTables 获取对应数据库的表
func (c *AutoCodeLogic) GetTables() (*types.Result, error) {
	dbName := global.GVA_CONFIG.Mysql.Dbname
	err, tables := model.AutoCodeServiceApp.GetTables(dbName)
	if err != nil {
		global.GVA_LOG.Error("查询table失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "查询table失败!"}, nil
	}
	type Tables struct {
		Tables []types.TableReq `json:"tables"`
	}
	return &types.Result{Code: 0, Msg: "获取成功",
		Data: Tables{
			Tables: tables,
		},
	}, nil
}
func (c *AutoCodeLogic) GetDB() (*types.Result, error) {
	if err, dbs := model.AutoCodeServiceApp.GetDB(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "获取失败!"}, nil
	} else {
		type Dbs struct {
			Dbs []types.DBReq
		}
		return &types.Result{Code: 0, Msg: "获取成功",
			Data: Dbs{
				Dbs: dbs,
			},
		}, nil
	}

}
func (c *AutoCodeLogic) GetColumn(r *http.Request) (*types.Result, error) {
	dbName := r.URL.Query().Get("dbName")
	tableName := r.URL.Query().Get("tableName")
	if err, columns := model.AutoCodeServiceApp.GetColumn(tableName, dbName); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "获取失败"}, nil
	} else {
		type Col struct {
			Columns []types.ColumnReq
		}
		return &types.Result{Code: 0, Msg: "获取成功", Data: Col{Columns: columns}}, nil
	}
}
