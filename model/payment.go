package model

import "time"

type Payment struct {
	Id          int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	BookingId   int       `json:"booking_id" gorm:"column:booking_id"`
	Method      string    `json:"method" gorm:"column:method"`
	PaymentDate time.Time `json:"payment_date" gorm:"column:payment_date;;autoCreateTime"`
	Discount    float32       `json:"discount" gorm:"column:discount"`
	Status      bool      `json:"status" gorm:"column:status;default:false"`
	FinalAmount float32   `json:"final_amount" gorm:"column:final_amount"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

type UpdateBooking struct {
	Status    bool      `json:"status" gorm:"column:status"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

func (UpdateBooking) TableName() string { return "bookings" }
