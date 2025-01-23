package model

type BookingService struct {
	Id        int `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	BookingId int `json:"booking_id" gorm:"column:booking_id"`
	ServiceId int `json:"service_id" gorm:"column:service_id"`
	Quantity  int `json:"quantity" gorm:"column:quantity"`
}

func (BookingService) TableName() string {
	return "booking_services"
}