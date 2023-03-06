package models

type DeliveryTicket struct {
	Id                     string  `json:"id" gorm:"primaryKey,not null,unique"`
	Weight                 float64 `json:"weight" gorm:"type:decimal(16,3)"` // Total weight of the delivery
	WeightSuccess          bool    `json:"weight_success"`
	WeightErrorType        string  `json:"weight_error_type"`        // Type of the occurred error
	WeightErrorDescription string  `json:"weight_error_description"` // Description of the occurred error
	Volume                 float64 `json:"volume"`
	Date                   string  `json:"date"`
	ImageFilepath          string  `json:"image_filepath"`
}

type FailedChecks struct {
	Id        string `json:"id" gorm:"primaryKey,not null,unique"`
	TicketId  string `json:"ticketId" gorm:"not null"`
	Quantity  string `json:"quantity"`
	CheckName string `json:"checkName"`
}

/* Deprecated
type DeliveryTicketItem struct {
	TicketId    string `json:"ticket_id" gorm:"not null,unique"` // composite keys
	ItemId      string `json:"item_id" gorm:"not null,unique"`   // composite keys // TODO actually make composite
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
	Pieces      uint   `json:"pieces"`
}

*/
