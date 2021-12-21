package model

import (
	"errors"
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/config"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/utils"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

type SystemConfigService struct {
}

var SystemConfigServiceApp = new(SystemConfigService)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSystemConfig
//@description: 读取配置文件
//@return: err error, conf config.Server
func (systemConfigService *SystemConfigService) GetSystemConfig() (err error, conf config.Server) {
	return nil, global.GVA_CONFIG
}

// @description   set system config,
//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetSystemConfig
//@description: 设置配置文件
//@param: system model.System
//@return: err error

func (systemConfigService *SystemConfigService) SetSystemConfig(system types.System) (err error) {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	err = global.GVA_VP.WriteConfig()
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetServerInfo
//@description: 获取服务器信息
//@return: server *utils.Server, err error

func (systemConfigService *SystemConfigService) GetServerInfo() (server *types.Server, err error) {
	var s types.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		global.GVA_LOG.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Rrm, err = utils.InitRAM(); err != nil {
		global.GVA_LOG.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		global.GVA_LOG.Error("func utils.InitDisk() Failed", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}
func (systemConfigService *SystemConfigService) Reload() error {
	if runtime.GOOS == "windows" {
		return errors.New("系统不支持")
	}
	pid := os.Getpid()
	cmd := exec.Command("kill", "-1", strconv.Itoa(pid))
	return cmd.Run()
}
