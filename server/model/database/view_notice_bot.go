package database

type ViewNoticeBot struct {
	NoticeId        int64  `xorm:"notice_id"`
	RepositoryId    int64  `xorm:"repository_id"`
	Message         string `xorm:"message"`
	NoticeForbidden bool   `xorm:"notice_forbidden"`
	BotId           int64  `xorm:"bot_id"`
	AccessToken     string `xorm:"access_token"`
	Secret          string `xorm:"secret"`
	BotForbidden    bool   `xorm:"bot_forbidden"`
	AtAll           bool   `xorm:"at_all"`
}

func (ViewNoticeBot) TableName() string {
	return "view_notice_bot"
}
