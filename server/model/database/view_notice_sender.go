package database

type ViewNoticeSender struct {
	NoticeId        int64  `xorm:"notice_id"`
	RepositoryId    int64  `xorm:"repository_id"`
	Message         string `xorm:"message"`
	NoticeForbidden bool   `xorm:"notice_forbidden"`
	SenderId        int64  `xorm:"sender_id"`
	SenderName      string `xorm:"sender_name"`
	SenderEmail     string `xorm:"sender_email"`
	SenderPhone     string `xorm:"sender_phone"`
}

func (ViewNoticeSender) TableName() string {
	return "view_notice_sender"
}
