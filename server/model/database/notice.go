package database

import "time"

type Notice struct {
	Id           int64     `xorm:"'id' autoincr"`
	RepositoryId int64     `xorm:"repository_id"`
	Message      string    `xorm:"message"`
	Forbidden    bool      `xorm:"forbidden"`
	CreateTime   time.Time `xorm:"'create_time' created"`
	UpdateTime   time.Time `xorm:"'update_time' updated"`
}

func (Notice) TableName() string {
	return "notice"
}
