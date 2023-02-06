package model

import (
	"time"
)

var (
	swiperTableName   = "swiper"
	activityTableName = "activity"
	userTableName     = "user"
	enlistTableName   = "enlist"
	orderTableName    = "order"
)

type Model struct {
	Id int64 `gorm:"primaryKey;autoIncrement" json:"ID" form:"ID"`
}

type ActivityDO struct {
	Model
	Name        string    `gorm:"column:name;type:varchar(50);index;not null;default:'';comment:活动名称"`
	Label       string    `gorm:"column:label;type:varchar(50);index;not null;default:'';comment:活动标签"`
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
	Model
	HeadImage       string    `gorm:"column:head_image;type:varchar(50);index;not null;default:'';comment:头像"`
	NickName        string    `gorm:"column:nick_name;type:varchar(50);index;not null;default:'';comment:昵称"`
	Gender          string    `gorm:"type:enum('male', 'female');default:'male';comment:性别"`
	Height          int32     `gorm:"column:height;type:int;comment:身高(cm)"`
	Weight          int32     `gorm:"column:weight;type:int;comment:体重(kg)"`
	Hometown        string    `gorm:"column:hometown;type:varchar(50);not null;default:'';comment:籍贯"`
	Location        string    `gorm:"column:location;type:varchar(50);not null;default:'';comment:所在地"`
	EmotionalStatus string    `gorm:"type:enum('SINGLE_WITHOUT_MARRIAGE_HISTORY', 'SINGLE_AND_DIVORCED');default:'SINGLE_AND_DIVORCED';comment:目前情感状态"`
	Education       string    `gorm:"type:enum('UNIFIED_BACHELOR', 'PART_TIME_BACHELOR','TRANSFORMED_BACHELOR','FULL_TIME_MASTER', 'PART_TIME_MASTER','FULL_TIME_DOCTOR','PART_TIME_DOCTOR','OTHER');default:'UNIFIED_BACHELOR';comment:最高学历"`
	University      string    `gorm:"column:university;type:varchar(50);not null;default:'';comment:毕业(在读)院校"`
	Occupation      string    `gorm:"column:occupation;type:varchar(50);not null;default:'';comment:职业"`
	Company         string    `gorm:"column:company;type:varchar(50);not null;default:'';comment:公司"`
	WechatNumber    string    `gorm:"column:wechat_number;type:varchar(50);not null;default:'';comment:微信号"`
	PhoneNumber     string    `gorm:"column:phone_number;type:varchar(50);comment:手机号"`
	OpenID          string    `gorm:"column:open_id;type:varchar(50);not null;default:'';comment:OpenID"`
	Token           string    `gorm:"column:token;type:varchar(50);not null;default:'';comment:Token"`
	CreatedAt       time.Time `gorm:"column:created_at;type:timestamp;autoCreateTime;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt       time.Time `gorm:"column:updated_at;type:timestamp;autoUpdateTime;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`
}

func (UserDO) TableName() string {
	return userTableName
}

type SwiperDO struct {
	Model
	Img string `gorm:"column:img;type:string;"`
	URL string `gorm:"column:url;type:string;"`
}

func (SwiperDO) TableName() string {
	return swiperTableName
}

type EnlistDO struct {
	Model
	user1Id int64 `gorm:"column:user_1_id;type:int;"`
	user2Id int64 `gorm:"column:user_2_id;type:int;"`
}

func (EnlistDO) TableName() string {
	return enlistTableName
}

type OrderDO struct {
	Model
	userId     int64 `gorm:"column:user_id;type:int;"`
	activityId int64 `gorm:"column:activity_id;type:int;"`
}

func (OrderDO) TableName() string {
	return orderTableName
}
