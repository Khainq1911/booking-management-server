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

type BookingWithCustomerAndRoom struct {
	Booking
	RoomId           int       `json:"room_id" gorm:"column:room_id"`
	RoomName         string    `json:"room_name" gorm:"column:room_name"`
	RoomDescription  string    `json:"room_description" gorm:"column:room_description"`
	RoomFloor        int       `json:"room_floor" gorm:"column:room_floor"`
	ImgUrl           string    `json:"img_url" gorm:"column:img_url"`
	TyperoomName     string    `json:"typeroom_name" gorm:"column:typeroom_name"`
	HourlyPrice      float32   `json:"hourly_price" gorm:"column:hourly_price"`
	DailyPrice       float32   `json:"daily_price" gorm:"column:daily_price"`
	MaxCapacity      int       `json:"max_capacity" gorm:"column:max_capacity"`
	StandardCapacity int       `json:"standard_capacity" gorm:"column:standard_capacity"`
	CustomerName     string    `json:"customer_name" gorm:"column:customer_name"`
	CustomerDob      time.Time `json:"customer_dob" gorm:"column:customer_dob"`
	CustomerEmail    string    `json:"customer_email" gorm:"column:customer_email"`
	CustomerPhone    string    `json:"customer_phone" gorm:"column:customer_phone"`
}

func (BookingWithCustomerAndRoom) TableName() string { return "bookings" }
