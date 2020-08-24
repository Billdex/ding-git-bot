package model

import "time"

type Repository struct {
	Id         int64     `xorm:"'id' autoincr"`
	Name       string    `xorm:"name"`
	Secret     string    `xorm:"secret"`
	Url        string    `xorm:"url"`
	CreateTime time.Time `xorm:"'create_time' created"`
	UpdateTime time.Time `xorm:"'update_time' updated"`
}

func (Repository) TableName() string {
	return "repository"
}
