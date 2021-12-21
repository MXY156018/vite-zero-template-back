package main

import (
	"flag"
	"fmt"
	"go-zero-template/cmd/core"
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/initialize"
	"go-zero-template/cmd/internal/config"
	"go-zero-template/cmd/internal/handler"
	"go-zero-template/cmd/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/go-zero-template-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Server
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	global.GVA_VP = core.Viper()      // 初始化Viper
	global.GVA_LOG = core.Zap(c)       // 初始化zap日志库
	global.GVA_DB = initialize.Gorm(c) // gorm连接数据库
	global.GVA_CONFIG = c
	initialize.Timer()

	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	//server.Use(middleware.JWTAuth)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
