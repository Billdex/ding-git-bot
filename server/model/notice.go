package model

import "time"

type Notice struct {
	Id           int64     `xorm:"'id' autoincr"`
	RepositoryId int64     `xorm:"repository_id"`
	BotId        int64     `xorm:"bot_id"`
	Events       []int64   `xorm:"events_id"`
	Message      string    `xorm:"message"`
	Senders      []string  `xorm:"sender"`
	Followers    []string  `xorm:"follower"`
	Forbidden    bool      `xorm:"forbidden"`
	CreateTime   time.Time `xorm:"'create_time' created"`
	UpdateTime   time.Time `xorm:"'update_time' updated"`
}

func (Notice) TableName() string {
	return "notice"
}
