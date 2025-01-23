package model

import "time"

type Booking struct {
	Id           int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	EmployeeId   int       `json:"employee_id" gorm:"column:employee_id"`
	GuestId      int       `json:"guest_id" gorm:"column:guest_id"`
	RoomId       int       `json:"room_id" gorm:"column:room_id"`
	CheckInTime  time.Time `json:"check_in_time" gorm:"column:check_in_time"`
	CheckOutTime time.Time `json:"check_out_time" gorm:"column:check_out_time"`
	TotalPrice   float32   `json:"total_price" gorm:"column:total_price"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

type RoomStatusUpdate struct {
	BookingStatus bool      `json:"booking_status" gorm:"column:booking_status"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

func (RoomStatusUpdate) TableName() string { return "rooms" }
