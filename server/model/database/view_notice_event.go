package database

type ViewNoticeEvent struct {
	NoticeId        int64  `xorm:"notice_id"`
	RepositoryId    int64  `xorm:"repository_id"`
	Message         string `xorm:"message"`
	NoticeForbidden bool   `xorm:"notice_forbidden"`
	EventId         int64  `xorm:"event_id"`
	Event           string `xorm:"event"`
}

func (ViewNoticeEvent) TableName() string {
	return "view_notice_event"
}
