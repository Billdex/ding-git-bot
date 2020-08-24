package main

import (
	"ding-git-bot/config"
	"ding-git-bot/log"
	"fmt"
)

func main() {
	// 初始化配置信息
	err := config.InitConfig()
	if nil != err {
		fmt.Println("读取配置文件出错！", err)
		return
	}

	// 初始化logger
	log.InitLog(config.AppConfig.Log.Style, config.AppConfig.Log.File, config.AppConfig.Log.Level)
	defer log.Sync()
}
