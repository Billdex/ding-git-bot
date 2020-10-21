package database

type NoticeBot struct {
	Id       int64 `xorm:"'id' autoincr'"`
	NoticeId int64 `xorm:"notice_id"`
	BotId    int64 `xorm:"bot_id"`
	AtAll    bool  `xorm:"at_all"`
}

func (NoticeBot) TableName() string {
	return "notice_bot"
}
