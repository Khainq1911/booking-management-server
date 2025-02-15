package model

import (
	"time"
)

type ShiftAssignment struct {
	Id             int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	EmployeeID     int       `gorm:"column:employee_id;not null" json:"employee_id"`
	ShiftID        int       `gorm:"column:shift_id;not null" json:"shift_id"`
	AssignmentDate time.Time `gorm:"column:assignment_date;not null;type:date" json:"assignment_date"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (ShiftAssignment) TableName() string { return "shift_assignments" }

type UpdateShiftAssignment struct {
	EmployeeID int       `gorm:"column:employee_id;not null" json:"employee_id"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (UpdateShiftAssignment) TableName() string { return "shift_assignments" }

type GetScheduler struct {
	Id             int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	EmployeeID     int       `gorm:"column:employee_id;not null" json:"employee_id"`
	Username       string    `json:"username" gorm:"column:username"`
	Phone          string    `json:"phone" gorm:"column:phone"`
	Status         string    `json:"status" gorm:"column:status"`
	ShiftID        int       `gorm:"column:shift_id;not null" json:"shift_id"`
	StartTime      string    `json:"start_time" gorm:"column:start_time;type:time"`
	EndTime        string    `json:"end_time" gorm:"column:end_time;type:time"`
	AssignmentDate string    `gorm:"column:assignment_date;not null;type:date" json:"assignment_date"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (GetScheduler) TableName() string { return "shift_assignments" }
