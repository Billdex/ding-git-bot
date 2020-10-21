package database

type NoticeEvent struct {
	Id       int64 `xorm:"'id' autoincr'"`
	NoticeId int64 `xorm:"notice_id"`
	EventId  int64 `xorm:"event_id"`
}

func (NoticeEvent) TableName() string {
	return "notice_event"
}
