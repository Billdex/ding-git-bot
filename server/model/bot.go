package model

import "time"

type Bot struct {
	Id           int64     `xorm:"'id' autoincr"`
	RepositoryId int64     `xorm:"repository_id"`
	WebHook      string    `xorm:"webhook"`
	Secret       string    `xorm:"secret"`
	Forbidden    bool      `xorm:"forbidden"`
	CreateTime   time.Time `xorm:"'create_time' created"`
	UpdateTime   time.Time `xorm:"'update_time' updated"`
}

func (Bot) TableName() string {
	return "bot"
}
