package main

import (
	"graduation/dao/mysql"
	"graduation/logger"
	"graduation/pkg/snowflake"
	"graduation/routes"
	"graduation/settings"
	"fmt"

	"go.uber.org/zap"
)

func main() {
	// 加载配置
	if err := settings.Init(); err != nil {
		fmt.Println("init settings failed, err:%v\n", err)
		return
	}
	// 初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Println("init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	// 初始化MySQL
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Println("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()

	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Println("init snowflake failed, err:%v\n", err)
		return
	}

	// 注册路由
	r := routes.Setup(settings.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}

}