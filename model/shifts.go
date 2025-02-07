package model

import "time"

type Shift struct {
	Id           int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	EmployeeId   int       `json:"employee_id" gorm:"column:employee_id"`
	EmployeeName string    `json:"employee_name" gorm:"column:employee_name"`
	StartTime    time.Time `json:"start_time" gorm:"column:start_time"`
	EndTime      time.Time `json:"end_time" gorm:"column:end_time"`
	TotalHours   float32   `json:"total_hours" gorm:"column:total_hours"`
	Status       string    `json:"status" gorm:"column:status"`
	Notes        string    `json:"notes" gorm:"column:notes"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

type UpdateShift struct {
	EmployeeId int       `json:"employee_id" gorm:"column:employee_id"`
	Notes      string    `json:"notes" gorm:"column:notes"`
	Status     string    `json:"status" gorm:"column:status"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

func (UpdateShift) TableName() string { return "shifts" }
