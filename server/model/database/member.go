package database

import "time"

type Member struct {
	Id           int64     `xorm:"'id' autoincr"`
	RepositoryId int64     `xorm:"repository_id"`
	Name         string    `xorm:"name"`
	Email        string    `xorm:"email"`
	Phone        string    `xorm:"phone"`
	CreateTime   time.Time `xorm:"'create_time' created"`
	UpdateTime   time.Time `xorm:"'update_time' updated"`
}

func (Member) TableName() string {
	return "member"
}
