package model

import "time"

type Response struct {
	StatusCode int
	Message    string
	Data       interface{}
}

type Employee struct {
	Username  string
	Email     string
	Phone     string
	Password  string
	Position  string
	HiredDate time.Time
	Dob       time.Time
	Salary    float32
}
