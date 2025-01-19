package model

import "time"

type Room struct {
	Id             int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name           string    `json:"name" gorm:"column:name"`
	Description    string    `json:"description" gorm:"column:description"`
	Floor          int       `json:"floor" gorm:"column:floor"`
	CleaningStatus bool      `json:"cleaning_status" gorm:"column:cleaning_status"`
	BookingStatus  bool      `json:"booking_status" gorm:"column:booking_status"`
	TypeId         int       `json:"type_id" gorm:"type_id"`
	ImgUrl         string    `json:"img_url" gorm:"column:img_url"`
	CreatedAt      time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

type RoomAction struct {
	Name           string    `json:"name" gorm:"column:name"`
	Description    string    `json:"description" gorm:"column:description"`
	Floor          int       `json:"floor" gorm:"column:floor"`
	CleaningStatus bool      `json:"cleaning_status" gorm:"column:cleaning_status"`
	BookingStatus  bool      `json:"booking_status" gorm:"column:booking_status"`
	TypeId         int       `json:"type_id" gorm:"type_id"`
	ImgUrl         string    `json:"img_url" gorm:"column:img_url"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

func (RoomAction) TableName() string { return "rooms" }

type RoomWithType struct {
	Room
	TyperoomName        string  `json:"typeroom_name" gorm:"column:typeroom_name"`
	TyperoomDescription string  `json:"typeroom_description" gorm:"column:typeroom_description"`
	HourlyPrice         float32 `json:"hourly_price" gorm:"column:hourly_price"`
	DailyPrice          float32 `json:"daily_price" gorm:"column:daily_price"`
}
func (RoomWithType) TableName() string { return "rooms" }