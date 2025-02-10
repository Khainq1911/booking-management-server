package model

import "time"

type Shift struct {
	Id         int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	ShiftName  string    `json:"shift_name" gorm:"column:shift_name"`
	StartTime  string    `json:"start_time" gorm:"column:start_time;type:time"`
	EndTime    string    `json:"end_time" gorm:"column:end_time;type:time"`
	TotalHours float32   `json:"total_hours" gorm:"column:total_hours"`
	Notes      string    `json:"notes" gorm:"column:notes"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

type UpdateShift struct {
	EmployeeId int       `json:"employee_id" gorm:"column:employee_id"`
	Notes      string    `json:"notes" gorm:"column:notes"`
	Status     string    `json:"status" gorm:"column:status"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

func (UpdateShift) TableName() string { return "shifts" }
