package database

import "time"

type Bot struct {
	Id          int64     `xorm:"'id' autoincr"`
	AccessToken string    `xorm:"access_token"`
	Secret      string    `xorm:"secret"`
	Forbidden   bool      `xorm:"forbidden"`
	CreateTime  time.Time `xorm:"'create_time' created"`
	UpdateTime  time.Time `xorm:"'update_time' updated"`
}

func (Bot) TableName() string {
	return "bot"
}
