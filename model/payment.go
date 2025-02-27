package model

import "time"

type Payment struct {
	Id          int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	BookingId   int       `json:"booking_id" gorm:"column:booking_id"`
	Method      string    `json:"method" gorm:"column:method"`
	PaymentDate time.Time `json:"payment_date" gorm:"column:payment_date;autoCreateTime"`
	Discount    float32   `json:"discount" gorm:"column:discount"`
	Status      bool      `json:"status" gorm:"column:status"`
	FinalAmount float32   `json:"final_amount" gorm:"column:final_amount"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

type UpdateBooking struct {
	Status    bool      `json:"status" gorm:"column:status"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

func (UpdateBooking) TableName() string { return "bookings" }

type ListPayment struct {
    Payment       Payment  `json:"payment" gorm:"embedded"` 
    CheckInTime   time.Time `json:"check_in_time" gorm:"column:check_in_time"`
    CheckOutTime  time.Time `json:"check_out_time" gorm:"column:check_out_time"`
    TotalPrice    float32   `json:"total_price" gorm:"column:total_price"`
    EmployeeName  string    `json:"employee_name" gorm:"column:employee_name"`
    EmployeePhone string    `json:"employee_phone" gorm:"column:employee_phone"`
    CustomerName  string    `json:"customer_name" gorm:"column:customer_name"`
    CustomerPhone string    `json:"customer_phone" gorm:"column:customer_phone"`
    RoomName      string    `json:"room_name" gorm:"column:room_name"`
}

