package database

type NoticeFollower struct {
	Id         int64 `xorm:"'id' autoincr'"`
	NoticeId   int64 `xorm:"notice_id"`
	FollowerId int64 `xorm:"follower_id"`
}

func (NoticeFollower) TableName() string {
	return "notice_follower"
}
