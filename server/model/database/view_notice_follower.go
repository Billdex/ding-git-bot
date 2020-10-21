package database

type ViewNoticeFollower struct {
	NoticeId        int64  `xorm:"notice_id"`
	RepositoryId    int64  `xorm:"repository_id"`
	Message         string `xorm:"message"`
	NoticeForbidden bool   `xorm:"notice_forbidden"`
	FollowerId      int64  `xorm:"follower_id"`
	FollowerName    string `xorm:"follower_name"`
	FollowerEmail   string `xorm:"follower_email"`
	FollowerPhone   string `xorm:"follower_phone"`
}

func (ViewNoticeFollower) TableName() string {
	return "view_notice_follower"
}
