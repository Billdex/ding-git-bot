package database

type NoticeSender struct {
	Id       int64 `xorm:"'id' autoincr'"`
	NoticeId int64 `xorm:"notice_id"`
	SenderId int64 `xorm:"sender_id"`
}

func (NoticeSender) TableName() string {
	return "notice_sender"
}
