package model

import "time"

type Response struct {
	StatusCode int
	Message    string
	Data       interface{}
}

type Employee struct {
	Id        int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Username  string    `json:"username" gorm:"column:username"`
	Dob       time.Time `json:"dob" gorm:"column:dob"`
	Email     string    `json:"email" gorm:"column:email"`
	Phone     string    `json:"phone" gorm:"column:phone"`
	Role      string    `json:"role" gorm:"column:role;default:employee"`
	Address   string    `json:"address" gorm:"column:address"`
	Password  string    `json:"password" gorm:"column:password"`
	Position  string    `json:"position" gorm:"column:position"`
	Salary    float32   `json:"salary" gorm:"column:salary"`
	HiredDate time.Time `json:"hired_date" gorm:"column:hired_date"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

type UpdateEmploy struct {
	Username  string    `json:"username" gorm:"column:username"`
	Dob       time.Time `json:"dob" gorm:"column:dob"`
	Email     string    `json:"email" gorm:"column:email"`
	Phone     string    `json:"phone" gorm:"column:phone"`
	Address   string    `json:"address" gorm:"column:address"`
	Position  string    `json:"position" gorm:"column:position"`
	Salary    float32   `json:"salary" gorm:"column:salary"`
	HiredDate time.Time `json:"hired_date" gorm:"column:hired_date"`
}
