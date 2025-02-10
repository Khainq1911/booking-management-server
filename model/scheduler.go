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

func (ShiftAssignment) TableName () string {return "shift_assignments"}

type UpdateShiftAssignment struct {
	EmployeeID     int       `gorm:"column:employee_id;not null" json:"employee_id"`
	UpdatedAt      time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}
func (UpdateShiftAssignment) TableName () string {return "shift_assignments"}