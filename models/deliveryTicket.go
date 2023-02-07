package models

type DeliveryTicket struct {
	Id     string `json:"id" gorm:"not null,unique"`
	Weight uint   `json:"weight"` // Total weight of the delivery
	// TODO decide on a proper database schema
}
