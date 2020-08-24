package model

type Event struct {
	Id   int64  `xorm:"'id' autoincr"`
	Name string `xorm:"name"`
}

func (Event) TableName() string {
	return "event"
}
