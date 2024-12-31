package model

import "time"

type Response struct {
	StatusCode int
	Message    string
	Data       interface{}
}

type Employee struct {
	Id        int
	Username  string
	Dob       time.Time
	Email     string
	Phone     string
	Password  string
	Position  string
	Salary    float32
	HiredDate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
