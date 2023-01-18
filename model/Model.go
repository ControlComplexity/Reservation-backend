package model

import (
	"time"
)

var (
	activityTableName = "activity"
	userTableName     = "user"
)

type ActivityDO struct {
	ID          int64     `gorm:"column:id;type:serial;primaryKey;comment:自增id"`
	Name        string    `gorm:"column:name;type:varchar(50);index;not null;default:'';comment:活动名称"`
	Price       float32   `gorm:"column:price;type:float;index;not null;default:0;comment:价格"`
	Location    string    `gorm:"column:location;type:varchar(50);index;not null;default:'';comment:地点"`
	Time        time.Time `gorm:"column:time;type:timestamp;autoUpdateTime;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`
	SmallImg    string    `gorm:"column:small_img;type:varchar(50);index;not null;default:'';comment:小图"`
	BigImg      string    `gorm:"column:big_img;type:varchar(50);index;not null;default:'';comment:大图"`
	Description string    `gorm:"column:description;type:varchar(50);index;not null;default:'';comment:描述"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;autoCreateTime;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;autoUpdateTime;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`
}

func (ActivityDO) TableName() string {
	return activityTableName
}

type UserDO struct {
	ID          int64     `gorm:"column:id;type:serial;primaryKey;comment:自增id"`
	Name        string    `gorm:"column:name;type:varchar(50);index;not null;default:'';comment:用户名"`
	NickName    string    `gorm:"column:nick_name;type:varchar(50);index;not null;default:'';comment:昵称"`
	OpenID      string    `gorm:"column:open_id;type:varchar(50);index;not null;default:'';comment:OpenID"`
	Token       string    `gorm:"column:token;type:varchar(50);index;not null;default:'';comment:Token"`
	PhoneNumber int64     `gorm:"column:phone_number;type:int;comment:手机号"`
	HeadImage   string    `gorm:"column:head_image;type:varchar(50);index;not null;default:'';comment:头像"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;autoCreateTime;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;autoUpdateTime;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`
}

func (UserDO) TableName() string {
	return userTableName
}
