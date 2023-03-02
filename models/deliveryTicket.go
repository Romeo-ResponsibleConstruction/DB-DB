package models

type DeliveryTicket struct {
	Id            string  `json:"id" gorm:"not null,unique"`
	Weight        float64 `json:"weight" gorm:"type:decimal(16,3)"` // Total weight of the delivery
	Volume        float64 `json:"volume"`
	Date          string  `json:"date"`
	ImageFilepath string  `json:"image_filepath"`
}

type DeliveryTicketItem struct {
	TicketId    string `json:"ticket_id" gorm:"not null,unique"` // composite keys
	ItemId      string `json:"item_id" gorm:"not null,unique"`   // composite keys // TODO actually make composite
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
	Pieces      uint   `json:"pieces"`
}
