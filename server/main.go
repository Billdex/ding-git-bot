package main

import (
	"ding-git-bot/config"
	"ding-git-bot/log"
	"ding-git-bot/model/database"
	"ding-git-bot/router"
	"fmt"
)

func main() {
	// 初始化配置信息
	err := config.InitConfig()
	if err != nil {
		fmt.Println("读取配置文件出错！", err)
		return
	}

	// 初始化logger
	err = log.InitLog(config.AppConfig.Log.Style, config.AppConfig.Log.File, config.AppConfig.Log.Level)
	if err != nil {
		fmt.Println("初始化logger出错！", err)
		return
	}
	defer log.Sync()

	// 初始化数据库
	err = database.InitMysql()
	if err != nil {
		log.Error("初始化数据库失败！", err)
		return
	}

	// 启动web服务
	err = router.InitRouter()
	if err != nil {
		log.Error("Web服务启动失败！", err)
	}
}
