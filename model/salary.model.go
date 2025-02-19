package model

import "time"

type Salary struct {
	Id             int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	EmployeeId     int       `json:"employee_id" gorm:"column:employee_id"`
	PayPeriodStart string    `json:"pay_period_start" gorm:"column:pay_period_start;not null"`
	PayPeriodEnd   string    `json:"pay_period_end" gorm:"column:pay_period_end;not null"`
	BaseSalary     float64   `json:"base_salary" gorm:"column:base_salary; not null"`
	Overtime       float64   `json:"overtime" gorm:"column:overtime;default:0"`
	Bonus          float64   `json:"bonus" gorm:"column:bonus;default:0"`
	Deductions     float64   `json:"deductions" gorm:"column:deductions;default:0"`
	NetSalary      float64   `json:"net_salary" gorm:"column:net_salary;not null"`
	PaymentDate    time.Time `json:"payment_date" gorm:"column:payment_date;default:current_timestamp"`
	Notes          string    `json:"notes,omitempty" gorm:"column:notes"`
	Status         string    `json:"status" gorm:"status"`
	CreatedAt      time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

func (Salary) TableName() string { return "payrolls" }

type UpdateSalary struct {
	BaseSalary float64   `json:"base_salary" gorm:"column:base_salary; not null"`
	Overtime   float64   `json:"overtime" gorm:"column:overtime;default:0"`
	Bonus      float64   `json:"bonus" gorm:"column:bonus;default:0"`
	Deductions float64   `json:"deductions" gorm:"column:deductions;default:0"`
	NetSalary  float64   `json:"net_salary" gorm:"column:net_salary;not null"`
	Status     string    `json:"status" gorm:"status"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

func (UpdateSalary) TableName() string { return "payrolls" }

type ListSalary struct {
	Id             int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	EmployeeId     int       `json:"employee_id" gorm:"column:employee_id"`
	EmployeeName   string    `json:"employee_name"`
	PayPeriodStart string    `json:"pay_period_start" gorm:"column:pay_period_start;not null"`
	PayPeriodEnd   string    `json:"pay_period_end" gorm:"column:pay_period_end;not null"`
	BaseSalary     float64   `json:"base_salary" gorm:"column:base_salary; not null"`
	Overtime       float64   `json:"overtime" gorm:"column:overtime;default:0"`
	Bonus          float64   `json:"bonus" gorm:"column:bonus;default:0"`
	Deductions     float64   `json:"deductions" gorm:"column:deductions;default:0"`
	NetSalary      float64   `json:"net_salary" gorm:"column:net_salary;not null"`
	PaymentDate    time.Time `json:"payment_date" gorm:"column:payment_date;default:current_timestamp"`
	Notes          string    `json:"notes,omitempty" gorm:"column:notes"`
	Status         string    `json:"status" gorm:"status"`
	CreatedAt      time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

func (ListSalary) TableName() string { return "payrolls" }
