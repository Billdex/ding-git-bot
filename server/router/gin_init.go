package router

import (
	"ding-git-bot/config"
	. "ding-git-bot/controller"
	"github.com/gin-gonic/gin"
	"strconv"
)

//初始化路由
func InitRouter() error {
	r := gin.Default()

	r.POST("/git", WebHookParsing)

	port := strconv.Itoa(config.AppConfig.Server.Port)
	err := r.Run(":" + port)
	return err
}
