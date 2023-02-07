package models

type JSONDeliveryTicket struct {
	Weight float32 `json:"weight"` // Total weight of the delivery
	// Duplicate of DeliveryTicket, but with id removed, because that doesn't want to be supplied by the OCR stage
}
