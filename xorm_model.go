package carbon

type XormModel struct {
	Id        int64             `xorm:"not null pk autoincr INT(11)" json:"id"`
	CreatedAt ToDateTimeString  `xorm:"created_at created" json:"creat_time" description:"创建时间"`
	UpdatedAt ToDateTimeString  `xorm:"updated_at updated" json:"update_time" description:"更新时间"`
	DeletedAt *ToDateTimeString `xorm:"deleted_time deleted" json:"deleted_time" description:"删除时间"`
}
