package database

import (
	"ding-git-bot/config"
	"ding-git-bot/log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var DBEngine *xorm.Engine

func InitMysql() error {
	var err error
	user := config.AppConfig.DBConfig.User
	password := config.AppConfig.DBConfig.Password
	host := config.AppConfig.DBConfig.Host
	port := config.AppConfig.DBConfig.Port
	database := config.AppConfig.DBConfig.Database
	//addr := user + ":" + password + "@tcp(" + host + ":" + port +")/" + database + "?charset=utf8&loc=Local"
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Local", user, password, host, port, database)
	log.Info("Init Mysql as : ", addr)

	DBEngine, err = xorm.NewEngine("mysql", addr)
	if err != nil {
		return err
	}

	DBEngine.ShowSQL()

	log.Info("初始化数据库成功！")
	return nil
}
